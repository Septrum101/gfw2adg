package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
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

	b := bytes.NewBufferString(strings.Join(conf.DefaultDNS, "\n"))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("[/%s/]%s\n", strings.Join(conf.CustomProxyDomain, "/"), conf.ProxyDNS))

	gfw, err := handleGfwList(conf.ProxyDNS, conf.GFWListURL)
	if err != nil {
		panic(err)
	}
	b.WriteString(gfw)

	log.Print("Write to custom_dns.conf")
	err = os.WriteFile("custom_dns.conf", b.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
