package player_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/player"
)

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("read league correctly", func(t *testing.T) {
		// Arrange
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		store := player.NewFileSystemPlayerStore(database)

		// Act
		actual := store.GetLeague()

		// Assert
		expected := []player.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}
		assert.Equal(t, expected, actual)
	})
}
