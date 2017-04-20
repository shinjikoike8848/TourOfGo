package main

import (
	"testing"
)

const loopbackString = "127.0.0.1"
const googleDNSString = "8.8.8.8"

func TestIPAddr(t *testing.T) {
	loopback := IPAddr{127, 0, 0, 1}
	googleDNS := IPAddr{8, 8, 8, 8}

	if loopback.String() != loopbackString {
		t.Errorf("loopback should be %s", loopbackString)
	}

	if googleDNS.String() != googleDNSString {
		t.Errorf("googleDns should be %s", googleDNSString)
	}
}
