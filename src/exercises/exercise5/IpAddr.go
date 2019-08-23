package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IpAddr [4]byte

/* hard way*/
func (ip IpAddr) _String() string {
	s := ""
	for i := 0; i < len(ip); i += 1 {
		part := ip[i]
		s = fmt.Sprintf("%v%v", s, part)
		if i < (len(ip) - 1) {
			s = fmt.Sprintf("%v%v", s, ".")
		}
	}
	return fmt.Sprintf("%v", s)
}

// not so simple
func (ip IpAddr) String() string {
	var ips []string
	for i := 0; i < len(ip); i += 1 {
		ips = append(ips, strconv.Itoa(int(ip[i])))
	}
	return strings.Join(ips, ".")
}

func main() {
	var hosts = make(map[string]IpAddr)
	hosts["loopback"] = IpAddr{127, 0, 0, 1}
	hosts["googleDNS"] = IpAddr{8, 8, 8, 8}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
