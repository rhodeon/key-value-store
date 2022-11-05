package main

import (
	"errors"
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

var ErrNoSuchKey = errors.New("no such key")

func main() {
	engine := echo.New()

	engine.PUT("/:key", putValueHandler)
	engine.GET("/:key", getValueHandler)
	engine.DELETE("/:key", deleteValueHandler)

	log.Fatalln(engine.Start(":8080"))
}
