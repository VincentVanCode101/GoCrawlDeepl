package network

import (
	"errors"
	"net"
	"testing"
	"time"
)

// Mock function for DialTimeout
func mockDialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	if network == "tcp" && address == "8.8.8.8:53" {
		return nil, errors.New("mocked no connection")
	}
	return &net.TCPConn{}, nil
}

func TestCheckInternetConnection(t *testing.T) {
	originalDialTimeout := dialTimeout
	defer func() { dialTimeout = originalDialTimeout }()

	dialTimeout = mockDialTimeout

	err := CheckInternetConnection()

	if err == nil {
		t.Fatalf("Expected an error, but got nil")
	}

	expectedErr := "no internet connection: unable to connect to Google's DNS server"
	if err.Error() != expectedErr {
		t.Fatalf("Expected error %q, but got %q", expectedErr, err.Error())
	}
}

func TestCheckInternetConnection_ConnectionAvailable(t *testing.T) {

	originalDialTimeout := dialTimeout
	defer func() { dialTimeout = originalDialTimeout }()

	dialTimeout = func(network, address string, timeout time.Duration) (net.Conn, error) {
		return &net.TCPConn{}, nil
	}

	err := CheckInternetConnection()
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
}
