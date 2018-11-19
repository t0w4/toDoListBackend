package main

import (
	"log"
	"net/http"
	"os"

	"github.com/t0w4/toDoListBackend/controller"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", controller.CreateTask).Methods(http.MethodPost)
	router.HandleFunc("/tasks", controller.GetTasks).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{uuid}", controller.GetTask).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{uuid}", controller.PutTask).Methods(http.MethodPut)
	router.HandleFunc("/tasks/{uuid}", controller.DeleteTask).Methods(http.MethodDelete, http.MethodOptions)
	log.Print(http.ListenAndServe("0.0.0.0:8080", router))
	os.Exit(1)
}
