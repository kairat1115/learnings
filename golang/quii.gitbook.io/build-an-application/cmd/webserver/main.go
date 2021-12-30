package main

import (
	poker "learn-go-with-tests/build-an-application"
	"log"
	"net/http"
)

const (
	dbFileName = "game.db.json"
)

func main() {
	store, closeStore, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal("problem creating player store", err)
	}
	defer closeStore()

	game := poker.NewTexasHoldem(store, poker.BlindAlerterFunc(poker.Alerter))

	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatal("problem creating player server", err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
