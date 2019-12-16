package main

import (
	"log"

	"github.com/sula7/network-info/config"
	"github.com/sula7/network-info/v1"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatalln("unable to get config from env:", err)
	}

	r := v1.NewApi()

	err = v1.Start(r, cfg.BindPort)
	if err != nil {
		log.Fatalln("unable to start HTTP server:", err)
	}
}
