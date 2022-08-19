package server

import (
	"net/http"
	"todo/pkg/data"

	"github.com/gin-gonic/gin"
)

// AddTodo adds an todo to the list
func AddTodo(context *gin.Context) {
	var newTodo data.Tododata

	if context.Request.Body == nil {
		context.String(http.StatusBadRequest, "Body requiert")
	}

	err := context.BindJSON(&newTodo)

	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
	}

	data.Todo = append(data.Todo, newTodo)
	data.MessageQueue <- "Added todo to list"
	context.IndentedJSON(http.StatusCreated, newTodo)
}

// CheckTodo sets an item to done
func CheckTodo(context *gin.Context) {
	todo := context.Param("id")

	for i := range data.Todo {
		if data.Todo[i].Title == todo {
			data.Todo[i].Done = true
			context.JSON(http.StatusAccepted, data.Todo[i])
			data.MessageQueue <- "Checked todo"
			return
		}
	}
	context.String(http.StatusNotFound, "Id not found")
}
