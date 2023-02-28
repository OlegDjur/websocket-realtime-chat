package main

import (
	"log"
	"realtime-chat/pkg/db"
)

func main() {
	_, err := db.InitDB()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

}
