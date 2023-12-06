package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func FetchData(url string, data interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Check if the request was successful (HTTP status code 200)
	if response.StatusCode != http.StatusOK {
		fmt.Println("API request failed with status:", response.Status)
		return errors.New("API request failed with status:" + response.Status)
	}

	// Decode the JSON response into a slice of Product structs
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return err
	}

	return nil
}
