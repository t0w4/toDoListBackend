package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"toDoListBackend/db"
	"toDoListBackend/model"
	"toDoListBackend/view"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Init()
	if err != nil {
		view.RendorInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("db connect error: %v", err)})
		return
	}
	defer conn.Close()

	tasks, err := model.GetTasks(conn)
	if err != nil {
		view.RendorInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("get tasks error: %v", err)})
		return
	}
	view.RenderTasks(w, tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Init()
	if err != nil {
		view.RendorInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("db connect error: %v", err)})
		return
	}
	defer conn.Close()

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

	err = model.CreateTask(conn, &task)
	if err != nil {
		view.RendorInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("create task error: %v", err)})
		return
	}
	w.WriteHeader(http.StatusCreated)
}
