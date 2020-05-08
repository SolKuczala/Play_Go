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

var MYtURN bool

func main() {
	baseURL := getUserParams()
	//va a empezar la jugada pidiendo un board,asumo que empieza este player
	//me imagino una funcion que si le llega true del toss coin hace get del board
	MYtURN = true
	if MYtURN == true {
		getBoard(baseURL, 3)
		//mando juego
		playRandom(baseURL)
		//ya no es mas mi turno
		MYtURN = false
		//espero
		time.Sleep(7 * time.Second)
		//pido status
		respStruct, err := getStatus(baseURL)
		if err != nil {
			fmt.Println(err)
		} else {
			for respStruct.Lastplayer == "X" {
				time.Sleep(10 * time.Second)
				_, err := getStatus(baseURL)
				if err != nil {
					fmt.Println(err)
				}
			} //una vez que rompe
			playRandom(baseURL)
		}
		//		si la resp distinto de

	}
	getStatus(baseURL)

}

func getUserParams() string {
	var port = flag.Int("port", 8080, "port number")
	var ip = flag.String("ip", "127.0.0.1", "ip")
	flag.Parse()
	return fmt.Sprintf("http://%s:%d", *ip, *port)
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

//randomPlay
func playRandom(baseURL string) error {
	fmt.Println("playing...")
	rand.Seed(time.Now().UnixNano())
	randomNum := randomNum(0, 2)
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
	//fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
	return nil
}

//opposite game
func toNotLoose() {

}

//to actually want to win
func active() {

}
func randomNum(min int, max int) int {
	return rand.Intn(max-min) + min
}

type Body struct {
	Board      [][]string `json: "board"`
	Lastplayer string     `json: "last-player"`
	Status     string     `json: "status"`
}

/*Gets status of the game, return error from standard packages*/
func getStatus(baseURL string) (Body, error) {
	fmt.Println("getting status...")
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
