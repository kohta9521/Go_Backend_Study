package models

import "database/sql"

// Todo構造体
type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

type TodoModel struct {
	DB *sql.DB
}

func NewTodoModel(DB *sql.DB) *TodoModel {
	return &TodoModel{DB: DB}
}

// 全権取得関数
func (m *TodoModel) All() ([]Todo, error) {
	rows, err := m.DB.Query("SELECT id, task FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
        var todo Todo
        if err := rows.Scan(&todo.ID, &todo.Task); err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }

	return todos, nil
}

// 新規作成関数

func (m *TodoModel) Insert(task string) (int error) {
	result, err := m.DB.Exec("INSERT INTO todos (task) VALUES (?)", task)
	if err != nil {
        return 0, err
    }
 
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }
 
    return int(id), nil
}
