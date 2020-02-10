package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	log.Info("Reached the Home Page!")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte)
}
