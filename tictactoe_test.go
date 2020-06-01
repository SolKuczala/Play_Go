package tictactoe

import (
	"errors"
	"testing"
)

type testCase struct {
	char       string
	coordinate Coord
	game       *Game
	err        error
}

func TestPlay(t *testing.T) {
	testCases := []testCase{
		testCase{
			char:       "X",
			coordinate: Coord{X: 0, Y: 0},
			game: &Game{
				Board: [][]string{
					[]string{"#", "#", "#"},
					[]string{"#", "#", "#"},
					[]string{"#", "#", "#"},
				},
				Lastplayed: "#",
				Status:     GameStatusOngoing,
			},
			err: nil,
		},
		testCase{
			char:       "O",
			coordinate: Coord{X: 0, Y: 0},
			game: &Game{
				Board: [][]string{
					[]string{"X", "#", "#"},
					[]string{"#", "#", "#"},
					[]string{"#", "#", "#"},
				},
				Lastplayed: "X",
				Status:     GameStatusOngoing,
			},
			err: errors.New("placeholder"),
		},
	}

	for _, expected := range testCases {
		err := Play(expected.char, expected.coordinate, expected.game)
		t.Logf("[%v]", err)
		if !(expected.err == nil && err == nil || expected.err != nil && err != nil) {
			t.Fail()
		}
	}
}
