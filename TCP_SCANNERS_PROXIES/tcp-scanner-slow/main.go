package main

import (
	"fmt"
	"net"
)

func main() {
	// net.Dial(network, address string)
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)

		conn, err := net.Dial("tcp", address)

		if err != nil {
			fmt.Printf("%d close\n", i)
			continue // fail to connect
		}

		// connected
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
