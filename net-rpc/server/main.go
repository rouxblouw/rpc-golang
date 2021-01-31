package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

type API int

var database []Item

func (a *API) GetDB(nothing string, reply *[]Item) error {
	*reply = database
	return nil
}

func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, v := range database {
		if v.Title == title {
			getItem = v
		}
	}

	*reply = getItem
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item
	for id, v := range database {
		if v.Title == edit.Title {
			database[id] = Item{edit.Title, edit.Body}
			changed = database[id]
		}
	}

	*reply = changed
	return nil
}

func DeleteItem(item Item, reply *Item) error {
	var deleted Item
	for id, v := range database {
		if v.Title == item.Title {
			database = append(database[:id], database[id+1:]...)
			deleted = item
		}
	}
	*reply = deleted
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
