// Package models defines the Models required in the project
package models

import (
	"time"
)

//Task represents and defines a Task-List Item
type Task struct {
	TimeCreatedModified int64  `json:"timeCreated" bson:"timeCreated"`
	TaskTitle           string `json:"taskTitle" bson:"taskTitle"`
	DueDate             int64  `json:"dueDate" bson:"dueDate"`
	TaskDone            bool   `json:"taskDone" bson:"taskDone"`
}

//NewTask is responsible for creating a new instance of Task type
func NewTask(taskTitle string, dueDate int64, taskDone bool) *Task {

	now := time.Now()
	unixTimestamp := now.Unix()
	t := Task{TimeCreatedModified: unixTimestamp, TaskTitle: taskTitle, DueDate: dueDate, TaskDone: taskDone}
	return &t
}
