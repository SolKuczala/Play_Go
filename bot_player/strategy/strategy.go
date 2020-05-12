package strategy

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

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
		Gen: func(board [][]string, player string) (int, int, error) {
			for i := 0; i < len(board); i++ {
				for j := 0; j < len(board); j++ {
					if board[i][j] == "#" {
						return j, i, nil
					}
				}
			}
			return 0, 0, errors.New("chau")
		},
	},
}

type PlayStrategy struct {
	Gen func(board [][]string, player string) (int, int, error)
}

func (ps PlayStrategy) Play(baseURL, player string, tries int, board [][]string) error {
	isOk := false
	for !isOk && tries > 0 {
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
		return errors.New("To many tries")
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
