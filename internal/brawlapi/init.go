package brawlapi

import (
	"bufio"
	"log"
	"os"
)

var token string

const officialApiURI = "https://api.brawlstars.com"

func init() {
	file, err := os.Open("./secret/brawl-api.token")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	token = scanner.Text()
}
