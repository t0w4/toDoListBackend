package view

import (
	"encoding/json"
	"fmt"
	"net/http"
	"toDoListBackend/model"
)

type tasksResponse struct {
	Total int           `json:"total"`
	Tasks []*model.Task `json:"tasks"`
}

func RenderTasks(w http.ResponseWriter, tasks []*model.Task) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	s, err := json.Marshal(tasksResponse{Total: len(tasks), Tasks: tasks})
	if err != nil {
		RendorInternalServerError(w, 500, []string{"cant't encode tasks response json"})
		return
	}
	fmt.Fprintln(w, string(s))
}
