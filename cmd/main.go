package main

import (
	"log"
	"net/http"
	"os"

	"github.com/pmuens/goker/goker"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store := &goker.FileSystemPlayerStore{Database: db}
	server := goker.NewPlayerServer(store)

	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatalf("could not listen on port 3000 %v", err)
	}
}
