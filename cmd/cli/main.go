package main

import (
	"fmt"
	"log"
	"os"

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
	cli := goker.NewCLI(os.Stdin, os.Stdout, game)

	fmt.Println("Let's player poker")
	fmt.Println("Type `{Name} wins` to record a win")

	cli.PlayPoker()
}
