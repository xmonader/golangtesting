package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetTodo(t *testing.T) {
	want := Todo{ID: 1, Title: "todo"}

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	}))

	c := NewClient(s.URL + "/todos")
	got, err := c.GetTodo(1)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Unexpected todo returned. Got %q, want %q", got, want)
	}
}

func TestListTodos(t *testing.T) {
	want := map[int]Todo{1: Todo{ID: 1, Title: "todo"}, 2: Todo{ID: 2, Title: "todo2"}}

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	}))

	c := NewClient(s.URL + "/todos")
	got, err := c.ListTodos()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected todos returned. Got %q, want %q", got, want)
	}
}

func TestCreateTodo(t *testing.T) {
	want := Todo{ID: 1, Title: "todo"}

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	}))

	c := NewClient(s.URL + "/todos")
	got, err := c.NewTodo("todo")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected todo returned. Got %q, want %q", got, want)
	}
}
func TestUpdateTodo(t *testing.T) {
	want := Todo{ID: 1, Title: "todooo"}

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(want)
	}))

	c := NewClient(s.URL + "/todos")
	got, err := c.UpdateTodo(1, "todooo")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected todo returned. Got %q, want %q", got, want)
	}
}
