package main

import (
	"fmt"
	"log"
	"os"
	"poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem creating a file system player store, %v", err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	game := poker.NewCLI(store, os.Stdin, poker.BlindAlerterFunc(poker.StdOutAlerter))
	game.PlayPoker()
}
