package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	path := flag.String("f", "/dev/sda", "The file to inspect the first 512 bytes of.")
	outFile := flag.String("o", "", "Write output to this file.")
	flag.Parse()

	var out io.Writer
	var err error
	if *outFile == "" {
		out = os.Stdout
	} else {
		out, err = os.Create(*outFile)
		if err != nil {
			log.Fatal("Could not create file: " + err.Error())
		}
		defer out.(*os.File).Close()
	}

	log.Println("[+] Reading boot sectory of " + *path)
	file, err := os.Open(*path)
	if err != nil {
		log.Fatal("Count not open file: " + err.Error())
	}
	defer file.Close()

	bytes := make([]byte, 512)
	bytesRead, err := io.ReadFull(file, bytes)
	if err != nil {
		log.Fatalf("Error reading 512 bytes (got %d) from %s: %s", bytesRead, *path, err.Error())
	}

	log.Printf("Bytes read: %d\n\n", bytesRead)
	log.Printf("Data as decimal:\n\t%d\n\n", bytes)
	log.Printf("Data as hex:\n\t%x\n\n", bytes)
	log.Printf("Data as string:\n\t%s\n", bytes)

	if *outFile != "" {
		fmt.Fprintf(out, "Bytes read: %d\n\n", bytesRead)
		fmt.Fprintf(out, "Data as decimal:\n\t%d\n\n", bytes)
		fmt.Fprintf(out, "Data as hex:\n\t%x\n\n", bytes)
		fmt.Fprintf(out, "Data as string:\n\t%s\n", bytes)
	}
}
