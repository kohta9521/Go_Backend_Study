package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"src/study/Go_Backend_Study/handson/others/todo/controllers"
	"src/study/Go_Backend_Study/handson/others/todo/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)


func main() {
	db, err := sql.Open("mysql", "user:userpassword@tcp(localhost:3306)/todo_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	todoModel := models.NewTodoModel(db)
	todoHandler := controllers.NewTodoController(todoModel)

	router := mux.NewRouter()

	router.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
	router.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")

	fmt.Println("Server starting at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}