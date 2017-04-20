package main

import "fmt"

// IPAddr express bytes of separated by comma.
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ipa IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipa[0], ipa[1], ipa[2], ipa[3])
}
