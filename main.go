package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var allTodos = []todo{
	{ID: "1", Title: "Wash dishes", Completed: false},
	{ID: "2", Title: "Do laundry", Completed: false},
	{ID: "3", Title: "Clean room", Completed: false},
}

func getAllTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, allTodos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getAllTodos)
	err := router.Run("todos:3000")
	if err != nil {
		return
	}
}
