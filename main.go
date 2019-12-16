package main

import (
	"log"

	"github.com/sula7/network-info/config"
)

func main() {
	_, err := config.Get()
	if err != nil {
		log.Fatalln("unable to get config from env:", err)
	}
}
