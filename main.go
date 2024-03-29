package main

import (
	"github.com/gin-contrib/static"
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
	context.JSON(http.StatusOK, allTodos)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./static", true)))
	router.GET("/api/v1/alltodos", getAllTodos)
	return router
}

func main() {
	router := setupRouter()
	err := router.Run(":3000")
	if err != nil {
		return
	}
}
