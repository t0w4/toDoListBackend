package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Init() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@/todoList")
	if err != nil {
		return nil, err
	}
	return db, nil
}
