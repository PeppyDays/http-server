package main

import (
	"log"
	"net/http"

	"example.com/player"
)

func main() {
	store := player.NewInMemoryPlayerStore()
	server := player.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":8000", server))
}
