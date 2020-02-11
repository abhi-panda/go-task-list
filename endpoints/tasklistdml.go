package endpoints

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-task-list/models"
	"go-task-list/utilities"
	"io"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	layoutISO = "2006-01-02"
)

type successMessage struct {
	Message string `json:"message" bson:"message"`
}

//TaskHandlerInput defines the input json structure for dml endpoints
type TaskHandlerInput struct {
	TaskTitle string `json:"taskTitle" bson:"taskTitle"`
	DueDate   string `json:"dueDate" bson:"dueDate"`
	TaskDone  bool   `json:"taskDone" bson:"taskDone"`
}

var logFile *os.File
var errors map[string]string = make(map[string]string)

func mapToString(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

//ValidateTaskEntry function  validates the input before they are processed into the database
func ValidateTaskEntry(t *TaskHandlerInput) bool {

	// Check if Task Title was filled out
	if t.TaskTitle == "" {
		errors["titleError"] = "The Task Title field is required."
	}

	// Check if Task Due Date was filled out
	if t.DueDate == "" {
		errors["dueDateError"] = "The Task Due Date field is required."
	}

	if !utilities.CheckDateStringFormat(t.DueDate) {
		errors["dueDateError"] = "The Due Date is not of pattern YYYY-MM-DD"
	}

	if len(errors) > 0 {
		return false
	}
	return true

}

//UpdateTask function performs updation task on the task list
func UpdateTask(w http.ResponseWriter, r *http.Request, t *TaskHandlerInput, db *sql.DB) {
	validInput := ValidateTaskEntry(t)

	if !validInput {
		log.Error(mapToString(errors))
		http.Error(w, mapToString(errors), http.StatusInternalServerError)
	} else {
		log.Info("Input Validated in Update Request")

		dd, _ := time.Parse(layoutISO, t.DueDate)
		unixdd := dd.Unix() + 66599
		nt := models.NewTask(t.TaskTitle, unixdd, t.TaskDone)

		var ct models.Task

		if err := db.QueryRow("SELECT * FROM TaskList WHERE TaskTitle = ", nt.TaskTitle).Scan(&ct); err == sql.ErrNoRows {
			log.Error("Task with Title "+nt.TaskTitle+" Not Found for Update!", err)
			http.Error(w, "Task with Title "+nt.TaskTitle+" Not Found for Update!", http.StatusInternalServerError)
		}

		log.Info("Record Found in DB!")
		statement, err := db.Prepare("UPDATE TaskList SET DueDate = ?, TaskDone = ?,TimeCreatedModified = ? WHERE TaskTitle = ?")
		if err != nil {
			log.Error("Error Creating Update Statement", err)
			http.Error(w, "Error Creating Update Statement", http.StatusInternalServerError)
		}

		_, execerr := statement.Exec(nt.DueDate, nt.TaskDone, nt.TimeCreatedModified, nt.TaskTitle)

		if execerr != nil {
			log.Error("Error Executing Update Statement", execerr)
			http.Error(w, "Error Executing Update Statement", http.StatusInternalServerError)
		} else {
			log.Info("UPDATE Successful!!")
			sm := successMessage{Message: "Update Completed Successfully for " + nt.TaskTitle}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(sm)
		}
	}

}

//CreateTask function performs updation task on the task list
func CreateTask(w http.ResponseWriter, r *http.Request, t *TaskHandlerInput, db *sql.DB) {
	validInput := ValidateTaskEntry(t)

	if !validInput {
		log.Error(mapToString(errors))
		http.Error(w, mapToString(errors), http.StatusInternalServerError)
	} else {
		log.Info("Input Validated in Create Request")
		dd, _ := time.Parse(layoutISO, t.DueDate)
		unixdd := dd.Unix() + 66599
		nt := models.NewTask(t.TaskTitle, unixdd, t.TaskDone)

		statement, err := db.Prepare("INSERT INTO TaskList (TaskTitle, DueDate,TaskDone,TimeCreatedModified) VALUES (?, ?, ?, ?)")
		if err != nil {
			log.Error("Error Creating Create Statement", err)
			http.Error(w, "Error Creating Create Statement", http.StatusInternalServerError)
		}

		_, execerr := statement.Exec(nt.TaskTitle, nt.DueDate, nt.TaskDone, nt.TimeCreatedModified)

		if execerr != nil {
			log.Error("Error Executing Create Statement", execerr)
			http.Error(w, "Error Executing Create Statement", http.StatusInternalServerError)
		} else {
			log.Info("New Task Created!!")
			sm := successMessage{Message: "Create Completed Successfully for " + nt.TaskTitle}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(sm)
		}
	}

}

//DeleteTask function performs updation task on the task list
func DeleteTask(w http.ResponseWriter, r *http.Request, t *TaskHandlerInput, db *sql.DB) {

	var ct models.Task

	if err := db.QueryRow("SELECT * FROM TaskList WHERE TaskTitle = ", t.TaskTitle).Scan(&ct); err == sql.ErrNoRows {
		log.Error("Task with Title "+t.TaskTitle+" Not Found for deletion!", err)
		http.Error(w, "Task with Title "+t.TaskTitle+" Not Found for deletion!", http.StatusInternalServerError)
	}
	log.Info("Record Found in DB!")

	statement, err := db.Prepare("DELETE FROM TaskList WHERE TaskTitle = ?")
	if err != nil {
		log.Error("Error Creating Delete Statement", err)
		http.Error(w, "Error Creating Delete Statement", http.StatusInternalServerError)
	}

	_, execerr := statement.Exec(t.TaskTitle)
	if execerr != nil {
		log.Error("Error Executing Delete Statement", execerr)
		http.Error(w, "Error Executing Delete Statement", http.StatusInternalServerError)
	} else {
		log.Info("Task Deleted!!")
		sm := successMessage{Message: "Record Deleted with Task Title : " + t.TaskTitle}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(sm)
	}

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
