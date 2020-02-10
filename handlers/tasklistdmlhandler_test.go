package handlers_test

import (
	"bytes"
	"database/sql"
	"fmt"
	"go-task-list/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestTaskListDMLCreateTask(t *testing.T) {
	db, dberr := sql.Open("sqlite3", "./../database/tasklist.db")

	if dberr != nil {
		t.Fatal(dberr)
	} else {
		var jsonStr = []byte(`{
			"TaskTitle":"TaskTest",
			"DueDate":"2020-02-11",
			"TaskDone":false
		}`)

		req, err := http.NewRequest("POST", "/tasklist", bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		handler := http.Handler(handlers.TaskListDMLHandler(db))
		handler.ServeHTTP(rr, req)

		fmt.Println("Create Task : ")
		fmt.Println(rr.Body.String())

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}

func TestTaskListDMLUpdateTask(t *testing.T) {
	db, dberr := sql.Open("sqlite3", "./../database/tasklist.db")

	if dberr != nil {
		t.Fatal(dberr)
	} else {
		var jsonStr = []byte(`{
			"TaskTitle":"TaskTest",
			"DueDate":"2020-02-11",
			"TaskDone":true
		}`)

		req, err := http.NewRequest("PUT", "/tasklist", bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		handler := http.Handler(handlers.TaskListDMLHandler(db))
		handler.ServeHTTP(rr, req)

		fmt.Println("Update Task : ")
		fmt.Println(rr.Body.String())

		if status := rr.Code; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}

func TestTaskListDMLDeleteTask(t *testing.T) {
	db, dberr := sql.Open("sqlite3", "./../database/tasklist.db")

	if dberr != nil {
		t.Fatal(dberr)
	} else {
		var jsonStr = []byte(`{
			"TaskTitle":"TaskTest",
			"DueDate":"2020-02-11",
			"TaskDone":true
		}`)

		req, err := http.NewRequest("DELETE", "/tasklist", bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		handler := http.Handler(handlers.TaskListDMLHandler(db))
		handler.ServeHTTP(rr, req)

		fmt.Println("Delete Task : ")
		fmt.Println(rr.Body.String())

		if status := rr.Code; status != http.StatusAccepted {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}
