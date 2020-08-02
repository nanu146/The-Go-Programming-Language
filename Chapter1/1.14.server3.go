package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler3(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "%s %s %s\n", req.Method, req.URL, req.Proto)

	for k, v := range req.Header {
		fmt.Fprintf(resp, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(resp, "Host = %q\n", req.Host)
	fmt.Fprintf(resp, "RemoteAddr = %q\n", req.RemoteAddr)

	if err := req.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range req.Form {
		fmt.Fprintf(resp, "Form[%q] = %q\n", k, v)
	}
}
