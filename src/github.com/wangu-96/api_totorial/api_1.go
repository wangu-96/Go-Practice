package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var todos = []todo{
	{ID: "1", Title: "Buy groceries", Description: "Milk, eggs, bread, and fruits"},
	{ID: "2", Title: "Morning workout", Description: "30 minutes of cardio and stretching"},
	{ID: "3", Title: "Read a book", Description: "Finish at least one chapter of 'Atomic Habits'"},
	{ID: "4", Title: "Work on project", Description: "Complete the user authentication module"},
	{ID: "5", Title: "Call Mom", Description: "Check in and see how she’s doing"},
}

func getTodos(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, todos)
}

func updateTodo(c *gin.Context) {

	id := c.Param("id")
	var updatedTodo todo

	if err := c.BindJSON(&updatedTodo); err != nil {
		return
	}

	for i, t := range todos {
		if t.ID == id {
			todos[i].Title = updatedTodo.Title
			todos[i].Description = updatedTodo.Description
			c.IndentedJSON(http.StatusOK, updatedTodo)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})

}

func getTodoById(c *gin.Context) {

	id := c.Param("id")

	for _, t := range todos {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})

}

func addTodo(c *gin.Context) {

	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)

}

func main() {

	gin := gin.Default()

	gin.GET("/todos", getTodos)
	gin.POST("/todos", addTodo)
	gin.PUT("/todos/:id", updateTodo)
	gin.GET("/todos/:id", getTodoById)
	gin.Run("localhost:3001")

}
