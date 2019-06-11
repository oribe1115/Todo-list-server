package model

import (
	"time"
)

type Task struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Finish     bool   `json:"finish,omitempty"`
	Importance int    `json:"importance,omitempty"`
	Hoverstar  int    `json:"hoverstar,omitempty"`
}

type TaskInDB struct {
	Task
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func CreateTable() error {
	err := db.CreateTable(&TaskInDB{}).Error
	if err != nil {
		return err
	}

	return err
}
