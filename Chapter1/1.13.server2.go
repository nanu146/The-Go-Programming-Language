package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler2(resp http.ResponseWriter, req *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(resp, "path: %s \n", req.URL.Path)
}

func counter(resp http.ResponseWriter, req *http.Request) {
	mu.Lock()
	fmt.Fprintf(resp, "this web page has beeb visited %d times", count)
	mu.Unlock()
}
