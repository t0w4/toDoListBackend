package model

import (
	"database/sql"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetTasks(db *sql.DB) ([]*Task, error) {
	var tasks []*Task

	rows, err := db.Query(`
      SELECT 
        id,                               
        uuid, 
        title, 
        detail, 
        status,
        created_at, 
        updated_at 
       from tasks`)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var task Task
		var createdDatetime string
		var updateDatetime string
		err := rows.Scan(
			&(task.ID),
			&(task.UUID),
			&(task.Title),
			&(task.Detail),
			&(task.Status),
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
