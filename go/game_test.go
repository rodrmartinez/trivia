package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"testing"
)

func TestGameOutput(t *testing.T) {

	golden_file, err := ioutil.ReadFile("output.txt")
	if err != nil {
		fmt.Print(err)
	}
	expected := strings.Split(string(golden_file), "\n")

	notAWinner := false

	game := NewGame()

	game.Add("Chet")
	game.Add("Pat")
	game.Add("Sue")

	for {
		game.Roll(rand.Intn(5) + 1)

		if rand.Intn(9) == 7 {
			notAWinner = game.WrongAnswer()
		} else {
			notAWinner = game.WasCorrectlyAnswered()

		}

		if !notAWinner {
			break
		}
	}
	output := strings.Split(game.output.String(), "\n")

	for i, line := range output {
		if line != expected[i] {
			t.Errorf(" Expected \"%s\", but got \" %s\" ", expected[i], line)
		}
	}

}
