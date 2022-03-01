package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	ip := flag.String("i", "", "IP address to convert to its hostname.")
	flag.Parse()

	if *ip == "" {
		flag.Usage()
		os.Exit(1)
	}

	validIp := net.ParseIP(*ip)
	if validIp == nil {
		log.Fatalln("The provided IP is not valid.")
	}

	fmt.Println("Found hostnames:")
	hostnames, err := net.LookupAddr(validIp.String())
	if err != nil {
		log.Fatal(err)
	}
	for _, hostname := range hostnames {
		fmt.Println(hostname)
	}
}
