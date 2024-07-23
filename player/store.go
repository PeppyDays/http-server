package player

// TODO: This is temporal implementation, so move it into test later.
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
