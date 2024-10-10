package main

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/go-resty/resty/v2"
)

func getGFWList(gfwListUrl string) ([]string, error) {
	resp, err := resty.New().R().Get(gfwListUrl)
	if err != nil {
		return nil, err
	}

	var gfw []string
	for _, val := range strings.Split(strings.TrimSpace(resp.String()), "\n") {
		domain := strings.TrimSpace(val)
		if domain != "" {
			gfw = append(gfw, domain)
		}
	}
	log.Printf("Get %d gfw domains", len(gfw))

	if len(gfw) == 0 {
		return nil, errors.New("no gfw domains found")
	}

	return gfw, nil
}

func handleDomains(dns string, domains map[string]bool) (string, error) {
	if len(domains) == 0 {
		return "", fmt.Errorf("no domains found")
	}

	var domainList []string
	for val := range domains {
		if val != "" {
			domainList = append(domainList, val)
		}
	}
	slices.Sort(domainList)
	log.Printf("Converted %d domains to AdGuard format", len(domainList))

	return fmt.Sprintf("[/%s/]%s", strings.Join(domainList, "/"), dns), nil
}
