package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("total time elapsed %.2f", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- fmt.Sprintf("error: %v\n", err)
		return
	} else {
		nbytes, err := io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close() // dont leak resources

		if err != nil {
			ch <- fmt.Sprintf("error: %v\n", err)
		}

		sec := time.Since(start).Seconds()

		ch <- fmt.Sprintf("%.2f %7d %v", sec, nbytes, url)

	}

}
