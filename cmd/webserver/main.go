package main

import (
	"log"
	"net/http"

	"github.com/pmuens/goker/goker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := goker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server, _ := goker.NewPlayerServer(store)

	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatalf("could not listen on port 3000 %v", err)
	}
}
