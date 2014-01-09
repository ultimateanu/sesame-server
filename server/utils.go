package server

import (
	"html"
	"net"
	"os"
	"strings"
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

func Escape(s string) string {
	s = html.EscapeString(s)
	return strings.Replace(s, " ", "&nbsp;", -1)
}
