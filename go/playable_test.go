package main

import (
	"testing"
)

func TestGameIsPlayable(t *testing.T) {
	game := NewGame()
	players := [3]string{"chet", "Pat", "Sue"}

	for i, player := range players {
		game.Add(player)
		if i+1 != len(game.players) {
			t.Errorf(" %s should be player number %d, but got %d ", player, i, len(game.players))
		}
	}

	if len(game.players) != len(players) {
		t.Errorf(" Expected %d players but got %d ", len(players), len(game.players))
	}

	if len(game.players) < 2 {
		t.Errorf(" 2 players is the minimum required to play ")
	}

}
