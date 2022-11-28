package main

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type FakeTodoApp func(*http.Request) (*http.Response, error)

func (f FakeTodoApp) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestGetTodo2(t *testing.T) {
	client := &http.Client{
		Transport: FakeTodoApp(func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				Body: io.NopCloser(strings.NewReader(`{"id": 1, "title": "todo1"}`)),
			}, nil
		}),
	}

	cl := NewClientWithHTTPClient("/todo", client)

	got, err := cl.GetTodo(1)
	if err != nil {
		t.Fatal(err)
	}
	want := Todo{ID: 1, Title: "todo1"}
	if got != want {
		t.Errorf("Unexpected todo returned. Want %q, got %q", want, got)
	}
}

func TestListTodos2(t *testing.T) {
	want := map[int]Todo{1: Todo{ID: 1, Title: "todo1"}, 2: Todo{ID: 2, Title: "todo2"}}
	client := &http.Client{
		Transport: FakeTodoApp(func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				Body: io.NopCloser(strings.NewReader(`{ "1": {"id": 1, "title": "todo1"}, "2": {"id": 2, "title": "todo2"}}`)),
			}, nil
		}),
	}

	cl := NewClientWithHTTPClient("/todos", client)

	got, err := cl.ListTodos()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected todos returned. Want %q, got %q", want, got)
	}
}
func TestCreateTodo2(t *testing.T) {
	want := Todo{ID: 1, Title: "todo1"}
	client := &http.Client{
		Transport: FakeTodoApp(func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				Body: io.NopCloser(strings.NewReader(`{"id": 1, "title": "todo1"}`)),
			}, nil
		}),
	}

	cl := NewClientWithHTTPClient("", client)

	got, err := cl.NewTodo("todo1")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected todo returned. Want %q, got %q", want, got)
	}
}
func TestUpdateTodo2(t *testing.T) {
	want := Todo{ID: 1, Title: "todoo"}
	client := &http.Client{
		Transport: FakeTodoApp(func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				Body: io.NopCloser(strings.NewReader(`{"id": 1, "title": "todoo"}`)),
			}, nil
		}),
	}

	cl := NewClientWithHTTPClient("/todo", client)

	got, err := cl.UpdateTodo(1, "todoo")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected todo returned. Want %q, got %q", want, got)
	}
}
