package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Options struct {
	printBytes bool
	printChars bool
	printLines bool
	printWords bool
}

type stats struct {
	bytes uint64
	lines uint64
	words uint64
	chars uint64
	filename string
}


// CalculateStats function reads the file and calculates the number of bytes, lines, words, and characters in the file.
// and returns a stats struct with the calculated values.
func CalculateStats(reader *bufio.Reader) stats {
	var prevChar rune
	var bytesCount uint64
	var linesCount uint64
	var wordsCount uint64
	var charsCount uint64


	// The infinite loop is used to read the file until the end of the file is reached.
	// The reader.ReadRune() method reads a single UTF-8 encoded Unicode code point from the file.
	// The prevChar variable is used to keep track of the previous character read from the file. This is used to determine if the current character is a word.
	// In the case where you would have a double space, the prevChar variable would be a space and the current character would be a space, so it would not be counted as a word.


	for {
		cr, br, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// If end of file is reached, check if the last character is not a space
				if prevChar != rune(0) && !unicode.IsSpace(prevChar) {
					// If it's not a space, then it's a word
					wordsCount++
				}
				// We've reached the end of the file, so we can break the loop
				break
			}
			log.Fatal(err)
		}

		// br is the number of bytes read
		bytesCount += uint64(br)
		// Increase the character count
		charsCount++

		// If the current character is a newline, increase the line count
		if cr == '\n' {
			linesCount++
		}

		// If the previous character is not a space and the current character is a space, increase the word count
		if !unicode.IsSpace(prevChar) && unicode.IsSpace(cr) {
			wordsCount++
		}

		// Set the previous character to the current character
		prevChar = cr
	}

	return stats{bytes: bytesCount, lines: linesCount, words: wordsCount, chars: charsCount}
}


// CalculateStatsWithTotals function reads the file and calculates the number of bytes, lines, words, and characters in the file.
func CalculateStatsWithTotals(reader *bufio.Reader, filename string, options Options, totals *stats) {
	fileStats := CalculateStats(reader)
	fileStats.filename = filename

	fmt.Println(formatStats(options, fileStats, filename))

	totals.bytes += fileStats.bytes
	totals.lines += fileStats.lines
	totals.words += fileStats.words
}



func formatStats(commandLineOptions Options, fileStats stats, filename string) string {
	var cols []string

	if commandLineOptions.printBytes {
		cols = append(cols, strconv.FormatUint(fileStats.bytes, 10))
	}

	if commandLineOptions.printLines {
		cols = append(cols, strconv.FormatUint(fileStats.lines, 10))
	}

	if commandLineOptions.printWords {
		cols = append(cols, strconv.FormatUint(fileStats.words, 10))
	}

	if commandLineOptions.printChars {
		cols = append(cols, strconv.FormatUint(fileStats.chars, 10))
	}

	cols = append(cols, filename)

	return strings.Join(cols, "\t")
}
