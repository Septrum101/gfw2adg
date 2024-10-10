package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
)

func handleGfwList(dns string, gfwUrl string) (string, error) {
	log.Print("Fetching gfwlist..")
	resp, err := resty.New().R().Get(gfwUrl)
	if err != nil {
		return "", err
	}

	rawList := strings.Split(resp.String(), "\n")
	log.Printf("converting %d domains..", len(rawList))

	var mergedDomains string
	count := 0
	for _, val := range rawList {
		if val != "" {
			mergedDomains += val + "/"
			count++
		}
	}
	log.Printf("Valid domains: %d", count)

	return fmt.Sprintf("[/%s]%s", mergedDomains, dns), nil
}
