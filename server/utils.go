package server

import (
	"net"
	"os"
)

func GetLocalIp() ([]string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	addrs, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}

	return addrs, nil
}
