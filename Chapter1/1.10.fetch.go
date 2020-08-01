package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		resbytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}

		fmt.Printf("%s\n", resbytes)

	}
}
