package model

import (
	"database/sql"
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

func GetTasks(db *sql.DB) ([]*Task, error) {
	var task Task
	var tasks []*Task
	var createdDatetime string
	var updateDatetime string

	rows, err := db.Query(`
      SELECT 
        id,                               
        uuid, 
        title, 
        detail, 
        created_at, 
        updated_at 
       from tasks`)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
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
			return nil, err
		}
		task.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updateDatetime) // 日時はこの日付じゃないといけない
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}
