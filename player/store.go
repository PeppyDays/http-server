package player

// TODO: This is temporal implementation, so move it into test later.
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
