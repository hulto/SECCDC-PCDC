package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	port := flag.Int("p", 1025, "Port to listen on.")
	flag.Parse()

	// Create listener
	// Listens on the given port on all available IP addresses
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Error listening on port %d.\n%s\n", *port, err.Error())
	}
	log.Printf("Listening on port %d.\n", *port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connecion.", err)
		}
		go func(con net.Conn) {
			log.Printf("Received connection from %s.\n", conn.RemoteAddr())
			buf := make([]byte, 1024)
			nbyte, err := con.Read(buf)
			if err != nil {
				log.Println("Error reading from connection.", err)
			}
			// Dummy response
			conn.Write([]byte("An unknown error has occured."))
			// Printing what was sent
			trimmed := bytes.TrimRight(buf, "\x00")
			log.Printf("Read %d bytes from %s.\n%s\n", nbyte, conn.RemoteAddr(), trimmed)
			con.Close()
		}(conn)
	}
}
