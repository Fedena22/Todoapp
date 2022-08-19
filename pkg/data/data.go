package data

// Todo in memory data
var Todo = []Tododata{
	{ListenGroup: "Peter", Title: "1", Description: "Test1", Done: false},
	{ListenGroup: "Pan", Title: "1", Description: "Test2", Done: true},
}

// MessageQueue message queue for actions
var MessageQueue chan string
