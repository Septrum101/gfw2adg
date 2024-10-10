package main

import (
	"bytes"
	"log"
	"os"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"
)

func main() {
	f, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	conf := new(Config)
	if err := yaml.NewDecoder(f).Decode(conf); err != nil {
		panic(err)
	}

	slices.Sort(conf.DefaultDNS)
	b := bytes.NewBufferString(strings.Join(conf.DefaultDNS, "\n"))
	b.WriteString("\n")

	gfwDomainMap := make(map[string]bool)
	gfw, err := getGFWList(conf.GFWListURL)
	if err != nil {
		panic(err)
	}
	for _, val := range gfw {
		gfwDomainMap[val] = true
	}

	if len(conf.CustomProxyDomain) != 0 {
		log.Printf("Get %d custom domains", len(conf.CustomProxyDomain))
		customDomainMap := make(map[string]bool)
		for _, val := range conf.CustomProxyDomain {
			if _, ok := gfwDomainMap[val]; !ok {
				customDomainMap[val] = true
			} else {
				log.Printf("Found duplicate domain: %s\n", val)
			}
		}
		if customStr, err := handleDomains(conf.ProxyDNS, customDomainMap); err != nil {
			log.Print(err)
		} else {
			b.WriteString(customStr)
			b.WriteString("\n")
		}
	}

	gfwStr, err := handleDomains(conf.ProxyDNS, gfwDomainMap)
	if err != nil {
		panic(err)
	}
	b.WriteString(gfwStr)
	b.WriteString("\n")

	log.Print("Write to custom_dns.conf")
	err = os.WriteFile("custom_dns.conf", b.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
