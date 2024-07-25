package player

import (
	"io"
)

type InMemoryPlayerStore struct {
	scores map[string]int
	league []Player
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *InMemoryPlayerStore) IncreasePlayerScore(name string) {
	s.scores[name]++
}

func (s *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range s.scores {
		league = append(league, Player{Name: name, Wins: wins})
	}
	return league
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, []Player{}}
}

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (s *FileSystemPlayerStore) GetLeague() []Player {
	_, _ = s.database.Seek(0, io.SeekStart)
	league, _ := ParseLeague(s.database)
	return league
}

func NewFileSystemPlayerStore(database io.ReadSeeker) *FileSystemPlayerStore {
	return &FileSystemPlayerStore{database: database}
}
