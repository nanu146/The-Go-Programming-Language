package main

import (
	"bufio" // for buffered input output scanner
	"fmt"   // for pretty printing the output on console
	"os"    // for getting inputs from user on command line
)

func main() {
	counts := make(map[string]int)      // map of key type string and value type int for holding sentence counts
	input := bufio.NewScanner(os.Stdin) // scanner for standard input

	for input.Scan() {
		// when you add an item for the first time the zero value of int will be 0
		// and ++ operator increments it by 1 making the first entry value to be 1
		counts[input.Text()]++
	}

	for key, value := range counts {
		if value > 1 {
			fmt.Println(key, " -> Appears: ", value, " times.")
		}
	}
}
