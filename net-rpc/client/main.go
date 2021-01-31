package main

import (
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	a := Item{"First", "Item"}
	// b := Item{"Second", "Item"}
	// c := Item{"Third", "Item"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.GetDB", "", &db)

	log.Println(db)
}
