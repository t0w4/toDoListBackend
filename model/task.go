package model

import (
	"time"

	"toDoListBackend/db"

	"github.com/google/uuid"
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

func GetTasks() ([]*Task, error) {
	conn, err := db.Init()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var tasks []*Task

	rows, err := conn.Query(`
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

func GetTask(taskUUID string) (*Task, error) {
	conn, err := db.Init()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var task Task
	var createdDatetime string
	var updateDatetime string

	err = conn.QueryRow(`
      SELECT 
        id,                               
        uuid, 
        title, 
        detail, 
        status,
        created_at, 
        updated_at 
       from tasks
       where uuid = ?`,
		taskUUID).Scan(
		&(task.ID),
		&(task.UUID),
		&(task.Title),
		&(task.Detail),
		&(task.Status),
		&createdDatetime,
		&updateDatetime,
	)
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

	return &task, nil
}

func CheckTaskExist(taskUUID string) (bool, error) {
	conn, err := db.Init()
	if err != nil {
		return false, err
	}
	defer conn.Close()

	var fetchRecordCount int
	err = conn.QueryRow(`
      SELECT
       count(*)
       from tasks
       where uuid = ?`,
		taskUUID).Scan(
		&fetchRecordCount,
	)
	if err != nil {
		return false, err
	}

	if fetchRecordCount > 0 {
		return true, nil
	}
	return false, nil
}

func CreateTask(task *Task) (int64, error) {
	conn, err := db.Init()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	result, err := conn.Exec(
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
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}

func GetTaskByID(taskID int64) (*Task, error) {
	conn, err := db.Init()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var task Task
	var createdDatetime string
	var updateDatetime string

	err = conn.QueryRow(`
      SELECT 
        id,                               
        uuid, 
        title, 
        detail, 
        status,
        created_at, 
        updated_at 
       from tasks
       where id = ?`,
		taskID).Scan(
		&(task.ID),
		&(task.UUID),
		&(task.Title),
		&(task.Detail),
		&(task.Status),
		&createdDatetime,
		&updateDatetime,
	)
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

	return &task, nil
}

func UpdateTask(task *Task, taskUUID string) error {
	conn, err := db.Init()
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Exec(
		`UPDATE tasks 
                   SET title = ?, 
                       detail = ?, 
                       status = ?,
                       updated_at = ?
                 WHERE uuid = ?`,
		task.Title,
		task.Detail,
		task.Status,
		time.Now(),
		taskUUID,
	)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTask(taskUUID string) error {
	conn, err := db.Init()
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Exec(
		`DELETE FROM tasks 
                 WHERE uuid = ?`,
		taskUUID,
	)
	if err != nil {
		return err
	}
	return nil
}
