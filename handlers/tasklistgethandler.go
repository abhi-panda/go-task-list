// Package handlers contain all the handlers required for the application
package handlers

import (
	"database/sql"
	"go-task-list/endpoints"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var logFile *os.File

//TaskListGetHandler functions handles all the Get requests that come in for task-list
func TaskListGetHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)
		switch v["type"] {
		case "all":
			endpoints.GetAllTasks(w, r, db)
		case "today":
			endpoints.GetByTodayTasks(w, r, db)
		case "overdue":
			endpoints.GetOverdueTasks(w, r, db)
		case "alltodo":
			endpoints.GetAllTodoTasks(w, r, db)
		default:
			endpoints.GetByTitle(w, r, db, v["type"])
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
