package main

import (
	"fmt"
	"log"

	"github.com/podanypepa/f/net"
)

func main() {

	host, ip, err := net.GetIP()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s, %s\n", ip, host)

}
