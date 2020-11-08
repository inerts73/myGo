package main

import "fmt"

// IPAddr a byte slice for IP
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipAddr IPAddr) String() string {
	result := ""
	for _, v := range ipAddr {
		result += fmt.Sprintf("%v", v) + "."
	}
	return result[:len(result)-1]
}


func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	
}