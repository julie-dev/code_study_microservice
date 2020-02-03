package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"code_study_microservice/http-server-client/entities"
)

const port = 8080

func main() {
	c := CreateClient()
	resp, err := GetResponse(c)
	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(-1)
	}

	fmt.Println(resp.Message)
}

func CreateClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 5,
		},
		Timeout: 5 * time.Second,
	}

	return client
}

func GetResponse(client *http.Client) (resp *entities.ResponseData, err error) {
	reqData := entities.RequestData{Name:"julie"}
	data, _ := json.Marshal(reqData)

	url := fmt.Sprintf("http://localhost:%v/helloworld", port)
	request, _ := http.NewRequest("GET", url, bytes.NewBuffer(data))

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed with status: %v", response.Status)
	}

	decoder := json.NewDecoder(response.Body)
	var result entities.ResponseData
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
