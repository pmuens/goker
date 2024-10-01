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

	fmt.Println("Let's player poker")
	fmt.Println("Type `{Name} wins` to record a win")
	goker.NewCLI(os.Stdin, store, goker.BlindAlerterFunc(goker.StdOutAlerter)).PlayPoker()
}
