package main

// Event represents a transaction log entry
type Event struct {
	sequence  uint64    // unique id
	eventType EventType // action taken
	key       string    // key affected by action
	value     string    // value of a PUT operation
}

type EventType uint64

const (
	EVENT_TYPE_PUT EventType = iota + 1
	EVENT_TYPE_DELETE
)
