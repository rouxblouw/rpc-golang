package main

import (
	"log"
	"net/rpc"
)

// Todo is the API type
type Todo struct {
	Description string
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	// Add a todo
	var reply Todo
	a := Todo{"Do some RPC stuff"}
	client.Call("API.AddTodo", a, &reply)

	// Get All todos back
	var todos []Todo
	client.Call("API.GetTodos", "", &todos)

	log.Println(todos)
}
