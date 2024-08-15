package network

import (
	"errors"
	"net"
	"time"
)

var dialTimeout = net.DialTimeout

// CheckInternetConnection verifies if the internet connection is available.
func CheckInternetConnection() error {
	const googleDNS = "8.8.8.8:53"
	const timeout = 5 * time.Second

	conn, err := dialTimeout("tcp", googleDNS, timeout)
	if err != nil {
		return errors.New("no internet connection: unable to connect to Google's DNS server")
	}
	defer conn.Close()

	return nil
}
