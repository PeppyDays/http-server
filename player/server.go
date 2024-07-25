package player

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	http.Handler
	store PlayerStore
}

func (p *PlayerServer) getScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	p.store.IncreasePlayerScore(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) handlePlayer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getScore(w, r)
	case http.MethodPost:
		p.processWin(w, r)
	}
}

func (p *PlayerServer) handleLeague(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	league := p.store.GetLeague()
	if err := json.NewEncoder(w).Encode(league); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{
		store: store,
	}
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.handleLeague))
	router.Handle("/players/", http.HandlerFunc(p.handlePlayer))
	p.Handler = router
	return p
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	IncreasePlayerScore(name string)
	GetLeague() []Player
}

type Player struct {
	Name string
	Wins int
}
