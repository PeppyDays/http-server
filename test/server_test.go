package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/player"
)

func TestProcessingWinsAndRetrievingThem(t *testing.T) {
	store := player.NewInMemoryPlayerStore()
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

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := player.NewInMemoryPlayerStore()
	server := player.NewPlayerServer(store)

	name := "Pepper"
	server.ServeHTTP(httptest.NewRecorder(), arrangePostScoreRequest(name))
	server.ServeHTTP(httptest.NewRecorder(), arrangePostScoreRequest(name))
	server.ServeHTTP(httptest.NewRecorder(), arrangePostScoreRequest(name))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, arrangeGetScoreRequest(name))
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "3", response.Body.String())
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, arrangeGetLeagueRequest())
		assert.Equal(t, http.StatusOK, response.Code)

		actual, err := player.DecodeLeague(response.Body)
		expected := player.League{
			{Name: "Pepper", Wins: 3},
		}
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
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

func arrangeGetLeagueRequest() *http.Request {
	request, _ := http.NewRequest(
		http.MethodGet,
		"/league",
		nil,
	)
	return request
}
