package http

import (
	"encoding/json"
	"net/http"
)

type getFutureTasksHandler struct {
	todoService TODOService
}

func newGetFutureTasksHandler(todoService TODOService) *getFutureTasksHandler {
	return &getFutureTasksHandler{todoService: todoService}
}

func (g *getFutureTasksHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	tasks, err := g.todoService.FutureTasks()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(tasks)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
