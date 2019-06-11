package model

import (
	"time"
)

type Task struct {
	TaskForClient
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type TaskForClient struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Finish     bool   `json:"finish,omitempty"`
	Importance int    `json:"importance,omitempty"`
	Hoverstar  int    `json:"hoverstar,omitempty"`
}

func CreateTable() error {
	err := db.CreateTable(&Task{}).Error
	if err != nil {
		return err
	}

	return err
}
