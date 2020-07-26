package main

import (
	"fmt"     // to pretty print output
	"os"      // to get user arguments
	"strings" // for string manipulations
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
