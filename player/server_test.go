package player_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/player"
)

func TestGetPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		// Arrange
		store := &StubPlayerStore{
			scores: map[string]int{"Pepper": 20},
		}
		server := player.NewPlayerServer(store)
		request := arrangeGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		// Act
		server.ServeHTTP(response, request)

		// Assert
		actual := response.Body.String()
		expected := "20"
		assert.Equal(t, expected, actual)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		// Arrange
		store := &StubPlayerStore{
			scores: map[string]int{"Floyd": 10},
		}
		server := player.NewPlayerServer(store)
		request := arrangeGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		// Act
		server.ServeHTTP(response, request)

		// Assert
		actual := response.Body.String()
		expected := "10"
		assert.Equal(t, expected, actual)
	})

	t.Run("returns 200 on existing player", func(t *testing.T) {
		// Arrange
		store := &StubPlayerStore{
			scores: map[string]int{"Arine": 10},
		}
		server := player.NewPlayerServer(store)
		request := arrangeGetScoreRequest("Arine")
		response := httptest.NewRecorder()

		// Act
		server.ServeHTTP(response, request)

		// Assert
		actual := response.Code
		expected := http.StatusOK
		assert.Equal(t, expected, actual)
	})

	t.Run("returns 404 on missing player", func(t *testing.T) {
		// Arrange
		store := &StubPlayerStore{}
		server := player.NewPlayerServer(store)
		request := arrangeGetScoreRequest("Arine")
		response := httptest.NewRecorder()

		// Act
		server.ServeHTTP(response, request)

		// Assert
		actual := response.Code
		expected := http.StatusNotFound
		assert.Equal(t, expected, actual)
	})
}

func TestProcessWins(t *testing.T) {
	t.Run("returns accepted on POST", func(t *testing.T) {
		// Arrange
		store := &StubPlayerStore{
			scores: map[string]int{},
		}
		server := player.NewPlayerServer(store)
		request := arrangePostScoreRequest("Arine")
		response := httptest.NewRecorder()

		// Act
		server.ServeHTTP(response, request)

		// Assert
		actual := response.Code
		expected := http.StatusAccepted
		assert.Equal(t, expected, actual)
	})

	t.Run("records wins", func(t *testing.T) {
		// Arrange
		store := &StubPlayerStore{
			scores: map[string]int{},
		}
		server := player.NewPlayerServer(store)
		request := arrangePostScoreRequest("Arine")
		response := httptest.NewRecorder()

		// Act
		server.ServeHTTP(response, request)

		// Assert
		assert.Equal(t, []string{"Arine"}, store.increasePlayerScoreCalls)
	})
}

func TestGetLeague(t *testing.T) {
	t.Run("returns 200", func(t *testing.T) {
		// Arrnage
		store := &StubPlayerStore{}
		server := player.NewPlayerServer(store)
		request := arrangeGetLeagueRequest()
		response := httptest.NewRecorder()

		// Act
		server.ServeHTTP(response, request)

		// Assert
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("returns content-type as application/json", func(t *testing.T) {
		// Arrnage
		store := &StubPlayerStore{}
		server := player.NewPlayerServer(store)
		request := arrangeGetLeagueRequest()
		response := httptest.NewRecorder()

		// Act
		server.ServeHTTP(response, request)

		// Assert
		actual := response.Result().Header.Get("Content-Type")
		expected := "application/json"
		assert.Equal(t, expected, actual)
	})

	t.Run("returns the league correctly", func(t *testing.T) {
		// Arrnage
		expected := []player.Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		store := &StubPlayerStore{league: expected}
		server := player.NewPlayerServer(store)
		request := arrangeGetLeagueRequest()
		response := httptest.NewRecorder()

		// Act
		server.ServeHTTP(response, request)

		// Assert
		actual := parseLeague(t, response.Body)
		assert.Equal(t, expected, actual)
	})
}

type StubPlayerStore struct {
	scores map[string]int
	league []player.Player

	increasePlayerScoreCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) IncreasePlayerScore(name string) {
	s.increasePlayerScoreCalls = append(s.increasePlayerScoreCalls, name)
}

func (s *StubPlayerStore) GetLeague() []player.Player {
	return s.league
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

func parseLeague(t testing.TB, body io.Reader) (league []player.Player) {
	if err := json.NewDecoder(body).Decode(&league); err != nil {
		t.Errorf("failed to decode body as league")
	}
	return
}
