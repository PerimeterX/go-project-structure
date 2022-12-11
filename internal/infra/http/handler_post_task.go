package http

import (
	"encoding/json"
	"net/http"
	"todoapp/internal/core/todo"
	"todoapp/pkg/todoapp"
)

type postTaskHandler struct {
	todoService *todo.Service
}

func newPostTaskHandler(todoService *todo.Service) *postTaskHandler {
	return &postTaskHandler{todoService: todoService}
}

func (p *postTaskHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	task := &todoapp.Task{}
	dec := json.NewDecoder(request.Body)
	err := dec.Decode(task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = p.todoService.Insert(task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}
