package main 

import (
	"net"
	"io"
	"log"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "joescatcam.website:80")
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}

	defer dst.Close()

	// Run in goroutine to prevent io.Copy from blocking
	go func() {
		// ensure data from the inbound connection is copied to the joescatcam.website
		// io.Copy(dst, src)

		// copy from the joesproxy to joescatcam
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	// Copy our destination's output back to our source
	// copy from joescatcam to joesproxy
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}

}


func main() {
	// Listen on local port 80
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		// for each request, handle once
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}
