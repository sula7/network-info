package network

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/mostlygeek/arp"
	"github.com/sparrc/go-ping"
)

func GetIPInfo(ip string) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		log.Printf("unable to ping %s IP address: %v\n", ip, err)
		return
	}

	pinger.SetPrivileged(true)
	pinger.Count = 10
	pinger.Interval = time.Second
	pinger.Timeout = time.Second * 2
	pinger.Run()

	hostname, err := net.LookupAddr(ip)

	ttl, err := GetTtl(ip)
	if err != nil {
		log.Printf("unable to get ttl of %s: %v\n", ip, err)
	}

	var operatingSystem string
	switch ttl {
	case "64":
		operatingSystem = "Linux OS"
	case "128":
		operatingSystem = "Windows OS"
	case "254":
		operatingSystem = "Cisco"
	case "255":
		operatingSystem = "FreeBSD/Unix OS"
	default:
		operatingSystem = "OS not detected"
	}

	fmt.Printf("%s\t%s\t%v\t%s\n", ip, arp.Search(ip), hostname, operatingSystem)
}

func GetTtl(ip string) (string, error) {
	var arg string
	var out []byte
	var err error
	var ttl string

	switch runtime.GOOS {
	case "windows":
		arg = fmt.Sprintf(`ping -n 1 %s | FIND "TTL="`, ip)
		out, err = exec.Command("cmd", "/c", arg).Output()
		if err != nil {
			return "", err
		}

		ttl = string(out)
		log.Println("TTL before strings", ttl)

		ttl = strings.TrimLeft(ttl, "TTL=")
		ttl = strings.Replace(ttl, "TTL=", "", -1)
		ttl = strings.TrimSuffix(ttl, "\n") //need to trim cuz contains "\n" as default
		fmt.Println("TTL after strings", ttl)

	case "linux":
		arg = fmt.Sprintf(`ping -c 1 %s | awk -F'[ =]' '/ttl/ {print $8}'`, ip)
		out, err = exec.Command("bash", "-c", arg).Output()
		if err != nil {
			return "", err
		}

		ttl = string(out)
		ttl = strings.TrimSuffix(ttl, "\n") //need to trim cuz contains "\n" as default

	default:
		arg = fmt.Sprintf(`ping -c 1 %s | awk -F'[ =]' '/ttl/ {print $8}'`, ip)
		out, err = exec.Command("bash", "-c", arg).Output()
		if err != nil {
			return "", err
		}

		ttl = string(out)
		ttl = strings.TrimSuffix(ttl, "\n") //need to trim cuz contains "\n" as default
	}

	return ttl, nil
}
