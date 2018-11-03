package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"toDoListBackend/model"
	"toDoListBackend/view"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := model.GetTasks()
	if err != nil {
		view.RenderInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("get tasks error: %v", err)})
		return
	}
	view.RenderTasks(w, tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	task, err := model.GetTask(params["id"])
	if err != nil {
		view.RenderInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("get tasks error: %v", err)})
		return
	}
	view.RenderTask(w, task, http.StatusOK)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		view.RenderBadRequest(w, []string{fmt.Sprintf("read post body error: %v", err)})
		return
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		view.RenderInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("json parse error: %v", err)})
		return
	}

	insertID, err := model.CreateTask(&task)
	if err != nil {
		view.RenderInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("create task error: %v", err)})
		return
	}
	createdTask, err := model.GetTaskByID(insertID)
	if err != nil {
		view.RenderInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("get task error: %v", err)})
		return
	}
	view.RenderTask(w, createdTask, http.StatusCreated)
}
