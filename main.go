package main

import (
	"log"
	"net/http"
	"os"

	"example.com/player"
)

func main() {
	database, err := os.OpenFile("game.db.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("filed to jopen database")
	}
	store := player.NewFileSystemPlayerStore(database)
	server := player.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":8000", server))
}
