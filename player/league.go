package player

import (
	"encoding/json"
	"io"
)

func ParseLeague(reader io.Reader) (league []Player, err error) {
	err = json.NewDecoder(reader).Decode(&league)
	return
}
