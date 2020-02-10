// Package models defines the Models required in the project
package models

import (
	"time"
)

const (
	layoutISO = "2006-01-02"
)

//Task represents and defines a Task-List Item
type Task struct {
	TimeCreatedModified int64  `json:"timeCreated" bson:"timeCreated"`
	TaskTitle           string `json:"taskTitle" bson:"taskTitle"`
	DueDate             int64  `json:"dueDate" bson:"dueDate"`
	TaskDone            bool   `json:"taskDone" bson:"taskDone"`
}

//ResultantTask represents and defines a Task-List Item to be sent
type ResultantTask struct {
	TimeCreatedModified string `json:"timeCreated" bson:"timeCreated"`
	TaskTitle           string `json:"taskTitle" bson:"taskTitle"`
	DueDate             string `json:"dueDate" bson:"dueDate"`
	TaskDone            bool   `json:"taskDone" bson:"taskDone"`
}

//NewTask is responsible for creating a new instance of Task type
func NewTask(taskTitle string, dueDate int64, taskDone bool) *Task {

	now := time.Now()
	unixTimestamp := now.Unix()
	t := Task{TimeCreatedModified: unixTimestamp, TaskTitle: taskTitle, DueDate: dueDate, TaskDone: taskDone}
	return &t
}

//ConvertToResultantTask is responsible for converting system understandable JSON Task-list item into a more user friendly Task-list item JSON
func ConvertToResultantTask(t Task) ResultantTask {
	rt := ResultantTask{TimeCreatedModified: (time.Unix(t.TimeCreatedModified, 0).Format(layoutISO)), TaskTitle: t.TaskTitle, DueDate: (time.Unix(t.DueDate, 0).Format(layoutISO)), TaskDone: t.TaskDone}
	return rt
}
