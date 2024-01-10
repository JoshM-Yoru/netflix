package main

import (
	"log"
	"netflix/handler"
	"netflix/storage"
)

func main() {
	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := handler.NewAPIServer(":8342", store)

	server.Run()
}
