package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"toDoListBackend/db"
	"toDoListBackend/model"
	"toDoListBackend/view"

	"github.com/google/uuid"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Init()
	if err != nil {
		view.RendorInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("db connect error: %v", err)})
		return
	}
	defer conn.Close()
	if r.Method == http.MethodGet {
		index(w, conn)
		return
	}
	if r.Method == http.MethodPost {
		create(w, r, conn)
		return
	}
}

func index(w http.ResponseWriter, conn *sql.DB) {
	tasks, err := model.GetTasks(conn)
	if err != nil {
		view.RendorInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("get tasks error: %v", err)})
		return
	}
	view.RenderTasks(w, tasks)
}

func create(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	var task model.Task
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		view.RendorBadRequest(w, []string{fmt.Sprintf("read post body error: %v", err)})
		return
	}
	json.Unmarshal(body, &task)

	_, err = conn.Exec(
		`INSERT INTO tasks (
                                   uuid, 
                                   title, 
                                   detail, 
                                   status,
                                   created_at,
                                   updated_at
                                   ) VALUES (?, ?, ?, ?, ?, ?) `,
		uuid.New(),
		task.Title,
		task.Detail,
		task.Status,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		view.RendorInternalServerError(w, http.StatusInternalServerError, []string{fmt.Sprintf("create task error: %v", err)})
		return
	}
	w.WriteHeader(http.StatusCreated)
}
