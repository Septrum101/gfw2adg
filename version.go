package main

import (
	"fmt"
)

var (
	appName = "gfw2adg"
	version = "dev"
	date    = "unknown"
)

func init() {
	fmt.Printf("%s %s, built at %s\n", appName, version, date)
}
