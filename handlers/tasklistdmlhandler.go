// Package handlers contain all the handlers required for the application
package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"go-task-list/endpoints"

	log "github.com/sirupsen/logrus"
)

//TaskListDMLHandler functions handles all the Data Manipulation request that come in for task-list
func TaskListDMLHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var t endpoints.TaskHandlerInput
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&t)
		if err != nil {
			log.Error("Request Body not according to the contract", err)
			http.Error(w, "Request Body not according to the contract", http.StatusInternalServerError)
		}

		switch r.Method {

		case "PUT":
			endpoints.UpdateTask(w, r, &t, db)
		case "POST":
			endpoints.CreateTask(w, r, &t, db)
		case "DELETE":
			endpoints.DeleteTask(w, r, &t, db)
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
