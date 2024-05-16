package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Message represents a single message in the conversation
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Profile represents the user's profile
type Profile struct {
	Name              string `json:"name"`
	Age               string `json:"age"`
	Gender            string `json:"gender"`
	Job               string `json:"job"`
	FamilyBackground  string `json:"familyBackground"`
	FriendsBackground string `json:"friendsBackground"`
	AnythingElse      string `json:"anythingElse"`
}

// Payload is the top-level structure of the JSON payload
type Payload struct {
	Messages []Message `json:"messages"`
	Profile  Profile   `json:"profile"`
}

func Ask_AI(url string, data *Payload, client *http.Client) (string, error) {
	if data == nil {
		return "", fmt.Errorf("data cannot be nil")
	}

	// convert data to json
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error marshalling data: %v", err)
	}

	//create new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	//create new header
	req.Header.Set("Content-Type", "application/json")

	//initiate client request and get response
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	//response status code check
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("received non-2XX response status: %d %s", resp.StatusCode, resp.Status)
	}

	//response to readable data
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	// err = json.Unmarshal(responseBody, &jsonData)
	// if err != nil {
	// 	return "", fmt.Errorf("Error parsing response to json:", err, "\n")
	// }

	return string(responseBody), nil

}
