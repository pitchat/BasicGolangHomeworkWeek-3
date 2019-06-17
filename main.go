package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Student object for demo response student
type Student struct {
	ID   string `json:"student_id"`
	Name string `json:"name"`
}

var students = map[string]Student{
	"620001": Student{ID: "620001", Name: "A"},
	"620002": Student{ID: "620002", Name: "B"},
	"620003": Student{ID: "620003", Name: "C"},
}

//maxID student max id
var maxID uint64 = 620003

func getStudentHandler(c *gin.Context) {
	//curl -H "Content-Type: application/json" -X GET http://127.0.0.1:1234/students
	response := []Student{}
	for _, s := range students {
		response = append(response, s)
	}
	c.JSON(http.StatusOK, response)
}

func postStudentHandler(c *gin.Context) {

	//curl -H "Content-Type: application/json" -X POST -d '{"name":"new student2"}' http://127.0.0.1:1234/students
	s := Student{}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	maxID = maxID + 1
	var strID string
	strID = strconv.FormatUint(maxID, 10)
	s.ID = strID
	students[strID] = s
	c.JSON(http.StatusOK, "id of new student "+strID)
}

func main() {
	r := gin.Default()

	r.GET("/students", getStudentHandler)
	r.POST("/students", postStudentHandler)
	r.Run(":1234") //listen and serve on 0.0.0.0:8080

}
