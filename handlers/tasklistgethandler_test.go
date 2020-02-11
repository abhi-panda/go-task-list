package handlers_test

import (
	"database/sql"
	"fmt"
	"go-task-list/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func TestTaskListGetHandler(t *testing.T) {
	db, dberr := sql.Open("sqlite3", "./../database/tasklist.db")

	if dberr != nil {
		t.Fatal(dberr)
	} else {
		tt := []struct {
			routeVariable string
			shouldPass    bool
		}{
			{"all", true},
			{"alltodo", true},
			{"today", true},
			{"overdue", true},
			{"Task1", true},
		}

		for _, tc := range tt {
			path := "/tasklist/"
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}
			req = mux.SetURLVars(req, map[string]string{
				"type": tc.routeVariable,
			})
			rr := httptest.NewRecorder()
			handler := http.Handler(handlers.TaskListGetHandler(db))
			handler.ServeHTTP(rr, req)
			fmt.Println(tc.routeVariable + " : ")
			fmt.Println(rr.Body.String())

			if (rr.Code == http.StatusOK) != tc.shouldPass {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, http.StatusOK)
			}
		}
	}
}
