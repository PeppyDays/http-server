package player

import (
	"encoding/json"
	"io"
)

type League []Player

func DecodeLeague(reader io.Reader) (league League, err error) {
	err = json.NewDecoder(reader).Decode(&league)
	return
}

func (l League) Find(name string) *Player {
	for i := 0; i < len(l); i++ {
		if l[i].Name == name {
			return &l[i]
		}
	}
	return nil
}
