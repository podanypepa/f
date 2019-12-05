package main

import (
	"fmt"
	"os"

	"github.com/podanypepa/f/net"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("1. args = host_name, 2. args = port number")
		os.Exit(0)
	}

	o, err := net.IsOpened(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Printf("%t (%v)\n", o, err)
	} else {
		fmt.Printf("%t\n", o)
	}

}
