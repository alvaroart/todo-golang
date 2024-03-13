package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

		response := TodosResponse{
			Items: []Todo{
				{UUID: uuid.New(), Title: "Title 1"},
				{UUID: uuid.New(), Title: "Title 2"},
				{UUID: uuid.New(), Title: "Title 3"},
			},
		} // Sua lista de itens aqui
		c.JSON(http.StatusOK, response)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
