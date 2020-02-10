// Package handlers contain all the handlers required for the application
package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-task-list/models"
	"io"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	layoutISO = "2006-01-02"
)

type taskHandlerInput struct {
	TaskTitle string
	DueDate   string
	TaskDone  bool
}

var logFile *os.File
var errors map[string]string

func mapToString(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

//ValidateTaskEntry function  validates the input before they are processed into the database
func ValidateTaskEntry(t *taskHandlerInput) bool {

	// Check if Task Title was filled out
	if t.TaskTitle == "" {
		errors["titleError"] = "The Task Title field is required."
	}

	// Check if Task Due Date was filled out
	if t.DueDate == "" {
		errors["dueDateError"] = "The Task Due Date field is required."
	}

	if len(errors) > 0 {
		return false
	}
	return true

}

//UpdateTask function performs updation task on the task list
func UpdateTask(w http.ResponseWriter, r *http.Request, t *taskHandlerInput, db *sql.DB) {
	validInput := ValidateTaskEntry(t)

	if !validInput {
		log.Error(mapToString(errors))
		http.Error(w, mapToString(errors), http.StatusInternalServerError)
	}
	log.Info("Input Validated in Update Request")

	dd, _ := time.Parse(layoutISO, t.DueDate)
	unixdd := dd.Unix()
	nt := models.NewTask(t.TaskTitle, unixdd, t.TaskDone)

	var ct models.Task

	if err := db.QueryRow("SELECT * FROM TaskList WHERE TaskTitle = ", nt.TaskTitle).Scan(&ct); err == sql.ErrNoRows {
		log.Error("Task with Title " + nt.TaskTitle + " Not Found for Update!")
		http.Error(w, "Task with Title "+nt.TaskTitle+" Not Found for Update!", http.StatusInternalServerError)
	}

	log.Info("Record Found in DB!")
	statement, err := db.Prepare("UPDATE TaskList SET DueDate = ?, TaskDone = ?,TimeCreatedModified = ? WHERE TaskTitle = ?")
	if err != nil {
		log.Error("Error Creating Update Statement")
		http.Error(w, "Error Creating Update Statement", http.StatusInternalServerError)
	}

	_, execerr := statement.Exec(nt.DueDate, nt.TaskDone, nt.TimeCreatedModified, nt.TaskTitle)

	if execerr != nil {
		log.Error("Error Executing Update Statement")
		http.Error(w, "Error Executing Update Statement", http.StatusInternalServerError)
	}
	log.Info("UPDATE Successful!!")
}

//CreateTask function performs updation task on the task list
func CreateTask(w http.ResponseWriter, r *http.Request, t *taskHandlerInput, db *sql.DB) {
	validInput := ValidateTaskEntry(t)

	if !validInput {
		log.Error(mapToString(errors))
		http.Error(w, mapToString(errors), http.StatusInternalServerError)
	}
	log.Info("Input Validated in Create Request")
	dd, _ := time.Parse(layoutISO, t.DueDate)
	unixdd := dd.Unix()
	nt := models.NewTask(t.TaskTitle, unixdd, t.TaskDone)

	statement, err := db.Prepare("INSERT INTO TaskList (TaskTitle, DueDate,TaskDone,TimeCreatedModified) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Error("Error Creating Create Statement")
		http.Error(w, "Error Creating Create Statement", http.StatusInternalServerError)
	}

	_, execerr := statement.Exec(nt.TaskTitle, nt.DueDate, nt.TaskDone, nt.TimeCreatedModified)

	if execerr != nil {
		log.Error("Error Executing Create Statement")
		http.Error(w, "Error Executing Create Statement", http.StatusInternalServerError)
	}
	log.Info("New Task Created!!")
}

//DeleteTask function performs updation task on the task list
func DeleteTask(w http.ResponseWriter, r *http.Request, t *taskHandlerInput, db *sql.DB) {

	var ct models.Task

	if err := db.QueryRow("SELECT * FROM TaskList WHERE TaskTitle = ", t.TaskTitle).Scan(&ct); err == sql.ErrNoRows {
		log.Error("Task with Title " + t.TaskTitle + " Not Found for deletion!")
		http.Error(w, "Task with Title "+t.TaskTitle+" Not Found for deletion!", http.StatusInternalServerError)
	}
	log.Info("Record Found in DB!")

	statement, err := db.Prepare("DELETE FROM TaskList WHERE TaskTitle = ?")
	if err != nil {
		log.Error("Error Creating Delete Statement")
		http.Error(w, "Error Creating Delete Statement", http.StatusInternalServerError)
	}

	_, execerr := statement.Exec(t.TaskTitle)
	if execerr != nil {
		log.Error("Error Executing Delete Statement")
		http.Error(w, "Error Executing Delete Statement", http.StatusInternalServerError)
	}
	log.Info("Task Deleted!!")
}

//TaskListDMLHandler functions handles all the Data Manipulation request that come in for task-list
func TaskListDMLHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		errors = make(map[string]string)
		var t taskHandlerInput
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&t)
		if err != nil {
			log.Error("Request Body not according to the contract")
			http.Error(w, "Request Body not according to the contract", http.StatusInternalServerError)
		}

		switch r.Method {

		case "PUT":
			UpdateTask(w, r, &t, db)
		case "POST":
			CreateTask(w, r, &t, db)
		case "DELETE":
			DeleteTask(w, r, &t, db)
		default:
			log.Error("Method Not permitted")
			http.Error(w, "Method Not permitted", http.StatusInternalServerError)
		}

	})
}

func init() {
	var err error
	logFile, err = os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error starting http server : ", err)
		return
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

}
