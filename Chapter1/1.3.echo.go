package main

import (
	"fmt" // to pretty print the results
	"os"  // to access the command line arguments
)

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

}
