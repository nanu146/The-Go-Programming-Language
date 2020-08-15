// echos the command line arguments
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "truncate new line character")
var sep = flag.String("sep", " ", "Seperator")

func main() {
	flag.Parse()
	fmt.Printf(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
