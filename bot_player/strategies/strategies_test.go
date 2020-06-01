package strategies

import (
	"errors"
	"testing"
)

type testCase struct {
	board  [][]string
	player string
	x, y   int
	err    error
}

func TestLinear(t *testing.T) {
	testCases := []testCase{
		testCase{
			board: [][]string{
				[]string{"#", "#", "#"},
				[]string{"#", "#", "#"},
				[]string{"#", "#", "#"},
			},
			x: 0, y: 0, err: nil,
		},
		testCase{
			board: [][]string{
				[]string{"O", "O", "#"},
				[]string{"#", "#", "#"},
				[]string{"#", "#", "#"},
			},
			x: 0, y: 2, err: nil,
		},
		testCase{
			board: [][]string{
				[]string{"O", "O", "O"},
				[]string{"X", "#", "#"},
				[]string{"#", "#", "#"},
			},
			x: 1, y: 1, err: nil,
		},
		testCase{
			board: [][]string{
				[]string{"O", "O", "O"},
				[]string{"X", "X", "X"},
				[]string{"O", "O", "O"},
			},
			x: 0, y: 0, err: errors.New("chau"),
		},
	}

	player := "X"

	for _, tc := range testCases {
		x, y, err := linear(tc.board, player)
		t.Logf("%d:%d [%v]", x, y, err)
		if x != tc.x || y != tc.y {
			t.Fail()
		} else if !((tc.err == nil && err == nil) || (tc.err.Error() == err.Error())) {
			t.Fail()
		}
	}

}

func TestDonotLoose(t *testing.T) {
	testCases := []testCase{
		testCase{
			board: [][]string{
				[]string{"O", "O", "#"},
				[]string{"#", "#", "#"},
				[]string{"#", "#", "#"},
			},
			x: 0, y: 2, err: nil,
		},
		testCase{
			board: [][]string{
				[]string{"#", "O", "X"},
				[]string{"X", "O", "#"},
				[]string{"#", "#", "#"},
			},
			x: 2, y: 1, err: nil,
		},
		testCase{
			board: [][]string{
				[]string{"X", "X", "O"},
				[]string{"O", "X", "#"},
				[]string{"#", "O", "X"},
			},
			x: 1, y: 2, err: nil,
		},
	}

	player := "X"

	for _, tc := range testCases {
		x, y, err := donotLoose(tc.board, player)
		t.Logf("%d:%d [%v]", x, y, err)
		if x != tc.x || y != tc.y {
			t.Fail()
		} else if !((tc.err == nil && err == nil) || (tc.err.Error() == err.Error())) {
			t.Fail()
		}
	}

}

func TestTryToWin(t *testing.T) {
	testCases := []testCase{
		testCase{
			board: [][]string{
				[]string{"O", "O", "X"},
				[]string{"#", "X", "#"},
				[]string{"#", "#", "#"},
			},
			x: 2, y: 0, err: nil,
			player: "X",
		},
		testCase{
			board: [][]string{
				[]string{"#", "O", "O"},
				[]string{"X", "X", "#"},
				[]string{"#", "#", "#"},
			},
			x: 1, y: 2, err: nil,
			player: "X",
		},
		testCase{
			board: [][]string{
				[]string{"O", "X", "O"},
				[]string{"X", "O", "#"},
				[]string{"X", "X", "#"},
			},
			x: 2, y: 2, err: nil,
			player: "O",
		},
	}

	for _, tc := range testCases {
		x, y, err := tryToWin(tc.board, tc.player)
		t.Logf("%d:%d [%v]", x, y, err)
		if x != tc.x || y != tc.y {
			t.Fail()
		} else if !((tc.err == nil && err == nil) || (tc.err.Error() == err.Error())) {
			t.Fail()
		}
	}
}
