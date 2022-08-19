package main

import (
	"net/http"
	"todo/pkg/data"
	"todo/pkg/server"

	"github.com/gin-gonic/gin"
)

func main() {
	data.MessageQueue = make(chan string)

	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/todos", server.GetTodos)
	r.GET("/todos/:group", server.GetTodosFromGroup)
	r.GET("/eventstream", Sender)
	r.POST("/todos", server.AddTodo)
	r.POST("/todos/:id", server.CheckTodo)
	return r
}

// Sender sends msg to the cli
func Sender(context *gin.Context) {
	context.Header("Content-Type", "text/event-stream")
	context.Header("Cache-Control", "no-cache")
	context.Header("Connection", "keep-alive")
	for {
		select {
		case msg := <-data.MessageQueue:
			context.String(http.StatusOK, msg)
		}
	}
}
