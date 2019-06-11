package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/oribe1115/Todo-list-server/model"
)

func main() {
	_, err := model.EstablishConnection()
	if err != nil {
		log.Fatal("Cannot Connect to Database: %s", err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World.\n")
	})
	// e.GET("create/table", handler.CreateTableHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	e.Start(":" + port)
}
