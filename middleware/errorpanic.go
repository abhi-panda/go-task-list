// Package middleware contains all the middleware required by the project
package middleware

import (
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

var logFile *os.File

//ErrorPanicHandler is a middleware utility to handle Panic Situations for all calls
func ErrorPanicHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {

			if err := recover(); err != nil {
				log.Error("Encountered Error / Panic : %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}

		}()

		next.ServeHTTP(w, r)

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
