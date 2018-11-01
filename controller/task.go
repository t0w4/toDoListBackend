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
		view.RendorInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("db connect error: %v", err)})
		return
	}
	tasks, err := model.GetTasks(conn)
	if err != nil {
		view.RendorInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("get tasks error: %v", err)})
		return
	}
	view.RenderTasks(w, tasks)
}
