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
		testCasePlay{
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
		testCasePlay{
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
		testCasePlay{
			char:       "O",
			coordinate: Coord{X: 0, Y: 0},
			game: &Game{
				Board: [][]string{
					[]string{"X", "#", "#"},
					[]string{"#", "X", "#"},
					[]string{"#", "#", "X"},
				},
				Lastplayed: "X",
				Status:     GameStatusEndWithWinner,
			},
			err: errors.New("Cannot play, create another board to start again"),
		},
		testCasePlay{
			char:       "X",
			coordinate: Coord{X: 0, Y: 0},
			game: &Game{
				Board: [][]string{
					[]string{"X", "#", "#"},
					[]string{"#", "X", "#"},
					[]string{"#", "#", "X"},
				},
				Lastplayed: "X",
				Status:     GameStatusOngoing,
			},
			err: errors.New("Hey, let the other play too :)"),
		},
		testCasePlay{
			char:       "Z",
			coordinate: Coord{X: 0, Y: 0},
			game: &Game{
				Board: [][]string{
					[]string{"X", "#", "#"},
					[]string{"#", "X", "#"},
					[]string{"#", "#", "X"},
				},
				Lastplayed: "X",
				Status:     GameStatusOngoing,
			},
			err: errors.New("not a valid player"),
		},
		testCasePlay{
			char:       "X",
			coordinate: Coord{X: 3, Y: 3},
			game: &Game{
				Board: [][]string{
					[]string{"X", "#", "#"},
					[]string{"#", "X", "#"},
					[]string{"#", "#", "X"},
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
		testCaseCW{
			game: &Game{Board: [][]string{
				[]string{"X", "#", "#"},
				[]string{"#", "#", "#"},
				[]string{"#", "#", "#"},
			},
				Lastplayed: "",
				Status:     GameStatusOngoing,
			},
			expextedGameStatus: GameStatusOngoing,
		},
		testCaseCW{
			game: &Game{Board: [][]string{
				[]string{"X", "X", "X"},
				[]string{"#", "#", "#"},
				[]string{"#", "#", "#"},
			},
				Lastplayed: "X",
				Status:     GameStatusOngoing,
			},
			expextedGameStatus: GameStatusEndWithWinner,
		},
		testCaseCW{
			game: &Game{Board: [][]string{
				[]string{"O", "#", "#"},
				[]string{"#", "O", "#"},
				[]string{"#", "#", "O"},
			},
				Lastplayed: "O",
				Status:     GameStatusOngoing,
			},
			expextedGameStatus: GameStatusEndWithWinner,
		},
		testCaseCW{
			game: &Game{Board: [][]string{
				[]string{"X", "#", "#"},
				[]string{"X", "#", "#"},
				[]string{"X", "#", "#"},
			},
				Lastplayed: "X",
				Status:     GameStatusOngoing,
			},
			expextedGameStatus: GameStatusEndWithWinner,
		},
		testCaseCW{
			game: &Game{Board: [][]string{
				[]string{"#", "#", "O"},
				[]string{"#", "O", "#"},
				[]string{"O", "#", "#"},
			},
				Lastplayed: "O",
				Status:     GameStatusOngoing,
			},
			expextedGameStatus: GameStatusEndWithWinner,
		},
		testCaseCW{
			game: &Game{Board: [][]string{
				[]string{"X", "X", "O"},
				[]string{"O", "O", "X"},
				[]string{"X", "O", "O"},
			},
				Lastplayed: "X",
				Status:     GameStatusOngoing,
			},
			expextedGameStatus: GameStatusEndWithDraw,
		},
		testCaseCW{
			game: &Game{Board: [][]string{
				[]string{"#", "O", "#"},
				[]string{"#", "O", "#"},
				[]string{"#", "O", "#"},
			},
				Lastplayed: "O",
				Status:     GameStatusOngoing,
			},
			expextedGameStatus: GameStatusEndWithWinner,
		},
		testCaseCW{
			game: &Game{Board: [][]string{
				[]string{"#", "#", "#"},
				[]string{"#", "#", "#"},
				[]string{"O", "O", "O"},
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
