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

// GetTask は path に含まれる uuid に一致する tasks テーブルの レコードを返す
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskUUID := params["uuid"]
	exist, err := model.CheckTaskExist(taskUUID)
	if err != nil {
		view.RenderInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("check task exist error: %v", err)})
		return
	}
	if !exist {
		view.RenderNotFound(w, "tasks", taskUUID)
		return
	}

	task, err := model.GetTask(taskUUID)
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

func PutTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskUUID := params["uuid"]
	exist, err := model.CheckTaskExist(taskUUID)
	if err != nil {
		view.RenderInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("check task exist error: %v", err)})
		return
	}
	if !exist {
		view.RenderNotFound(w, "tasks", taskUUID)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		view.RenderBadRequest(w, []string{fmt.Sprintf("read post body error: %v", err)})
		return
	}

	var task model.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		view.RenderInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("json parse error: %v", err)})
		return
	}

	err = model.UpdateTask(&task, taskUUID)
	if err != nil {
		view.RenderInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("create task error: %v", err)})
		return
	}
	updatedTask, err := model.GetTask(taskUUID)
	if err != nil {
		view.RenderInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("get task error: %v", err)})
		return
	}
	view.RenderTask(w, updatedTask, http.StatusOK)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskUUID := params["uuid"]
	exist, err := model.CheckTaskExist(taskUUID)
	if err != nil {
		view.RenderInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("check task exist error: %v", err)})
		return
	}
	if !exist {
		view.RenderNotFound(w, "tasks", taskUUID)
		return
	}

	err = model.DeleteTask(taskUUID)
	if err != nil {
		view.RenderInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("create task error: %v", err)})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
