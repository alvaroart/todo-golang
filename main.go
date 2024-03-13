package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoDatabase interface {
	RetrieveTodoList() []Todo
}

type InMemoryTodoDatabase struct {
}

func NewInMemoryTodoDatabase() InMemoryTodoDatabase {
	return InMemoryTodoDatabase{}
}

func (imdb *InMemoryTodoDatabase) RetrieveTodoList() []Todo {
	return []Todo{
		{UUID: uuid.New(), Title: "Zeh cade vc !"},
		{UUID: uuid.New(), Title: "Title 2"},
		{UUID: uuid.New(), Title: "Title 3"},
	}
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
	r.GET("/todos", func(c *gin.Context) {
		db := NewInMemoryTodoDatabase()

		response := TodosResponse{
			Items: db.RetrieveTodoList(),
		} // Sua lista de itens aqui
		c.JSON(http.StatusOK, response)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
