package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Task struct {
	Id        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetTaskRow(db *sql.DB) (*Task, error) {
	defer db.Close()
	var task Task
	var createdDatetime string
	var updateDatetime string

	err := db.QueryRow(`
      SELECT 
        id,                               
        uuid, 
        title, 
        detail, 
        created_at, 
        updated_at 
       from tasks`).Scan(
		&(task.Id),
		&(task.UUID),
		&(task.Title),
		&(task.Detail),
		&createdDatetime,
		&updateDatetime)
	if err != nil {
		return nil, err
	}
	task.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdDatetime) // 日時はこの日付じゃないといけない
	if err != nil {
		fmt.Println(err)
	}
	task.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updateDatetime) // 日時はこの日付じゃないといけない
	if err != nil {
		fmt.Println(err)
	}

	return &task, nil
}
