package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	today := time.Now().Format("2006-01-02")
	url := "https://api.fitbit.com/1/user/-/activities/date/" + today + ".json"

	// Your token
	token := ""
	bearer := "Bearer " + token

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("Error new requst: ", err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error response: ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading: ", err)
	}

	log.Println(string([]byte(body)))
}
