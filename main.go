package main

import (
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/v4/middleware"

	"github.com/oribe1115/oribe-todo-list-server/handler"
	"github.com/oribe1115/oribe-todo-list-server/model"
)

func main() {
	_, err := model.EstablishConnection()
	if err != nil {
		log.Fatal("Cannot Connect to Database: %s", err)
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSWithConfig{
		AllowOrigins:     []string{"http://http://localhost:8080", "https://oribe1115-todolist.netlify.com/"},
		AllowCredentials: true,
	}))

	e.GET("create/table", handler.CreateTableHandler)
	e.GET("/", handler.GetAllTasksHandler)
	e.POST("/", handler.AddNewTaskHandler)
	e.PUT("/:id", handler.ChangeStatusFandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	e.Start(":" + port)
}
