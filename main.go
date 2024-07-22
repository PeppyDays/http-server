package main

import (
	"log"
	"net/http"

	"example.com/player"
)

func main() {
	store := &InMemoryPlayerStore{
		scores: map[string]int{
			"Arine": 10,
		},
	}
	server := player.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":8000", server))
}

type InMemoryPlayerStore struct {
	scores map[string]int
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *InMemoryPlayerStore) IncreasePlayerScore(name string) {
	score := s.scores[name] + 1
	s.scores[name] = score
}
