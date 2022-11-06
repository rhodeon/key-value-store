package main

import (
	transactionLogger "cloud-native-go/keystore-api/transaction-logger"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"sync"
)

var store = struct {
	sync.RWMutex
	data map[string]string
}{
	data: make(map[string]string),
}
var logger transactionLogger.TransactionLogger

var ErrNoSuchKey = errors.New("no such key")

func main() {
	if err := initializeTransactionLogger(); err != nil {
		log.Fatalf("unable to initialize logger: %s\n", err)
	}

	engine := echo.New()

	engine.PUT("/:key", putValueHandler)
	engine.GET("/:key", getValueHandler)
	engine.DELETE("/:key", deleteValueHandler)

	log.Fatalln(engine.Start(":8080"))
}

func initializeTransactionLogger() error {
	var err error

	logger, err = transactionLogger.NewFileTransactionLogger("transactions.log")
	if err != nil {
		return fmt.Errorf("failed to create Event logger: %w", err)
	}

	event, errs := logger.ReadEvents()
	e, ok := transactionLogger.Event{}, true

	for ok && err == nil {
		select {
		case err, ok = <-errs:
		case e, ok = <-event:
			switch e.EventType {
			case transactionLogger.EVENT_TYPE_PUT:
				err = putValue(e.Key, e.Value)

			case transactionLogger.EVENT_TYPE_DELETE:
				err = deleteValue(e.Key)
			}
		}
	}

	logger.Run()
	return err
}
