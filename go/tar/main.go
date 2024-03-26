package main

import (
	"archive/tar"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Step 1 - List the files in the archive: read the file, parse the headers to get the file names.

	lisArchiveContents := flag.Bool("t", false, "-t List archive contents to stdout.")
	// prompt user for filepath
	flag.Parse()
	flagArgs := flag.Args()

	if len(flagArgs) == 0 {
		log.Fatalln("No file name provided")
	}
	filePath := flagArgs[0]

	// open file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// create new tar reader
	tr := tar.NewReader(file)
	for {
		header, err := tr.Next()

		// exit for loop when end of file
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		// list archive contents if -t flag provided
		if *lisArchiveContents {
			fmt.Println(header.Name)
		}
	}

}
