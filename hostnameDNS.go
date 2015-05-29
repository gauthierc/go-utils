package main

import (
	"fmt"
	"net"
	"regexp"
)

func revers(ip4 string) {
	regE := regexp.MustCompile("[-a-z0-9]*")
	Hostname, err := net.LookupAddr(ip4)
	if err == nil {
		fmt.Printf("%s\n", regE.FindString(Hostname[0]))
	}
}

func main() {

	regS := regexp.MustCompile("[0-9.]*")
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		if addrs, err := inter.Addrs(); err == nil {
			for _, addr := range addrs {
				monip := regS.FindString(addr.String())
				if monip != "127.0.0.1" {
					revers(monip)
				}
			}
		}
	}

}
