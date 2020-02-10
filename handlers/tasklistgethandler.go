// Package handlers contain all the handlers required for the application
package handlers

import (
	"database/sql"
	"go-task-list/models"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

//GetAllTasks function gets all tasks from the task list db
func GetAllTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	getAllResults := make([]models.Task, 0, 5)
	rows, queryerr := db.Query("SELECT TimeCreatedModified, TaskTitle, DueDate, TaskDone FROM TaskList")
	if queryerr != nil {
		log.Error("Error Querying Get All Statement")
		http.Error(w, "Error Querying Get All Statement", http.StatusInternalServerError)
	}
	t := models.Task{TimeCreatedModified: unixTimestamp, TaskTitle: taskTitle, DueDate: dueDate, TaskDone: taskDone}

}

//GetByTodayTasks function gets all tasks from the task list db which are due by taday's date
func GetByTodayTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {

}

//GetOverdueTasks function gets all tasks from the task list db which are overdue
func GetOverdueTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {

}

//TaskListGetHandler functions handles all the Get requests that come in for task-list
func TaskListGetHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)
		if v["bydate"] == "all" {
			GetAllTasks(w, r, db)
		} else if v["bydate"] == "" {
			log.Error("Wrong Get Call, Check Syntax")
			http.Error(w, "Wrong Get Call, Check Syntax", http.StatusInternalServerError)
		} else if v["bydate"] == "overdue" {
			GetOverdueTasks(w, r, db)
		}
		if err != nil {
			log.Error("Wrong By Date Variable Provided")
			http.Error(w, "Wrong By Date Variable Provided", http.StatusInternalServerError)
		}
		GetByTodayTasks(w, r, db)
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
