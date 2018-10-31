package controller

import (
	"fmt"
	"net/http"
	"toDoListBackend/db"
	"toDoListBackend/model"
	"toDoListBackend/view"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Init()
	defer conn.Close()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	tasks, err := model.GetTasks(conn)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	view.RenderTasks(w, tasks)
}
