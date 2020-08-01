package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		if resp, err := http.Get(url); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		} else {
			fmt.Println("printing with buffio")
			//printWithBuffio(resp.Body)
			fmt.Println("printing with io.copy")
			printwithCopy(resp.Body, os.Stdout)
		}

	}
}

func printWithBuffio(reader io.Reader) {
	input := bufio.NewScanner(reader)

	for input.Scan() {
		fmt.Println(input.Text())
	}
}

func printwithCopy(reader io.Reader, out io.Writer) {
	_, err := io.Copy(out, reader)
	if err != nil {
		fmt.Fprint(os.Stderr, "%v\n", err)
	}

}
