package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func getPeers(trnt map[string]interface{}) (string, error) {
	baseURL := trnt["announce"].(string)

	fmt.Println(baseURL)

	info, ok := trnt["info"].(map[string]interface{})

	fmt.Println(info)

	if info == nil || !ok {
		return "", errors.New("No info section")
	}

	u, err := url.Parse(baseURL)

	if err != nil {
		return "", errors.New("could not parse URL")
	}

	infoHash, err := getHash(info)

	if err != nil {
		return "", err
	}

	// Add query parameters
	queryParams := url.Values{}
	queryParams.Add("info_hash", string(infoHash))
	queryParams.Add("peer_id", "aaaaaaaaaaaaaaaaaaaa")
	queryParams.Add("port", "6881")
	queryParams.Add("uploaded", "0")
	queryParams.Add("downloaded", "0")
	queryParams.Add("left", info["length"].(string))
	queryParams.Add("compact", "1")
	u.RawQuery = queryParams.Encode() // Attach the query params to the URL

	// Create a new GET request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make the request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read the response body: %v", err)
		}

		// fmt.Println(string(body))
		return string(body), nil
	}
	return "", nil
	// Read and print the response body

}
