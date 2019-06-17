package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Todo object for demo response student
type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[string]*Todo{}

func getTodosHandler(c *gin.Context) {
	//curl -H "Content-Type: application/json" -X GET http://127.0.0.1:1234/todos
	tt := []*Todo{}
	for _, t := range todos {
		tt = append(tt, t)
	}
	c.JSON(http.StatusOK, tt)
}

func getTodoByIDHandler(c *gin.Context) {
	//curl -H "Content-Type: application/json" -X GET http://127.0.0.1:1234/todos/1

	id := c.Param("id")
	t, ok := todos[id]
	if !ok {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, t)
}

func createTodosHandler(c *gin.Context) {

	//curl -H "Content-Type: application/json" -X POST -d '{"title":"Wake up","status","active"}' http://127.0.0.1:1234/todos
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	i := len(todos)
	i++
	id := strconv.Itoa(i)
	t.ID = id
	todos[id] = &t
	c.JSON(http.StatusCreated, t)
}
func updateTodosHandler(c *gin.Context) {
	id := c.Param("id")
	t := todos[id]
	if err := c.ShouldBindJSON(t); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, t)

}
func deleteTodosHandler(c *gin.Context) {
	id := c.Param("id")
	delete(todos, id)
	c.JSON(http.StatusOK, gin.H{"status": "success"})

}

func main() {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/todos", getTodosHandler)
	api.GET("/todos/:id", getTodoByIDHandler)
	api.POST("/todos", createTodosHandler)
	api.PUT("/todos/:id", updateTodosHandler)
	api.DELETE("/todos/:id", deleteTodosHandler)
	router.Run(":1234") //listen and serve on 0.0.0.0:1234

}
