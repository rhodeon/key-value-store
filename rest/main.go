package main

import (
	"cloud-native-go/common"
	transaction_logger2 "cloud-native-go/common/transaction-logger"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

var store = common.Store{
	Data: make(map[string]string),
}
var logger transaction_logger2.TransactionLogger

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

	logger, err = transaction_logger2.NewFileTransactionLogger("transactions.log")
	if err != nil {
		return fmt.Errorf("failed to create Event logger: %w", err)
	}

	event, errs := logger.ReadEvents()
	e, ok := transaction_logger2.Event{}, true

	for ok && err == nil {
		select {
		case err, ok = <-errs:
		case e, ok = <-event:
			switch e.EventType {
			case transaction_logger2.EVENT_TYPE_PUT:
				err = store.PutValue(e.Key, e.Value)

			case transaction_logger2.EVENT_TYPE_DELETE:
				err = store.DeleteValue(e.Key)
			}
		}
	}

	logger.Run()
	return err
}
