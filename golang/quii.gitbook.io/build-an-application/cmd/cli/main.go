package main

import (
	"fmt"
	poker "learn-go-with-tests/build-an-application"
	"log"
	"os"
)

const (
	dbFileName = "game.db.json"
)

func main() {
	store, closeStore, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeStore()

	game := poker.NewTexasHoldem(store, poker.BlindAlerterFunc(poker.StdOutAlerter))

	fmt.Println("Let's play poker")
	poker.NewCLI(os.Stdin, os.Stdout, game).PlayPoker()
}
