package main

import (
	"log"
	"net/http"
	"os"
	"pocker"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := pocker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system store: %v", err)
	}

	server := pocker.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}
