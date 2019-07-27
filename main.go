package main

import (
	"log"
	"net/http"
	"os"

	"github.com/t0w4/toDoListBackend/db"

	"github.com/gorilla/mux"
	"github.com/t0w4/toDoListBackend/controller"
)

func main() {
	dbConn, err := db.Init()
	if err != nil {
		log.Printf("db init failed: %v", err)
		os.Exit(1)
	}

	tc := &controller.TaskController{dbConn}
	router := mux.NewRouter()
	router.HandleFunc("/tasks", tc.CreateTask).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/tasks", tc.GetTasks).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{uuid}", tc.GetTask).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{uuid}", tc.PutTask).Methods(http.MethodPut)
	router.HandleFunc("/tasks/{uuid}", tc.DeleteTask).Methods(http.MethodDelete, http.MethodOptions)
	log.Print(http.ListenAndServe("0.0.0.0:8080", router))
	os.Exit(1)
}
