package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Todo is the API type
type Todo struct {
	Description string
}

// The API type to register
type API struct {
	todos []Todo
}

// AddTodo add a Todo to the internal list
func (a *API) AddTodo(t Todo, reply *Todo) error {
	a.todos = append(a.todos, t)
	*reply = t
	return nil
}

// GetTodos returns all of the Todos
func (a *API) GetTodos(n string, reply *[]Todo) error {
	*reply = a.todos
	return nil
}

func main() {
	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal(err)
	}

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal(err)
	}
}
