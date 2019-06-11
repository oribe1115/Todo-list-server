package main

import (
	"log"

	"github.com/oribe1115/Todo-list-server/model"
)

func main() {
	_, err := model.EstablishConnection()
	if err != nil {
		log.Fatal("Cannot Connect to Database: %s", err)
	}
}
