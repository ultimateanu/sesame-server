package server

import (
	"net"
	"os"
	"unicode"
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

func UrlSafe(s string) string {
	r := make([]rune, len(s))
	for i, c := range s {
		r[i] = '-'
		if unicode.IsDigit(c) || unicode.IsLower(c) || unicode.IsUpper(c) || c == '.' || c == '_' || c == '~' {
			r[i] = c
		}
	}
	return string(r)
}
