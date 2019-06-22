package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Todo object for demo response
type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[string]*Todo{}

//curl -H "Content-Type: application/json" -X GET http://127.0.0.1:1234/api/todos
func getTodosHandler(c *gin.Context) {
	tt := []*Todo{}
	for _, t := range todos {
		tt = append(tt, t)
	}
	c.JSON(http.StatusOK, tt)
	
}

//curl -H "Content-Type: application/json" -X GET http://127.0.0.1:1234/api/todos/1
func getTodoByIDHandler(c *gin.Context) {
	id := c.Param("id")
	t, ok := todos[id]
	if !ok {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, t)

}

//curl -H "Content-Type: application/json" -X POST -d '{"title":"Wake up","status","active"}' http://127.0.0.1:1234/api/todos
func createTodosHandler(c *gin.Context) {
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

//curl -H "Content-Type: application/json" -X PUT -d '{"title":"Wake up","status","inactive"}' http://127.0.0.1:1234/api/todos/1
func updateTodosHandler(c *gin.Context) {
	id := c.Param("id")
	t := todos[id]
	if err := c.ShouldBindJSON(t); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, t)

}

//curl -H "Content-Type: application/json" -X DELETE http://127.0.0.1:1234/api/todos/1
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
