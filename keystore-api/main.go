package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
)

var store = make(map[string]string)
var ErrNoSuchKey = errors.New("no such key")

func main() {
	engine := echo.New()
	engine.PUT("/:key", putValueHandler)
	log.Fatalln(engine.Start(":8080"))
}
