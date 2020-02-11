package endpoints

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

	log "github.com/sirupsen/logrus"
)

//GetAllTasks function gets all tasks from the task list db
func GetAllTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	getAllResults := make([]models.ResultantTask, 0, 5)
	rows, queryerr := db.Query("SELECT TimeCreatedModified, TaskTitle, DueDate, TaskDone FROM TaskList")
	if queryerr != nil {
		log.Error("Error Querying Get All Statement", queryerr)
		http.Error(w, "Error Querying Get All Statement", http.StatusInternalServerError)
	}

	for rows.Next() {
		t := models.Task{TimeCreatedModified: 0, TaskTitle: "", DueDate: 0, TaskDone: false}
		rows.Scan(&t.TimeCreatedModified, &t.TaskTitle, &t.DueDate, &t.TaskDone)
		rt := models.ConvertToResultantTask(t)
		getAllResults = append(getAllResults, rt)
	}
	if len(getAllResults) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("No Tasks Found!"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(getAllResults)
	}
}

//GetAllTodoTasks function gets all tasks from the task list db which are not completed
func GetAllTodoTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	getAllResults := make([]models.ResultantTask, 0, 5)
	rows, queryerr := db.Query("SELECT TimeCreatedModified, TaskTitle, DueDate, TaskDone FROM TaskList WHERE TaskDone = false")
	if queryerr != nil {
		log.Error("Error Querying Get All Statement", queryerr)
		http.Error(w, "Error Querying Get All Statement", http.StatusInternalServerError)
	}

	for rows.Next() {
		t := models.Task{TimeCreatedModified: 0, TaskTitle: "", DueDate: 0, TaskDone: false}
		rows.Scan(&t.TimeCreatedModified, &t.TaskTitle, &t.DueDate, &t.TaskDone)
		rt := models.ConvertToResultantTask(t)
		getAllResults = append(getAllResults, rt)
	}
	if len(getAllResults) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("No Tasks Found for the criteria!"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(getAllResults)
	}

}

//GetByTodayTasks function gets all tasks from the task list db which are due by taday's date
func GetByTodayTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	getAllResults := make([]models.ResultantTask, 0, 5)
	qtime := time.Now()
	queryunixtime := qtime.Unix()
	querytime := (queryunixtime + utilities.SecondsLeftInDay())

	rows, queryerr := db.Query("SELECT TimeCreatedModified, TaskTitle, DueDate, TaskDone FROM TaskList WHERE DueDate = " + strconv.FormatInt(querytime, 10) + " AND TaskDone = false")
	if queryerr != nil {
		log.Error("Error Querying Get by Today Statement", queryerr)
		http.Error(w, "Error Querying Get by Today Statement", http.StatusInternalServerError)
	}

	for rows.Next() {
		t := models.Task{TimeCreatedModified: 0, TaskTitle: "", DueDate: 0, TaskDone: false}
		rows.Scan(&t.TimeCreatedModified, &t.TaskTitle, &t.DueDate, &t.TaskDone)
		rt := models.ConvertToResultantTask(t)
		getAllResults = append(getAllResults, rt)
	}
	if len(getAllResults) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("No Tasks Found for the criteria!"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(getAllResults)
	}
}

//GetOverdueTasks function gets all tasks from the task list db which are overdue
func GetOverdueTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	getAllResults := make([]models.ResultantTask, 0, 5)
	qtime := time.Now()
	queryunixtime := qtime.Unix()
	querytime := (queryunixtime - utilities.SecondsOccuredInDay())

	rows, queryerr := db.Query("SELECT TimeCreatedModified, TaskTitle, DueDate, TaskDone FROM TaskList WHERE DueDate < " + strconv.FormatInt(querytime, 10) + " AND TaskDone = false")
	if queryerr != nil {
		log.Error("Error Querying Get Overdue Statement", queryerr)
		http.Error(w, "Error Querying Get Overdue Statement", http.StatusInternalServerError)
	}

	for rows.Next() {
		t := models.Task{TimeCreatedModified: 0, TaskTitle: "", DueDate: 0, TaskDone: false}
		rows.Scan(&t.TimeCreatedModified, &t.TaskTitle, &t.DueDate, &t.TaskDone)
		rt := models.ConvertToResultantTask(t)
		getAllResults = append(getAllResults, rt)
	}
	if len(getAllResults) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("No Tasks Found for the criteria!"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(getAllResults)
	}
}

//GetByTitle provides the searched task from the db with matching title
func GetByTitle(w http.ResponseWriter, r *http.Request, db *sql.DB, title string) {
	getResult := models.ResultantTask{TimeCreatedModified: "", DueDate: "", TaskTitle: "", TaskDone: false}
	rows, queryerr := db.Query("SELECT TimeCreatedModified, TaskTitle, DueDate, TaskDone FROM TaskList WHERE TaskTitle = '" + title + "' ;")
	if queryerr != nil {
		log.Error("Error Querying Get By Title Statement", queryerr)
		http.Error(w, "Error Querying Get By Title Statement", http.StatusInternalServerError)
	} else {
		rows.Next()
		t := models.Task{TimeCreatedModified: 0, TaskTitle: "", DueDate: 0, TaskDone: false}
		rows.Scan(&t.TimeCreatedModified, &t.TaskTitle, &t.DueDate, &t.TaskDone)
		rt := models.ConvertToResultantTask(t)
		getResult = rt
		if getResult.TaskTitle == "" {
			log.Error("No Tasks Found for the criteria!", queryerr)
			http.Error(w, "No Tasks Found for the criteria!", http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(getResult)
		}
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
