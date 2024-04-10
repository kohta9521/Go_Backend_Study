package controllers

import (
	"encoding/json"
	"net/http"
	"src/study/Go_Backend_Study/handson/others/todo/models"
)


type TodoController struct {
	Model *models.TodoModel
}

func NewTodoController(m *models.TodoModel) *TodoController {
	return &TodoController{Model: m}
}

func (h *TodoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	// models/todo.goのAll()を使用してデータ取得
	todos, err := h.Model.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}

func (h *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// models/todo.goのInsert()を使用してデータ作成
	id, err := h.Model.Insert(todo.Task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	todo.ID = id
	json.NewEncoder(w).Encode(todo)
}