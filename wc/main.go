package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {

	// configure flags
	countBytesPtr := flag.Bool("c", false, "-c outputs the number of bytes in a file.")
	countLinesPtr := flag.Bool("l", false, "-l outputs the number of lines in a file.")
	countWordsPtr := flag.Bool("w", false, "-w outputs the number of words in a file.")
	countCharsPtr := flag.Bool("m", false, "-m outputs the number of characters in a file.")
	flag.Parse()
	
	// get the non-flag arguments
	flagArgs := flag.Args()
	flagArgsLen := len(flagArgs)

	// Exit program if no file path argument
	if flagArgsLen == 0 {
		panic("No file path argument")
	}
	filePath := flagArgs[0]

	
	// open file
	f, err := os.Open(filePath)
	check(err)
	// close file at the end
	defer f.Close()


	// get amount fo bytes in file
	s, err := f.Stat()
	check(err)
	bytes := s.Size()

	if *countBytesPtr {
		fmt.Printf("bytes: %v\n",  bytes)
	}

	br := bufio.NewReader(f)
	lineCount := 0
	wordCount := 0
	charCount := 0

	for {
		b,err := br.ReadByte()

		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Println(err)
			break
		}

		if b != '\n' && b != ' ' && b != 0 {
			charCount++
		}

		if b == '\n' {
			lineCount++
			wordCount++
		} 
		if b == ' ' {
			wordCount++
		}


		if err != nil {
			break
		}
	}

	if *countLinesPtr {
		fmt.Printf("lines: %v\n",  lineCount)
	}

	if *countWordsPtr {
		fmt.Printf("words: %v\n",  wordCount)
	}

	if *countCharsPtr {
		fmt.Printf("characters: %v\n",  charCount)
	}


	if !*countBytesPtr && !*countCharsPtr && !*countLinesPtr && !*countWordsPtr {
		fmt.Printf("bytes: %v\n",  bytes)
		fmt.Printf("lines: %v\n",  lineCount)
		fmt.Printf("words: %v\n",  wordCount)
		fmt.Printf("characters: %v\n",  charCount)

	}

}
