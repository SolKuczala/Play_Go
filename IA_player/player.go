package main

//esta en otro docker?
import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	baseURL := getUserParams()
	//va a empezar la jugada pidiendo un board
	getBoard(baseURL)
	playreq(baseURL)
	getStatus(baseURL)
	//read the board and depending where the opposite played you played

}

func getUserParams() string {
	var port = flag.Int("port", 8080, "port number")
	var ip = flag.String("ip", "127.0.0.1", "ip")
	flag.Parse()
	return fmt.Sprintf("http://%s:%d", *ip, *port)
}

func getBoard(baseURL string) error {
	fmt.Println("getting board...")
	rand.Seed(time.Now().UnixNano())
	randomNum := random(3, 9)
	resp, err := http.Get(baseURL + "/create-board/" + strconv.Itoa(randomNum))
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

func playreq(baseURL string) error {
	fmt.Println("playing...")
	rand.Seed(time.Now().UnixNano())
	randomNum := random(0, 2)
	client := &http.Client{}
	// Create request
	req, err := http.NewRequest("PUT", baseURL+"/send-play/X/"+strconv.Itoa(randomNum)+"/"+strconv.Itoa(randomNum), nil) //+strconv.Itoa(randomNum)+"/"+strconv.Itoa(randomNum)/ or url.Values{"player": {"X"}, "row": {"1"}, "column": {"0"}
	if err != nil {
		fmt.Println(err)
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
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
	return nil
}

func getStatus(baseURL string) error {
	fmt.Println("getting status...")
	resp, err := http.Get(baseURL + "/status")
	if err != nil {
		return fmt.Errorf("%g", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%g", err)
	}

	type Body struct {
		Board      [][]string `json: "board"`
		Lastplayer string     `json: "last-player"`
		Status     string     `json: "status"`
	}

	var bodyContent Body
	err = json.Unmarshal(body, &bodyContent)
	if err != nil {
		fmt.Println("error:", err)
		return err
	} else {
		fmt.Printf("%+v", bodyContent)
	}
	//log.Println(string(body))
	return nil
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
