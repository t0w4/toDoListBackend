package main

import (
	"log"
	"net/http"
	"os"
	"toDoListBackend/controller"
)

func main() {
	http.HandleFunc("/tasks", controller.TaskHandler)
	log.Print(http.ListenAndServe("localhost:8080", nil))
	os.Exit(1)
}
