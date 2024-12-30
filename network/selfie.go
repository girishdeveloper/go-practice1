package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
)

func main() {
	if runtime.GOOS == "windows" {
		panic("This script does not run on windows")
	}
	var userArgs = ""
	var cname = ""
	var addrs []string
	fmt.Printf("IP address length in IP4 is %d and IP address length in IP6 is %d\n", net.IPv4len, net.IPv6len)
	fmt.Println(os.Args)
	userArgs = strings.Join(os.Args[1:], " ")
	addrs, err := net.LookupHost(userArgs)
	if err != nil {
		panic(err)
	}
	fmt.Println(addrs)
	cname, err1 := net.LookupCNAME(userArgs)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(cname)
	fmt.Printf("CIDR for IP4 corresponding to /31 subnet is %v\n", net.CIDRMask(31, 32))
	fmt.Printf("CIDR for IP6 corresponding to /64 subnet is %v\n", net.CIDRMask(64, 128))
}
