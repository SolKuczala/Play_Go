package strategies

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const playerX string = "X"
const playerO string = "O"
const noPlayer string = "#"

type coord struct {
	i int
	j int
}

/*StrategiesMap : declares types of strategies to choose for the bot in game
 */
var StrategiesMap = map[string]PlayStrategy{
	"random": PlayStrategy{
		Gen: func(board [][]string, player string) (int, int, error) {
			rand.Seed(time.Now().UnixNano())
			randomNum := func(min, max int) int {
				return rand.Intn(max-min) + min
			}
			return randomNum(0, len(board)), randomNum(0, len(board)), nil
		},
	},
	"linear": PlayStrategy{
		Gen: linear,
	},
	"donot_loose": PlayStrategy{
		Gen: donotLoose,
	},
	"try_to_win": PlayStrategy{
		Gen: tryToWin,
	},
}

type PlayStrategy struct {
	//no deberia ser puntero a board??
	Gen func(board [][]string, player string) (int, int, error)
}

/*Play : inner function of Playstrategy to make the play*/
func (ps PlayStrategy) Play(baseURL, player string, tries int, board [][]string) error {
	isOk := false
	for !isOk && (tries > 0) {
		x, y, err := ps.Gen(board, player)
		if err != nil {
			return err
		}
		if err := ps.sendPlay(baseURL, player, x, y); err == nil {
			isOk = true
		} else {
			tries--
		}
	}
	if !isOk && tries <= 0 {
		return errors.New("Too many tries")
	}
	return nil
}

func (ps PlayStrategy) sendPlay(baseURL string, player string, coordX, coordY int) error {
	fmt.Println("playing...")
	client := &http.Client{}
	// Create request, send to my api
	req, err := http.NewRequest("PUT", baseURL+"/send-play/"+player+"/"+strconv.Itoa(coordX)+"/"+strconv.Itoa(coordY), nil) //+strconv.Itoa(randomNum)+"/"+strconv.Itoa(randomNum)/ or url.Values{"player": {"X"}, "row": {"1"}, "column": {"0"}
	if err != nil {
		fmt.Println(err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("%g", err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%g", err)
	}
	// Display Results
	fmt.Println("response Status : ", resp.Status)
	//fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))

	if resp.StatusCode != http.StatusOK {
		return errors.New("Response status is not 200")
	}
	return nil
}

var linear = func(board [][]string, player string) (int, int, error) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == "#" {
				return i, j, nil
			}
		}
	}
	return 0, 0, errors.New("chau")
}

var donotLoose = func(board [][]string, player string) (int, int, error) {
	eligiblePlaceToPlay := coord{i: -1, j: -1}
	var maxOpponentPlays int = 0

	for _, search := range []string{"row", "column", "diag1", "diag2"} {
		idealPlay(board, player, search, &eligiblePlaceToPlay, &maxOpponentPlays)
	}
	if eligiblePlaceToPlay.i == -1 {
		return linear(board, player)
	}
	return eligiblePlaceToPlay.i, eligiblePlaceToPlay.j, nil
}

var tryToWin = func(board [][]string, player string) (int, int, error) {
	if player == playerO {
		player = playerX
	} else {
		player = playerO
	}
	return donotLoose(board, player)
}

//TODO: a
func idealPlay(board [][]string, player string, search string, eligiblePlaceToPlay *coord, maxOpponentPlays *int, skipOnMatch bool) {
	if search == "row" || search == "column" {
		for i := 0; i < len(board); i++ {
			var OpponentQ int
			var sectionPlayerQ int
			var emptySpace coord

			for j := 0; j < len(board); j++ {
				var element string
				if search == "row" {
					element = board[i][j]
				} else {
					element = board[j][i]
				}

				switch element {
				case player:
					sectionPlayerQ++
				case "#":
					emptySpace = coord{i: i, j: j}
				default:
					OpponentQ++
				}
			} //end inner for
			//condiciones para jugar
			if sectionPlayerQ > 0 {
				continue
			}
			if OpponentQ > *maxOpponentPlays {
				*maxOpponentPlays = OpponentQ
				*eligiblePlaceToPlay = emptySpace
			}
		}
	} else if search == "diag1" {
		var OpponentQ int
		var sectionPlayerQ int
		var emptySpace coord

		for i := 0; i < len(board); i++ {
			element := board[i][i]
			switch element {
			case player:
				sectionPlayerQ++
			case "#":
				emptySpace = coord{i: i, j: i}
			default:
				OpponentQ++
			}
		}

		if sectionPlayerQ > 0 {
			return
		}
		if OpponentQ > *maxOpponentPlays {
			*maxOpponentPlays = OpponentQ
			*eligiblePlaceToPlay = emptySpace
		}

	} else if search == "diag2" {
		var OpponentQ int
		var sectionPlayerQ int
		var emptySpace coord
		for i := 0; i < len(board); i++ {
			element := board[i][len(board)-1-i]

			switch element {
			case player:
				sectionPlayerQ++
			case "#":
				emptySpace = coord{i: i, j: len(board) - i}
			default:
				OpponentQ++
			}
		}

		if sectionPlayerQ > 0 {
			return
		}
		if OpponentQ > *maxOpponentPlays {
			*maxOpponentPlays = OpponentQ
			*eligiblePlaceToPlay = emptySpace
		}

	}

}
