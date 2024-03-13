package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var DB TodoDatabase

type TodoDatabase interface {
	RetrieveTodoList() []Todo
	CreateTodo(t Todo)
}

type InMemoryTodoDatabase struct {
	todoList map[uuid.UUID]Todo
}

func NewInMemoryTodoDatabase() InMemoryTodoDatabase {
	result := InMemoryTodoDatabase{}
	result.todoList = make(map[uuid.UUID]Todo)
	return result
}

func (imdb *InMemoryTodoDatabase) RetrieveTodoList() []Todo {

	todos := make([]Todo, 0)
	for key, todo := range imdb.todoList {
		fmt.Printf("%s", key.String())
		todos = append(todos, todo)
	}
	return todos

}

func (imdb *InMemoryTodoDatabase) CreateTodo(t Todo) {
	uniqueId := t.UUID
	// t := Todo{UUID: uniqueId, Title: "Zeh cade vc !"}
	imdb.todoList[uniqueId] = t

}

type Todo struct {
	UUID  uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

type TodosResponse struct {
	Items []Todo `json:"items"`
}

func main() {
	r := gin.Default()
	DB := NewInMemoryTodoDatabase()
	r.GET("/todos", func(c *gin.Context) {

		response := TodosResponse{
			Items: DB.RetrieveTodoList(),
		} // Sua lista de itens aqui
		c.JSON(http.StatusOK, response)
	})

	r.POST("/todos", func(c *gin.Context) {
		t := Todo{UUID: uuid.New(), Title: "Zeh cade vc !"}
		DB.CreateTodo(t)
		c.JSON(http.StatusNoContent, gin.H{})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
