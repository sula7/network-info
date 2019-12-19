package main

import (
	"log"
	"os"

	"github.com/sula7/network-info/config"
	"github.com/sula7/network-info/network"
	"github.com/sula7/network-info/v1"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatalln("unable to get config from env:", err)
	}

	IPs := os.Args[0:]
	if len(IPs) > 1 {
		log.Println("Parsing IP addresses from args")
		for _, ip := range IPs {
			network.GetMacAddress(ip)
		}
	}

	v1.NewApi(cfg.BindPort)
}
