package brawlapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func sendRequest(endpoint string) (http.Response, error) {
	url := officialApiURI + endpoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{
		Timeout: time.Second * 10,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		var clientError ClientError
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(body, &clientError)
		err = clientError
		return *res, err
	}

	return *res, nil

}
