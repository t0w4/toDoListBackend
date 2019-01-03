package db

import (
	"database/sql"
	"os"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInit(t *testing.T) {
	t.Run("db init", func(t *testing.T) {
		actual, _ := Init()
		expected := &sql.DB{}
		if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
			t.Errorf("want %s instance, got %s", reflect.TypeOf(expected), reflect.TypeOf(actual))
		}
	})

	t.Run("db driver name", func(t *testing.T) {
		expected := "mysql"
		output := sql.Drivers()[0]
		if expected != output {
			t.Errorf("want %s , got %s", expected, output)
		}
	})
}

func Test_getParamString(t *testing.T) {
	type args struct {
		param        string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "get env", args: args{param: "GOPATH", defaultValue: "can't get env"}, want: os.Getenv("GOPATH")},
		{name: "get default value", args: args{param: "HOGE", defaultValue: "OK"}, want: "OK"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getParamString(tt.args.param, tt.args.defaultValue); got != tt.want {
				t.Errorf("getParamString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getConnectionString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "get connection string", want: "root:@tcp([localhost]:3306)/todoList"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConnectionString(); got != tt.want {
				t.Errorf("getConnectionString() = %v, want %v", got, tt.want)
			}
		})
	}
}
