package model

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func TestMain(m *testing.M) {

	setUp()

	code := m.Run()

	cleanUp()

	os.Exit(code)
}

func setUp() {
	err := os.Setenv("APP_ENV", "test")
	if err != nil {
		fmt.Fprintf(os.Stdout, "set env err: %v\n", err)
		os.Exit(1)
	}

	conn, err := connect()
	if err != nil {
		fmt.Fprintf(os.Stdout, "connect mysql err: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	createDatabase(conn)
	createTable(conn)
}

func cleanUp() {
	err := os.Unsetenv("APP_ENV")
	if err != nil {
		fmt.Fprintf(os.Stdout, "unset env err: %v\n", err)
		os.Exit(1)
	}

	conn, err := connect()
	if err != nil {
		fmt.Fprintf(os.Stdout, "connect mysql err: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	dropTable(conn)
}

func connect() (*sql.DB, error) {
	connectionString := getConnectionString()
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getParamString(param string, defaultValue string) string {
	env := os.Getenv(param)
	if env != "" {
		return env
	}
	return defaultValue
}

func getConnectionString() string {
	host := getParamString("MYSQL_DB_HOST", "localhost")
	port := getParamString("MYSQL_PORT", "3306")
	user := getParamString("MYSQL_USER", "root")
	pass := getParamString("MYSQL_PASSWORD", "")
	protocol := getParamString("MYSQL_PROTOCOL", "tcp")
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/",
		user, pass, protocol, host, port)
}

func createDatabase(conn *sql.DB) {
	createDbSQL, err := os.Open("../sqls/test/00_create_db.sql")
	if err != nil {
		fmt.Fprintf(os.Stdout, "open file err: %v\n", err)
		os.Exit(1)
	}
	defer createDbSQL.Close()

	buffer, err := ioutil.ReadAll(createDbSQL)
	if err != nil {
		fmt.Fprintf(os.Stdout, "read file err: %v\n", err)
		os.Exit(1)
	}
	_, err = conn.Exec(string(buffer))
	if err != nil {
		fmt.Fprintf(os.Stdout, "create database err: %v\n", err)
		os.Exit(1)
	}
	println("database created!")
}

func createTable(conn *sql.DB) {
	createTableSQL, err := os.Open("../sqls/test/03_create_table.sql")
	if err != nil {
		fmt.Fprintf(os.Stdout, "open file err: %v\n", err)
		os.Exit(1)
	}
	defer createTableSQL.Close()

	buffer, err := ioutil.ReadAll(createTableSQL)
	if err != nil {
		fmt.Fprintf(os.Stdout, "read file err: %v\n", err)
		os.Exit(1)
	}
	_, err = conn.Exec(string(buffer))
	if err != nil {
		fmt.Fprintf(os.Stdout, "create table err: %v\n", err)
		os.Exit(1)
	}
	println("table created!")
}

func dropTable(conn *sql.DB) {
	createTableSQL, err := os.Open("../sqls/test/04_drop_table.sql")
	if err != nil {
		fmt.Fprintf(os.Stdout, "open file err: %v\n", err)
		os.Exit(1)
	}
	defer createTableSQL.Close()

	buffer, err := ioutil.ReadAll(createTableSQL)
	if err != nil {
		fmt.Fprintf(os.Stdout, "read file err: %v\n", err)
		os.Exit(1)
	}
	_, err = conn.Exec(string(buffer))
	if err != nil {
		fmt.Fprintf(os.Stdout, "drop table err: %v\n", err)
		os.Exit(1)
	}
	println("table dropped!")
}

func truncateTable() {
	conn, err := connect()
	if err != nil {
		fmt.Fprintf(os.Stdout, "connect mysql err: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	truncateTableSQL, err := os.Open("../sqls/test/02_truncate_table.sql")
	if err != nil {
		fmt.Fprintf(os.Stdout, "open file err: %v\n", err)
		os.Exit(1)
	}
	defer truncateTableSQL.Close()

	buffer, err := ioutil.ReadAll(truncateTableSQL)
	if err != nil {
		fmt.Fprintf(os.Stdout, "read file err: %v\n", err)
		os.Exit(1)
	}
	_, err = conn.Exec(string(buffer))
	if err != nil {
		fmt.Fprintf(os.Stdout, "truncate table err: %v\n", err)
		os.Exit(1)
	}
	println("table truncated!")
}

func createTasks(quantity int) []*Task {
	tasks := make([]*Task, quantity)
	for i := 0; i < quantity; i++ {
		tasks[i] = &Task{
			ID:        i + 1,
			UUID:      uuid.New().String(),
			Title:     fmt.Sprintf("title%d", i+1),
			Detail:    fmt.Sprintf("detail%d", i+1),
			Status:    fmt.Sprintf("status%d", i+1),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}
	return tasks
}

func TestCreateTask(t *testing.T) {
	type args struct {
		task *Task
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
		before  func()
		after   func()
	}{
		{name: "正常系1", args: args{createTasks(1)[0]}, want: 1, wantErr: false, before: truncateTable},
		{name: "正常系2", args: args{createTasks(1)[0]}, want: 2, wantErr: false, before: truncateTable},
	}
	for _, tt := range tests {
		if tt.before != nil {
			tt.before()
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateTask(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateTask() = %v, want %v", got, tt.want)
			}
		})
		if tt.after != nil {
			tt.after()
		}
	}
}
