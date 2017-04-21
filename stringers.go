package main

import "fmt"

// IPAddr express bytes of separated by comma.
type IPAddr [4]byte

// String return formatted ipaddress string like *.*.*.*
func (ipa IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipa[0], ipa[1], ipa[2], ipa[3])
}
