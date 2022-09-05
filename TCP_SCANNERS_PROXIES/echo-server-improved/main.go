package main

import (
	"fmt"
	"io"
	"net"
	"bufio"
)

func echo(conn net.Conn) {
	defer conn.Close()

	// buffered reader 
	reader := bufio.NewReader(conn)
	// read string
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	log.Printf("Read %d bytes:%s", len(s), s)

	log.Println("Writing data")
	// buffered writer 
	writer := bufio.NewWriter(conn)
	
	// write the data into socket
	if _, err := writer.WriteStrings(s); err != nil {
		log.Fatalln("Unable to write data")
	}
	
	// flush write all the data to the underlying writer
	write.Flush()
}

// or even more simpler
func echo(conn net.Conn) {
	defer conn.Close()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}
