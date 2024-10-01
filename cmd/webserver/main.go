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

	game := goker.NewTexasHoldem(goker.BlindAlerterFunc(goker.Alerter), store)
	server, err := goker.NewPlayerServer(store, game)

	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatalf("could not listen on port 3000 %v", err)
	}
}
