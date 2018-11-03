package main

import (
	"log"
	"net/http"
	"os"
	"toDoListBackend/controller"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", controller.TaskHandler)
	log.Print(http.ListenAndServe("localhost:8080", router))
	os.Exit(1)
}
