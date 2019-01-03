package view

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/t0w4/toDoListBackend/model"
)

type tasksResponse struct {
	Total int           `json:"total"`
	Tasks []*model.Task `json:"tasks"`
}

func RenderTasks(w http.ResponseWriter, tasks []*model.Task) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	s, err := json.Marshal(tasksResponse{Total: len(tasks), Tasks: tasks})
	if err != nil {
		RenderInternalServerError(w, []string{"cant't encode tasks response json"})
		return
	}
	fmt.Fprintln(w, string(s))
}

func RenderTask(w http.ResponseWriter, task *model.Task, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	s, err := json.Marshal(task)
	if err != nil {
		RenderInternalServerError(w, []string{"cant't encode task response json"})
		return
	}
	fmt.Fprintln(w, string(s))
}
