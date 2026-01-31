package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func getJSON[T any](url string) (*T, error) {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating new GET request")
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing GET request")
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return nil, err
	}

	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("Error unmarshalling data")
		return nil, err
	}

	return &result, nil
}
