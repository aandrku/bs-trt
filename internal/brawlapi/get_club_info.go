package brawlapi

import (
	"encoding/json"
	"io"
	"log"
)

func GetClubInfo(tag string) (ClubInfo, error) {
	var clubInfo ClubInfo
	endpoint := "/v1/clubs/%23" + tag
	res, err := sendRequest(endpoint)
	if err != nil {
		return clubInfo, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &clubInfo)

	return clubInfo, nil
}
