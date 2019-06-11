package model

import (
	"errors"
	"time"
)

var (
	tableName string
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

	tableName = "tasks"

	return err
}

func AddNewTask(taskFromClinet TaskForClient) error {
	newTask := Task{}
	newTask.TaskForClient = taskFromClinet

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	newTask.CreatedAt = time.Now().In(jst)

	err := db.Table(tableName).Create(&newTask).Error

	if err != nil {
		return errors.New("faild to add task")
	}

	return nil
}
