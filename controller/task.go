package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"toDoListBackend/db"
	"toDoListBackend/model"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := db.Init()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	task, err := model.GetTaskRow(conn)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	s, err := json.Marshal(task)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, string(s))
	return
}
