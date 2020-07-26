package main

import (
	"fmt" // to pretty print the output on terminal
	"os"  // to capture command line arguments
)

func main() {
	var s, sep = "", ""

	for _, val := range os.Args[1:] {
		s += sep + val
		fmt.Println("the place is: ", &s)
		sep = " "
	}

	fmt.Println(s)
}
