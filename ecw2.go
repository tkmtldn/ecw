package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var URL = "https://status.ecwid.com/api/v2/status.json"

type Status struct {
	Description string `json:"description"`
	Indicator   string `json:"indicator"`
}
type Page struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	TimeZone  string `json:"time_zone"`
	UpdatedAt string `json:"updated_at"`
	Url       string `json:"url"`
}

type Statuses struct {
	Page   `json:"page"`
	Status `json:"status"`
}

func main() {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
		return
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	myStatuses := new(Statuses)
	json.Unmarshal(body, &myStatuses)

	fmt.Println(myStatuses.Description)
}
