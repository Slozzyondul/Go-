package interfaces1

//IPAddr type implement fmt.Stringer to print the address as a dotted quad.

import (
	"fmt"
)

type IPAddr [4]byte

// Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func IpAddressInterface() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
		"localhost": {127, 0, 0, 2},
		"facebook":  {138, 69, 85, 23},
		"twitter":   {23, 166, 216, 227},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
