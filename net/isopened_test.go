package net

import "testing"

type test struct {
	host string
	port interface{}
	res  bool
}

var tt = []test{
	{
		"www.seznam.cz",
		80,
		true,
	},
	{
		"www.seznam.cz",
		1,
		false,
	},
	{
		"www.seznam.cz",
		"80",
		true,
	},
	{
		"www.seznam.cz",
		"ee",
		false,
	},
	{
		"www.seznam.cz",
		1.2,
		false,
	},
}

func TestIsOpened(t *testing.T) {

	for _, tc := range tt {
		r, _ := IsOpened(tc.host, tc.port)
		if r != tc.res {
			t.Errorf("host: %s, port: %d, expected: %t, get: %t\n", tc.host, tc.port, tc.res, r)
		}

	}
}
