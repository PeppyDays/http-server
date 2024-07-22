package player

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		// Arrange
		store := &StubPlayerStore{
			scores: map[string]int{"Pepper": 20},
		}
		server := NewPlayerServer(store)
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
		server := NewPlayerServer(store)
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
		server := NewPlayerServer(store)
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
		server := NewPlayerServer(store)
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
		server := NewPlayerServer(store)
		request := arrangePostScoreRequest("Arine")
		response := httptest.NewRecorder()

		// Act
		server.ServeHTTP(response, request)

		//Assert
		actual := response.Code
		expected := http.StatusAccepted
		assert.Equal(t, expected, actual)
	})

	t.Run("records wins", func(t *testing.T) {
		// Arrange
		store := &StubPlayerStore{
			scores: map[string]int{},
		}
		server := NewPlayerServer(store)
		request := arrangePostScoreRequest("Arine")
		response := httptest.NewRecorder()

		// Act
		server.ServeHTTP(response, request)

		// Assert
		assert.Equal(t, []string{"Arine"}, store.increasePlayerScoreCalls)
	})
}

type StubPlayerStore struct {
	scores                   map[string]int
	increasePlayerScoreCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) IncreasePlayerScore(name string) {
	s.increasePlayerScoreCalls = append(s.increasePlayerScoreCalls, name)
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
