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

	fmt.Println("Let's play poker")
	fmt.Println("Type {name} wins to records a win")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
