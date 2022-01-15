package machineip

import (
	"net"
	"strings"
)

// Returns the host ip.
func IP() string {
	conn, err := net.Dial("udp", "1.1.1.1:80")
	if err != nil {
		return ""
	}

	defer conn.Close()

	// Given an ip like 192.162.7.20:44248, we split it in two parts [192.162.7.20, 44248].
	addressPieces := strings.Split(conn.LocalAddr().String(), ":")

	// And then return the ip without the port.
	return addressPieces[0]
}
