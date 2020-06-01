package main

// TEST SWITCH PLAYER
import (
	"testing"
)

func TestPlay(t *testing.T) {
	players := [][]string{
		[]string{"X", "O"},
		[]string{"O", "X"},
		[]string{"z", ""},
	}

	for _, playerCombo := range players {
		str, _ := switchPlayer(playerCombo[0])
		t.Logf("%+v", playerCombo)
		if str != playerCombo[1] {
			t.Fail()
		}
	}
}
