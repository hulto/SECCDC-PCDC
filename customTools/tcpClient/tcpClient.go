package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var err error
var conn net.Conn

func main() {
	port := flag.Int("p", -1, "Port to connect to.")
	ip := flag.String("n", "", "IP address to connect to.")
	flag.Parse()

	if *ip == "" || *port < 1 {
		flag.Usage()
		os.Exit(1)
	}

	// Connecting
	addr := fmt.Sprintf("%s:%d", *ip, *port)
	conn, err = net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		log.Fatal("Could not connect: " + err.Error())
	}

	output := make(chan string, 1)
	input := make(chan string, 1)

	go showOutput(output)
	go getInput(input)

	for {
		go func(connn net.Conn) {
			buf := make([]byte, 10000)
			bufio.NewReader(connn).Read(buf)
			//trim := bytes.TrimRight(buf, "\x00")
			output <- "\n" + string(buf) + "\n"
		}(conn)

		text := <-input
		fmt.Fprint(conn, text+"\n")
	}
}

func showOutput(out chan string) {
	for {
		fmt.Print(<-out)
	}
}

func getInput(in chan string) {
	for {
		buf := make([]byte, 1024)
		bufio.NewReader(os.Stdin).Read(buf)
		msg := string(bytes.TrimRight(buf, "\x00"))
		if msg == "\n" {
			continue
		}
		in <- msg
	}
}
