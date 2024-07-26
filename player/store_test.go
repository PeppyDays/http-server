package player_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/player"
)

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("read league correctly", func(t *testing.T) {
		// Arrange
		database := createTemporalFile(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer clearTemporalFile(database)
		store := player.NewFileSystemPlayerStore(database)

		// Act
		actual := store.GetLeague()

		// Assert
		expected := player.League{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("read twice correctly", func(t *testing.T) {
		// Arrange
		database := createTemporalFile(`[
			{"Name": "Cleo", "Wins": 10}
		]`)
		defer clearTemporalFile(database)
		store := player.NewFileSystemPlayerStore(database)
		_ = store.GetLeague()

		// Act
		actual := store.GetLeague()

		// Assert
		expected := player.League{
			{Name: "Cleo", Wins: 10},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("get player score", func(t *testing.T) {
		// Arrange
		database := createTemporalFile(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer clearTemporalFile(database)
		store := player.NewFileSystemPlayerStore(database)

		// Act
		actual := store.GetPlayerScore("Chris")

		// Assert
		expected := 33
		assert.Equal(t, expected, actual)
	})

	t.Run("increase wins for existing player", func(t *testing.T) {
		// Arrange
		database := createTemporalFile(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer clearTemporalFile(database)
		store := player.NewFileSystemPlayerStore(database)

		// Act
		store.IncreasePlayerScore("Chris")

		// Assert
		actual := store.GetPlayerScore("Chris")
		expected := 34
		assert.Equal(t, expected, actual)
	})
}

func createTemporalFile(initial string) *os.File {
	file, _ := os.CreateTemp("", "db")
	_, _ = file.Write([]byte(initial))
	return file
}

func clearTemporalFile(file *os.File) {
	file.Close()
	os.Remove(file.Name())
}
