package view

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/t0w4/toDoListBackend/model"
)

func createTasks(quantity int) []*model.Task {
	tasks := make([]*model.Task, quantity)
	for i := 0; i < quantity; i++ {
		tasks[i] = &model.Task{
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

func TestRenderTasks(t *testing.T) {
	type args struct {
		w     http.ResponseWriter
		tasks []*model.Task
	}
	type test struct {
		name string
		args args
	}
	type want struct {
		statusCode               int
		contentType              string
		AccessControlAllowOrigin string
		status                   string
	}
	res := httptest.NewRecorder()
	greenTests := []test{
		{
			name: "正常系",
			args: args{w: res, tasks: createTasks(2)},
		},
	}

	greenWant := want{
		statusCode:               http.StatusOK,
		contentType:              "application/json; charset=utf-8",
		AccessControlAllowOrigin: "*",
	}

	for _, tt := range greenTests {
		t.Run(tt.name, func(t *testing.T) {
			RenderTasks(tt.args.w, tt.args.tasks)
			if greenWant.statusCode != res.Code {
				t.Errorf("header: status code = %d, want %d", res.Code, greenWant.statusCode)
			}

			if greenWant.contentType != res.Header().Get("Content-Type") {
				t.Errorf("header: Content-Type = %s, want %s", res.Header().Get("Content-Type"), greenWant.contentType)
			}

			if greenWant.AccessControlAllowOrigin != res.Header().Get("Access-Control-Allow-Origin") {
				t.Errorf("header: Access-Control-Allow-Origin = %s, want %s", res.Header().Get("Access-Control-Allow-Origin"), greenWant.AccessControlAllowOrigin)
			}

			var body tasksResponse
			err := json.Unmarshal(res.Body.Bytes(), &body)

			if err != nil {
				t.Fatal("can't unmarshal json")
			}

			if len(tt.args.tasks) != body.Total {
				t.Errorf("body: total = %d, want %d", body.Total, len(tt.args.tasks))
			}

			for i := 0; i < len(tt.args.tasks); i++ {
				if tt.args.tasks[i].ID != body.Tasks[i].ID {
					t.Errorf("body: task.id = %d, want %d", body.Tasks[i].ID, tt.args.tasks[i].ID)
				}
				if tt.args.tasks[i].UUID != body.Tasks[i].UUID {
					t.Errorf("body: task.uuid = %s, want %s", body.Tasks[i].UUID, tt.args.tasks[i].UUID)
				}
				if tt.args.tasks[i].Title != body.Tasks[i].Title {
					t.Errorf("body: task.title = %s, want %s", body.Tasks[i].Title, tt.args.tasks[i].Title)
				}
				if tt.args.tasks[i].Detail != body.Tasks[i].Detail {
					t.Errorf("body: task.detail = %s, want %s", body.Tasks[i].Detail, tt.args.tasks[i].Detail)
				}
				if tt.args.tasks[i].Status != body.Tasks[i].Status {
					t.Errorf("body: task.status = %s, want %s", body.Tasks[i].Status, tt.args.tasks[i].Status)
				}
			}
		})
	}
}

func TestRenderTask(t *testing.T) {
	// httptest.NewRecorderを毎回初期化するため、テーブル駆動テストの引数から外している
	type args struct {
		task       *model.Task
		statusCode int
	}
	type test struct {
		name string
		args args
	}
	type want struct {
		statusCode               int
		contentType              string
		AccessControlAllowOrigin string
		status                   string
	}
	greenTests := []test{
		{name: "正常系: 200", args: args{task: createTasks(1)[0], statusCode: http.StatusOK}},
		{name: "正常系: 201", args: args{task: createTasks(1)[0], statusCode: http.StatusCreated}},
		{name: "正常系: 204", args: args{task: createTasks(1)[0], statusCode: http.StatusNoContent}},
	}

	greenWants := []want{
		{
			statusCode:               http.StatusOK,
			contentType:              "application/json; charset=utf-8",
			AccessControlAllowOrigin: "*",
		},
		{
			statusCode:               http.StatusCreated,
			contentType:              "application/json; charset=utf-8",
			AccessControlAllowOrigin: "*",
		},
		{
			statusCode:               http.StatusNoContent,
			contentType:              "application/json; charset=utf-8",
			AccessControlAllowOrigin: "*",
		},
	}

	for i, tt := range greenTests {
		res := httptest.NewRecorder()
		t.Run(tt.name, func(t *testing.T) {
			RenderTask(res, tt.args.task, tt.args.statusCode)
			if greenWants[i].statusCode != res.Code {
				t.Errorf("header: status code = %d, want %d", res.Code, greenWants[i].statusCode)
			}

			if greenWants[i].contentType != res.Header().Get("Content-Type") {
				t.Errorf("header: Content-Type = %s, want %s", res.Header().Get("Content-Type"), greenWants[i].contentType)
			}

			if greenWants[i].AccessControlAllowOrigin != res.Header().Get("Access-Control-Allow-Origin") {
				t.Errorf("header: Access-Control-Allow-Origin = %s, want %s", res.Header().Get("Access-Control-Allow-Origin"), greenWants[i].AccessControlAllowOrigin)
			}

			var body model.Task

			err := json.Unmarshal(res.Body.Bytes(), &body)

			if err != nil {
				t.Fatal("can't unmarshal json")
			}

			if tt.args.task.ID != body.ID {
				t.Errorf("body: task.id = %d, want %d", body.ID, tt.args.task.ID)
			}
			if tt.args.task.UUID != body.UUID {
				t.Errorf("body: task.uuid = %s, want %s", body.UUID, tt.args.task.UUID)
			}
			if tt.args.task.Title != body.Title {
				t.Errorf("body: task.title = %s, want %s", body.Title, tt.args.task.Title)
			}
			if tt.args.task.Detail != body.Detail {
				t.Errorf("body: task.detail = %s, want %s", body.Detail, tt.args.task.Detail)
			}
			if tt.args.task.Status != body.Status {
				t.Errorf("body: task.status = %s, want %s", body.Status, tt.args.task.Status)
			}
		})
	}
}
