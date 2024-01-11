package main

import (
	"log"
	"netflix/handler"
	"netflix/storage"
)

func main() {
    // creates a new pointer to a database connection
	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

    // inits the database connection
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

    // creates a pointer to a server struct
	server := handler.NewAPIServer(":8342", store)

    // spins up server
	server.Run()
}
