package net

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// IsOpened test if is on host is opened specifi port
func IsOpened(host string, port interface{}) (o bool, e error) {

	var p int

	switch port := port.(type) {
	case int:
		p = port
	case string:
		p, e = strconv.Atoi(port)
		if e != nil {
			return false, e
		}
	default:
		return false, fmt.Errorf("usuperted type of port: %+v", port)
	}

	timeout := 5 * time.Second
	target := fmt.Sprintf("%s:%d", host, p)

	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return false, err
	}

	_ = conn.Close()
	return true, nil
}
