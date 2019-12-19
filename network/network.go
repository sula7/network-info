package network

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"net"
	"time"

	"github.com/mostlygeek/arp"
	"github.com/sparrc/go-ping"
)

func GetMacAddress(ip string) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		log.Errorf("unable to ping %s IP address: %v", ip, err)
		return
	}

	pinger.SetPrivileged(true)
	pinger.Count = 10
	pinger.Interval = time.Second
	pinger.Timeout = time.Second * 2
	pinger.Run()

	hostname, err := net.LookupAddr(ip)
	fmt.Printf("%s\t%s\t%v\n", ip, arp.Search(ip), hostname)
}
