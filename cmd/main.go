package main

import (
	"log"
	"net/http"

	"github.com/pmuens/goker/goker"
)

func main() {
	store := goker.NewInMemoryPlayerStore()
	server := &goker.PlayerServer{Store: store}

	log.Fatal(http.ListenAndServe(":3000", server))
}
