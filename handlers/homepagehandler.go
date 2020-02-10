package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

//HomeHandler renders the api registry
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	log.Info("Reached the Home Page!")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to TaskList, use {/tasklist} for create,update,delete and {/tasklist/<type (all,overdue,today)>}"))
}
