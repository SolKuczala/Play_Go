package main

//esta en otro docker?
import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	//va a empezar la jugada pidiendo un board
	getBoard()
	//time.Sleep(2 * time.Second)
	playreq()
}

func getBoard() {
	rand.Seed(time.Now().UnixNano())
	randomNum := random(3, 9)
	resp, err := http.Get("http://127.0.0.1:8080/create-board/" + strconv.Itoa(randomNum))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func playreq() {
	//rand.Seed(time.Now().UnixNano())
	//randomNum := random(3, 9)
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("PUT", "http://127.0.0.1:8080/send-play/X/0/1", nil) //+strconv.Itoa(randomNum)+"/"+strconv.Itoa(randomNum)/ or url.Values{"player": {"X"}, "row": {"1"}, "column": {"0"}
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	//	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
