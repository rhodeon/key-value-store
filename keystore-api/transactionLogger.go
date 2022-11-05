package main

import (
	"bufio"
	"fmt"
	"os"
)

// transactionLogger is used to write transaction logs
// of mutating operations to a data source.
type transactionLogger interface {
	run()
	writePut(key string, value string)
	writeDelete(key string)
	readEvents() (<-chan Event, <-chan error)
	err() <-chan error
}

// transactionLogger implementation with a text file data source.
type fileTransactionLogger struct {
	events         chan<- Event // write-only channel for sending events
	errors         <-chan error // read-only channel for receiving errors during write operations
	latestSequence uint64       // most recent inserted event number
	file           *os.File     // transaction log file
}

func newFileTransactionLogger(filename string) (transactionLogger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("unable to create transaction log file: %s", err)
	}

	return &fileTransactionLogger{file: file}, nil
}

func (l *fileTransactionLogger) run() {
	events := make(chan Event, 16)
	l.events = events

	errors := make(chan error, 1)
	l.errors = errors

	go func() {
		for e := range events {
			l.latestSequence++

			// write event to log file
			_, err := fmt.Fprintf(
				l.file,
				"%d\t%d\t%s\t%s\n",
				l.latestSequence, e.eventType, e.key, e.value,
			)

			// pass any occurring error to the logger
			if err != nil {
				errors <- err
				return
			}

		}
	}()
}

// writePut sends the event of put operations to the logger.
func (l *fileTransactionLogger) writePut(key string, value string) {
	l.events <- Event{
		eventType: EVENT_TYPE_PUT,
		key:       key,
		value:     value,
	}
}

// writePut sends the event of delete operations to the logger.
// No value is set for delete.
func (l *fileTransactionLogger) writeDelete(key string) {
	l.events <- Event{
		eventType: EVENT_TYPE_DELETE,
		key:       key,
	}
}

func (l *fileTransactionLogger) err() <-chan error {
	return l.errors
}

// readEvents() fetches the events from a log file and returns a channel containing them.
// An error channel is also returned for any problem that occurs.
func (l *fileTransactionLogger) readEvents() (<-chan Event, <-chan error) {
	scanner := bufio.NewScanner(l.file)

	outEvent := make(chan Event)    // to hold events gotten from file
	outError := make(chan error, 1) // to hold any error that occurs during file reading

	go func() {
		var e Event

		defer close(outEvent)
		defer close(outError)

		// iterate over each line in the file, handling the log details
		for scanner.Scan() {
			line := scanner.Text()

			// scan the log event into the event instance e
			if _, err := fmt.Sscanf(
				line,
				"%d\t%d\t%s\t%s",
				&e.sequence, &e.eventType, &e.key, &e.value,
			); err != nil {
				outError <- fmt.Errorf("input parse error: %w", err)
				return
			}

			// sanity check to ensure the sequence is in order
			if l.latestSequence >= e.sequence {
				outError <- fmt.Errorf("transaction numbers out of sequence")
				return
			}

			// update the latest sequence of the logger and the event channel with the new event
			l.latestSequence = e.sequence
			outEvent <- e
		}

		// update error channel with any error met during scanning
		if err := scanner.Err(); err != nil {
			outError <- fmt.Errorf("transaction log read failure: %w", err)
			return
		}
	}()

	return outEvent, outError
}
