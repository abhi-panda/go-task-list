package main

import (
	"database/sql"
	"io"
	"net/http"
	"os"

	"go-task-list/handlers"
	"go-task-list/middleware"

	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var logFile *os.File

//WECSERVERPORT defines which port our server will be running
const (
	WEBSERVERPORT = ":3000"
)

func main() {
	mw := io.MultiWriter(os.Stdout, logFile)

	db, dberr := sql.Open("sqlite3", "./database/tasklist.db")

	if dberr != nil {
		log.Printf("error connecting to the database : ", dberr)
		log.Error(dberr)
	}
	statement, createtableerr := db.Prepare("CREATE TABLE IF NOT EXISTS TaskList (TaskTitle TEXT PRIMARY KEY, TimeCreatedModified INTEGER , DueDate INTEGER , TaskDone INTEGER)")
	if createtableerr != nil {
		log.Error("Error Creating Table Statement")
		panic("Error Creating Table Statement")
	}

	_, execerr := statement.Exec()

	if execerr != nil {
		log.Error("Error Executing Create Table Statement")
		panic("Error Executing Create Table Statement")
	} else {
		log.Info("Table Existed or New Table Created!!")
		log.Printf("Server Started at " + WEBSERVERPORT)
		log.Printf("Please use https://localhost" + WEBSERVERPORT)
		defer db.Close()

		r := mux.NewRouter()
		r.Handle("/tasklist", handlers.TaskListDMLHandler(db)).Methods("PUT", "POST", "DELETE")
		r.Handle("/tasklist/{type}", handlers.TaskListGetHandler(db)).Methods("GET")
		r.HandleFunc("/", handlers.HomeHandler).Methods("GET")

		http.Handle("/", middleware.ErrorPanicHandler(ghandlers.LoggingHandler(mw, r)))
		http.ListenAndServe(WEBSERVERPORT, nil)
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
