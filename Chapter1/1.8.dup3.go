// this program tries to read the whole file and then perform
// duplicate identification
// this works with files alone and not command line standard input
// use the file under ./data/text.txt as input for this program
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]

	for _, file := range files {
		if f, err := ioutil.ReadFile(file); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		} else {
			for _, line := range strings.Split(string(f), "\n") {
				counts[line]++
			}
		}
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Println("the sentence: \"", line, "\" appears ", count, " times.")
		}
	}

}
