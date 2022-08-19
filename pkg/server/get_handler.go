package server

import (
	"net/http"
	"todo/pkg/data"

	"github.com/gin-gonic/gin"
)

// GetTodos will give back the todos
func GetTodos(context *gin.Context) {
	context.JSON(http.StatusOK, data.Todo)
}

// GetTodosFromGroup gets list of specific group
func GetTodosFromGroup(context *gin.Context) {
	group := context.Param("group")
	var groupTodos []data.Tododata
	for i := range data.Todo {
		if data.Todo[i].ListenGroup == group {
			groupTodos = append(groupTodos, data.Todo[i])
		}
	}
	context.JSON(http.StatusOK, groupTodos)
}
