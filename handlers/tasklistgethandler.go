// Package handlers contain all the handlers required for the application
package handlers

import (
	"database/sql"
	"encoding/json"
	"go-task-list/models"
	"go-task-list/utilities"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

//GetAllTasks function gets all tasks from the task list db
func GetAllTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	getAllResults := make([]models.ResultantTask, 0, 5)
	rows, queryerr := db.Query("SELECT TimeCreatedModified, TaskTitle, DueDate, TaskDone FROM TaskList")
	if queryerr != nil {
		log.Error("Error Querying Get All Statement")
		http.Error(w, "Error Querying Get All Statement", http.StatusInternalServerError)
	}

	for rows.Next() {
		t := models.Task{TimeCreatedModified: 0, TaskTitle: "", DueDate: 0, TaskDone: false}
		rows.Scan(&t.TimeCreatedModified, &t.TaskTitle, &t.DueDate, &t.TaskDone)
		rt := models.ConvertToResultantTask(t)
		getAllResults = append(getAllResults, rt)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getAllResults)
}

//GetAllTodoTasks function gets all tasks from the task list db which are not completed
func GetAllTodoTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	getAllResults := make([]models.ResultantTask, 0, 5)
	rows, queryerr := db.Query("SELECT TimeCreatedModified, TaskTitle, DueDate, TaskDone FROM TaskList WHERE TaskDone = false")
	if queryerr != nil {
		log.Error("Error Querying Get All Statement")
		http.Error(w, "Error Querying Get All Statement", http.StatusInternalServerError)
	}

	for rows.Next() {
		t := models.Task{TimeCreatedModified: 0, TaskTitle: "", DueDate: 0, TaskDone: false}
		rows.Scan(&t.TimeCreatedModified, &t.TaskTitle, &t.DueDate, &t.TaskDone)
		rt := models.ConvertToResultantTask(t)
		getAllResults = append(getAllResults, rt)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getAllResults)
}

//GetByTodayTasks function gets all tasks from the task list db which are due by taday's date
func GetByTodayTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	getAllResults := make([]models.ResultantTask, 0, 5)
	qtime := time.Now()
	queryunixtime := qtime.Unix()
	querytime := (queryunixtime + utilities.SecondsLeftInDay())

	rows, queryerr := db.Query("SELECT TimeCreatedModified, TaskTitle, DueDate, TaskDone FROM TaskList WHERE DueDate = " + strconv.FormatInt(querytime, 10) + " AND TaskDone = false")
	if queryerr != nil {
		log.Error("Error Querying Get by Today Statement")
		http.Error(w, "Error Querying Get by Today Statement", http.StatusInternalServerError)
	}

	for rows.Next() {
		t := models.Task{TimeCreatedModified: 0, TaskTitle: "", DueDate: 0, TaskDone: false}
		rows.Scan(&t.TimeCreatedModified, &t.TaskTitle, &t.DueDate, &t.TaskDone)
		rt := models.ConvertToResultantTask(t)
		getAllResults = append(getAllResults, rt)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getAllResults)
}

//GetOverdueTasks function gets all tasks from the task list db which are overdue
func GetOverdueTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	getAllResults := make([]models.ResultantTask, 0, 5)
	qtime := time.Now()
	queryunixtime := qtime.Unix()
	querytime := (queryunixtime - utilities.SecondsOccuredInDay())

	rows, queryerr := db.Query("SELECT TimeCreatedModified, TaskTitle, DueDate, TaskDone FROM TaskList WHERE DueDate < " + strconv.FormatInt(querytime, 10) + " AND TaskDone = false")
	if queryerr != nil {
		log.Error("Error Querying Get Overdue Statement")
		http.Error(w, "Error Querying Get Overdue Statement", http.StatusInternalServerError)
	}

	for rows.Next() {
		t := models.Task{TimeCreatedModified: 0, TaskTitle: "", DueDate: 0, TaskDone: false}
		rows.Scan(&t.TimeCreatedModified, &t.TaskTitle, &t.DueDate, &t.TaskDone)
		rt := models.ConvertToResultantTask(t)
		getAllResults = append(getAllResults, rt)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getAllResults)
}

//TaskListGetHandler functions handles all the Get requests that come in for task-list
func TaskListGetHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)
		switch v["type"] {
		case "all":
			GetAllTasks(w, r, db)
		case "today":
			GetByTodayTasks(w, r, db)
		case "overdue":
			GetOverdueTasks(w, r, db)
		case "alltodo":
			GetAllTodoTasks(w, r, db)
		default:
			{
				log.Error("Method Not permitted! Wrong type Value")
				http.Error(w, "Method Not permitted! Wrong type Value", http.StatusInternalServerError)
			}
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
