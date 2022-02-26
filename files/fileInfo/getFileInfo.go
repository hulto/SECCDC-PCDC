package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var fileInfo os.FileInfo
	var err error

	fileName := flag.String("f", "", "File to analyze.")
	flag.Parse()
	if *fileName == "" {
		fmt.Println("No file specified!")
		flag.Usage()
		os.Exit(1)
	}

	fileInfo, err = os.Stat(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size (bytes):", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	log.Println("\n\nLast Modified:", fileInfo.ModTime(), "\n")
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n", fileInfo.Sys())
}
