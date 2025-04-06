package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gobelinor/info2parser"
)

func main() {
	filePath := flag.String("file", "", "Path to the $I file to parse")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide the path to the $I file using the -file flag.")
		os.Exit(1)
	}

	data, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	info2, err := info2parser.Parse(data)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("==== $I FILE ====")
	fmt.Printf("Header       : %d\n", info2.Header)
	fmt.Printf("FileSize       : %d octets\n", info2.FileSize)
	fmt.Printf("DeletionTime   : %s\n", info2parser.FiletimeToTime(info2.DeletionTime).Format(time.RFC3339))
	fmt.Printf("FileNameLength : %d\n", info2.FileNameLength)
	fmt.Printf("OriginalPath : %s\n", info2.OriginalPath)
}
