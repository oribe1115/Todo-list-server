package main

import (
	"log"

	"github.com/labstack/echo"

	"github.com/oribe1115/Todo-list-server/handler"
	"github.com/oribe1115/Todo-list-server/model"
)

func main() {
	_, err := model.EstablishConnection()
	if err != nil {
		log.Fatal("Cannot Connect to Database: %s", err)
	}

	e := echo.New()
	e.GET("create/table", handler.CreateTableHandler)
}
