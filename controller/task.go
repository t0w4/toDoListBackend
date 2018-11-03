package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"toDoListBackend/model"
	"toDoListBackend/view"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := model.GetTasks()
	if err != nil {
		view.RendorInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("get tasks error: %v", err)})
		return
	}
	view.RenderTasks(w, tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		view.RendorBadRequest(w, []string{fmt.Sprintf("read post body error: %v", err)})
		return
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		view.RendorInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("json parse error: %v", err)})
		return
	}

	err = model.CreateTask(&task)
	if err != nil {
		view.RendorInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("create task error: %v", err)})
		return
	}
	w.WriteHeader(http.StatusCreated)
}
