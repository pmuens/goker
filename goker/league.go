package goker

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func NewLeague(reader io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(reader).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}

	return nil
}
