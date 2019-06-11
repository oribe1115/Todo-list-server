package main

import (
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"

	"github.com/oribe1115/oribe-todo-list-server/handler"
	"github.com/oribe1115/oribe-todo-list-server/model"
)

func main() {
	_, err := model.EstablishConnection()
	if err != nil {
		log.Fatal("Cannot Connect to Database: %s", err)
	}

	e := echo.New()
	e.GET("create/table", handler.CreateTableHandler)
	e.GET("/", handler.GetAllTasksHandler)
	e.POST("/", handler.AddNewTaskHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	e.Start(":" + port)
}
