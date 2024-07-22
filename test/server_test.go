package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/player"
	"github.com/stretchr/testify/assert"
)

func TestProcessingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := player.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), arrangePostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), arrangePostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), arrangePostScoreRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, arrangeGetScoreRequest(player))

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "3", response.Body.String())
}

type InMemoryPlayerStore struct {
	scores map[string]int
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *InMemoryPlayerStore) IncreasePlayerScore(name string) {
	s.scores[name]++
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func arrangeGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/players/%s", name),
		nil,
	)
	return request
}

func arrangePostScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/players/%s", name),
		nil,
	)
	return request
}
