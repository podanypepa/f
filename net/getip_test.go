package net

import "testing"

func TestGetIP(t *testing.T) {

	host, ip, err := GetIP()
	if err != nil || ip == "" || host == "" {
		t.Error(err)
	}
}
