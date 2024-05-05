package brawlapi

import (
	"encoding/json"
	"io"
	"log"
)

func GetPlayerInfo(tag string) (PlayerInfo, error) {
	var playerInfo PlayerInfo
	endpoint := "/v1/players/%23" + tag
	res, err := sendRequest(endpoint)
	if err != nil {
		return playerInfo, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &playerInfo)

	return playerInfo, nil
}
