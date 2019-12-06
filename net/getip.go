package net

import (
	"net"
	"os"
	"strings"
)

// GetIP return hostname, IP and error if is some there
func GetIP() (host, ip string, err error) {
	if host, err = os.Hostname(); err != nil {
		return "", "", err
	}

	addr, err := net.InterfaceAddrs()
	if err != nil {
		return "", "", err
	}
	for _, a := range addr {
		if !strings.Contains(a.String(), ":") && !strings.Contains(a.String(), "127.") {
			ip = strings.Split(a.String(), "/")[0]
			break
		}
	}
	return host, ip, nil
}
