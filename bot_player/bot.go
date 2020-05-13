package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/SolKuczala/tic-tac-go/bot_player/strategy"
)

var waitTime = 1 * time.Millisecond

//var defaultTriesTreshold = 10

func main() {
	baseURL, strategyName, myPlayerChar, firstToPlay := getUserParams()
	//playchar default X, sino O
	fmt.Printf("Strategy: %s\n", strategyName)
	if firstToPlay {
		if err := getBoard(baseURL, 3); err != nil {
			panic("Can't create board :(")
		}
	}
	selectedStrategy := strategy.StrategiesMap[strategyName]
	ongoingGame := true
	for ongoingGame {
		//pido status
		response, err := getStatus(baseURL)
		if err != nil {
			fmt.Println("hay error de get status")
		} else {
			if response.Winners > 0 {
				break
			}
		}

		myTurn := false
		if response.Lastplayer == "#" {
			myTurn = firstToPlay
		} else if response.Lastplayer != myPlayerChar {
			myTurn = true
		}

		if myTurn {
			selectedStrategy.Play(baseURL, myPlayerChar, 1, response.Board)
		} else {
			time.Sleep(waitTime)
			continue
		}

	}
	fmt.Println("End of game")
}

func getUserParams() (string, string, string, bool) {
	var port = flag.Int("port", 8080, "port number")
	var ip = flag.String("ip", "127.0.0.1", "ip")
	var strategy = flag.String("strategy", "random", "strategy")
	var player = flag.String("player", "X", "player")
	var playfirst = flag.Bool("playfirst", false, "playfirst")
	flag.Parse()
	return fmt.Sprintf("http://%s:%d", *ip, *port), *strategy, *player, *playfirst
}

func getBoard(baseURL string, size int) error {
	fmt.Println("getting board...")
	resp, err := http.Get(baseURL + "/create-board/" + strconv.Itoa(size))
	if err != nil {
		return fmt.Errorf("%g", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%g", err)
	}
	log.Println(string(body))
	return nil
}

//opposite game
func toNotLoose() {

}

//to actually want to win
func active() {

}

type Body struct {
	Board      [][]string `json: "board"`
	Status     string     `json: "status"`
	Lastplayer string     `json: "lastPlayer"`
	Winners    int        `json: "winners"`
}

/*Gets status of the game, return error from standard packages*/
func getStatus(baseURL string) (Body, error) {
	//fmt.Println("getting status...")
	var bodyContent Body
	resp, err := http.Get(baseURL + "/status")
	if err != nil {
		return bodyContent, fmt.Errorf("%g", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bodyContent, fmt.Errorf("%g", err)
	}

	err = json.Unmarshal(body, &bodyContent)
	if err != nil {
		fmt.Println("error:", err)
		return bodyContent, err
	}
	return bodyContent, nil

	//log.Println(string(body))    fmt.Printf("%+v",
}
