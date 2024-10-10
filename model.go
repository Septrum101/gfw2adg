package main

type Config struct {
	ProxyDNS          string   `yaml:"ProxyDNS"`
	GFWListURL        string   `yaml:"GFWListURL"`
	DefaultDNS        []string `yaml:"DefaultDNS"`
	CustomProxyDomain []string `yaml:"CustomProxyDomain"`
}
