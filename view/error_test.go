package view

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRenderInternalServerError(t *testing.T) {
	res := httptest.NewRecorder()
	type args struct {
		w       http.ResponseWriter
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "header & body check", args: args{w: res, message: "can't open file"}},
	}

	want := struct {
		statusCode               int
		contentType              string
		AccessControlAllowOrigin string
		status                   string
	}{
		statusCode:               http.StatusInternalServerError,
		contentType:              "application/json; charset=utf-8",
		AccessControlAllowOrigin: "*",
		status: "internal server error",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RenderInternalServerError(tt.args.w, tt.args.message)
			if want.statusCode != res.Code {
				t.Errorf("header: status code = %d, want %d", res.Code, want.statusCode)
			}

			if want.contentType != res.Header().Get("Content-Type") {
				t.Errorf("header: Content-Type = %s, want %s", res.Header().Get("Content-Type"), want.contentType)
			}

			if want.AccessControlAllowOrigin != res.Header().Get("Access-Control-Allow-Origin") {
				t.Errorf("header: Access-Control-Allow-Origin = %s, want %s", res.Header().Get("Access-Control-Allow-Origin"), want.AccessControlAllowOrigin)
			}

			var body errorResponse
			err := json.Unmarshal(res.Body.Bytes(), &body)

			if err != nil {
				t.Fatal("can't unmarshal json")
			}

			if want.status != body.Status {
				t.Errorf("body: status = %s, want %s", body.Status, want.status)
			}

			if tt.args.message != body.ErrorMessage {
				t.Errorf("body: message = %s, want %s", body.ErrorMessage, tt.args.message)
			}
		})
	}
}

func TestRenderBadRequest(t *testing.T) {
	res := httptest.NewRecorder()
	type args struct {
		w        http.ResponseWriter
		messages []string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "header & body check", args: args{w: res, messages: []string{"title can't be empty", "created_at can't be empty"}}},
	}

	want := struct {
		statusCode               int
		contentType              string
		AccessControlAllowOrigin string
		status                   string
	}{
		statusCode:               http.StatusBadRequest,
		contentType:              "application/json; charset=utf-8",
		AccessControlAllowOrigin: "*",
		status: "bad request",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RenderBadRequest(tt.args.w, tt.args.messages)
			if want.statusCode != res.Code {
				t.Errorf("header: status code = %d, want %d", res.Code, want.statusCode)
			}

			if want.contentType != res.Header().Get("Content-Type") {
				t.Errorf("header: Content-Type = %s, want %s", res.Header().Get("Content-Type"), want.contentType)
			}

			if want.AccessControlAllowOrigin != res.Header().Get("Access-Control-Allow-Origin") {
				t.Errorf("header: Access-Control-Allow-Origin = %s, want %s", res.Header().Get("Access-Control-Allow-Origin"), want.AccessControlAllowOrigin)
			}

			var body errorsResponse
			err := json.Unmarshal(res.Body.Bytes(), &body)

			if err != nil {
				t.Fatal("can't unmarshal json")
			}

			if want.status != body.Status {
				t.Errorf("body: status = %s, want %s", body.Status, want.status)
			}
			if strings.Join(tt.args.messages, "") != strings.Join(body.ErrorMessages, "") {
				t.Errorf("body: messages = %v, want %v", body.ErrorMessages, tt.args.messages)
			}
		})
	}
}

func TestRenderNotFound(t *testing.T) {
	res := httptest.NewRecorder()
	type args struct {
		w         http.ResponseWriter
		tableName string
		uuid      string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "header & body check", args: args{w: res, tableName: "tasks", uuid: "cfdfc19a-895d-444d-bbb1-f47ad3496fb"}},
	}

	want := struct {
		statusCode               int
		contentType              string
		AccessControlAllowOrigin string
		status                   string
	}{
		statusCode:               http.StatusNotFound,
		contentType:              "application/json; charset=utf-8",
		AccessControlAllowOrigin: "*",
		status: "not found",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RenderNotFound(tt.args.w, tt.args.tableName, tt.args.uuid)
			if want.statusCode != res.Code {
				t.Errorf("header: status code = %d, want %d", res.Code, want.statusCode)
			}

			if want.contentType != res.Header().Get("Content-Type") {
				t.Errorf("header: Content-Type = %s, want %s", res.Header().Get("Content-Type"), want.contentType)
			}

			if want.AccessControlAllowOrigin != res.Header().Get("Access-Control-Allow-Origin") {
				t.Errorf("header: Access-Control-Allow-Origin = %s, want %s", res.Header().Get("Access-Control-Allow-Origin"), want.AccessControlAllowOrigin)
			}

			var body errorResponse
			err := json.Unmarshal(res.Body.Bytes(), &body)

			if err != nil {
				t.Fatal("can't unmarshal json")
			}

			if want.status != body.Status {
				t.Errorf("body: status = %s, want %s", body.Status, want.status)
			}

			if fmt.Sprintf("Couldn't find %s with 'uuid'=%s", tt.args.tableName, tt.args.uuid) != body.ErrorMessage {
				t.Errorf("body: message = %s, want %s", body.ErrorMessage, fmt.Sprintf("Couldn't find %s with 'uuid'=%s", tt.args.tableName, tt.args.uuid))
			}
		})
	}
}
