package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	name := flag.String("h", "", "Hostname to find the IP of.")
	flag.Parse()

	if *name == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("Looking for IP addreses:")
	ips, err := net.LookupHost(*name)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
