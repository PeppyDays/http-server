package player

import (
	"encoding/json"
	"io"
)

type InMemoryPlayerStore struct {
	scores map[string]int
	league League
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *InMemoryPlayerStore) IncreasePlayerScore(name string) {
	s.scores[name]++
}

func (s *InMemoryPlayerStore) GetLeague() League {
	var league League
	for name, wins := range s.scores {
		league = append(league, Player{Name: name, Wins: wins})
	}
	return league
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, League{}}
}

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (s *FileSystemPlayerStore) GetPlayerScore(name string) int {
	league := s.GetLeague()
	player := league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (s *FileSystemPlayerStore) IncreasePlayerScore(name string) {
	league := s.GetLeague()
	player := league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}
	_, _ = s.database.Seek(0, io.SeekStart)
	_ = json.NewEncoder(s.database).Encode(league)
}

func (s *FileSystemPlayerStore) GetLeague() League {
	_, _ = s.database.Seek(0, io.SeekStart)
	league, _ := DecodeLeague(s.database)
	return league
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	return &FileSystemPlayerStore{database: database}
}
