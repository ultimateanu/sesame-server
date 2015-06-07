package main

import (
	"errors"
	"fmt"
	"net"
)

func getLocalIp() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok {
			return nil, err
		}
		ip4 := ipnet.IP.To4()
		if ip4 == nil || ip4[0] == 127 {
			continue
		}
		return ip4, nil
	}
	return nil, errors.New("No local ip found")
}

func printLocalIp() {
	localIp, err := getLocalIp()
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Printf("Serving at http://%v:%d\n", localIp, port)
	}
}
