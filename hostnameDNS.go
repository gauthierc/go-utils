package main

import (
	"flag"
	"fmt"
	"net"
	"regexp"
)

func revers(ip4 string, fqdn bool, short bool, domain bool, ip bool) {
	Hostname, err := net.LookupAddr(ip4)
	if err == nil {
		switch {
		case domain:
			regE := regexp.MustCompile("^[^.]*[.](.*)[.]$")
			fmt.Printf("%s\n", regE.FindStringSubmatch(Hostname[0])[1])
		case fqdn:
			regE := regexp.MustCompile("^(.*)[.]$")
			fmt.Printf("%s\n", regE.FindStringSubmatch(Hostname[0])[1])
		case ip:
			fmt.Printf("%s\n", ip4)
		default:
			regE := regexp.MustCompile("^([^.]*)")
			fmt.Printf("%s\n", regE.FindStringSubmatch(Hostname[0])[1])
		}
	}
}

func main() {

	fqdnPtr := flag.Bool("f", false, "long host name (FQDN)")
	shortPtr := flag.Bool("s", false, "short host name (default)")
	domainPtr := flag.Bool("d", false, "DNS domain name")
	ipPtr := flag.Bool("i", false, "addresses for the host name")
	flag.Parse()

	regS := regexp.MustCompile("[0-9.]*")
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		if addrs, err := inter.Addrs(); err == nil {
			for _, addr := range addrs {
				monip := regS.FindString(addr.String())
				if monip != "127.0.0.1" {
					revers(monip, *fqdnPtr, *shortPtr, *domainPtr, *ipPtr)
				}
			}
		}
	}

}
