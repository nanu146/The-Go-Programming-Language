package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// any argument provided to the program should be a valid path to a file
	files := os.Args[1:]
	counts := make(map[string]int)

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		// this gets executed when a file reference is passed
		for _, fl := range files {
			if f, err := os.Open(fl); err != nil {
				// print error message
				fmt.Fprintf(os.Stderr, "%v\n", err)
				// after printing the error message continue working on other files
				continue
			} else {
				// calling countLines on each file
				countLines(f, counts)
			}

		}
	}

	for key, value := range counts {
		if value >= 1 {
			fmt.Printf("the sentence: \"%v\" appears \"%v\" times\n", key, value)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}
