package tictactoe

import (
	"errors"
	"testing"
)

type testCasePlay struct {
	char       string
	coordinate Coord
	game       *Game
	err        error
}

func TestPlay(t *testing.T) {
	testCases := []testCasePlay{
		{
			char:       "X",
			coordinate: Coord{X: 0, Y: 0},
			game: &Game{
				Board: [][]string{
					{"#", "#", "#"},
					{"#", "#", "#"},
					{"#", "#", "#"},
				},
				Lastplayed: "#",
				Status:     GameStatusOngoing,
			},
			err: nil,
		},
		{
			char:       "O",
			coordinate: Coord{X: 0, Y: 0},
			game: &Game{
				Board: [][]string{
					{"X", "#", "#"},
					{"#", "#", "#"},
					{"#", "#", "#"},
				},
				Lastplayed: "X",
				Status:     GameStatusOngoing,
			},
			err: errors.New("placeholder"),
		},
		{
			char:       "O",
			coordinate: Coord{X: 0, Y: 0},
			game: &Game{
				Board: [][]string{
					{"X", "#", "#"},
					{"#", "X", "#"},
					{"#", "#", "X"},
				},
				Lastplayed: "X",
				Status:     GameStatusEndWithWinner,
			},
			err: errors.New("Cannot play, create another board to start again"),
		},
		{
			char:       "X",
			coordinate: Coord{X: 0, Y: 0},
			game: &Game{
				Board: [][]string{
					{"X", "#", "#"},
					{"#", "X", "#"},
					{"#", "#", "X"},
				},
				Lastplayed: "X",
				Status:     GameStatusOngoing,
			},
			err: errors.New("Hey, let the other play too :)"),
		},
		{
			char:       "Z",
			coordinate: Coord{X: 0, Y: 0},
			game: &Game{
				Board: [][]string{
					{"X", "#", "#"},
					{"#", "X", "#"},
					{"#", "#", "X"},
				},
				Lastplayed: "X",
				Status:     GameStatusOngoing,
			},
			err: errors.New("not a valid player"),
		},
		{
			char:       "X",
			coordinate: Coord{X: 3, Y: 3},
			game: &Game{
				Board: [][]string{
					{"X", "#", "#"},
					{"#", "X", "#"},
					{"#", "#", "X"},
				},
				Lastplayed: "O",
				Status:     GameStatusOngoing,
			},
			err: errors.New("Oh oh, there is no board there my friend :/"),
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

type testCaseCW struct {
	game               *Game
	expextedGameStatus int
}

func TestCheckWinner(t *testing.T) {
	testCases := []testCaseCW{
		{
			game: &Game{Board: [][]string{
				{"X", "#", "#"},
				{"#", "#", "#"},
				{"#", "#", "#"},
			},
				Lastplayed: "",
				Status:     GameStatusOngoing,
			},
			expextedGameStatus: GameStatusOngoing,
		},
		{
			game: &Game{Board: [][]string{
				{"X", "X", "X"},
				{"#", "#", "#"},
				{"#", "#", "#"},
			},
				Lastplayed: "X",
				Status:     GameStatusOngoing,
			},
			expextedGameStatus: GameStatusEndWithWinner,
		},
		{
			game: &Game{Board: [][]string{
				{"O", "#", "#"},
				{"#", "O", "#"},
				{"#", "#", "O"},
			},
				Lastplayed: "O",
				Status:     GameStatusOngoing,
			},
			expextedGameStatus: GameStatusEndWithWinner,
		},
		{
			game: &Game{Board: [][]string{
				{"X", "#", "#"},
				{"X", "#", "#"},
				{"X", "#", "#"},
			},
				Lastplayed: "X",
				Status:     GameStatusOngoing,
			},
			expextedGameStatus: GameStatusEndWithWinner,
		},
		{
			game: &Game{Board: [][]string{
				{"#", "#", "O"},
				{"#", "O", "#"},
				{"O", "#", "#"},
			},
				Lastplayed: "O",
				Status:     GameStatusOngoing,
			},
			expextedGameStatus: GameStatusEndWithWinner,
		},
		{
			game: &Game{Board: [][]string{
				{"X", "X", "O"},
				{"O", "O", "X"},
				{"X", "O", "O"},
			},
				Lastplayed: "X",
				Status:     GameStatusOngoing,
			},
			expextedGameStatus: GameStatusEndWithDraw,
		},
		{
			game: &Game{Board: [][]string{
				{"#", "O", "#"},
				{"#", "O", "#"},
				{"#", "O", "#"},
			},
				Lastplayed: "O",
				Status:     GameStatusOngoing,
			},
			expextedGameStatus: GameStatusEndWithWinner,
		},
		{
			game: &Game{Board: [][]string{
				{"#", "#", "#"},
				{"#", "#", "#"},
				{"O", "O", "O"},
			},
				Lastplayed: "O",
				Status:     GameStatusOngoing,
			},
			expextedGameStatus: GameStatusEndWithWinner,
		},
	}
	for _, testCase := range testCases {
		checkWinner(testCase.game)
		t.Logf("%+v", testCase.game)
		t.Logf("%+v", testCase.expextedGameStatus)
		if testCase.game.Status != testCase.expextedGameStatus {
			t.Fail()
		}
	}
}

func TestNewGame(t *testing.T) {
	testCases := []int{-1, 0, 3, 9}

	for _, testCase := range testCases {
		game := NewGame(testCase)
		t.Logf("%v, [%v]", len(game.Board), testCase)
		if testCase < 0 {
			if len(game.Board) != 0 {
				t.Fail()
			}
		} else if len(game.Board) != testCase {
			t.Fail()
		}
	}
}
