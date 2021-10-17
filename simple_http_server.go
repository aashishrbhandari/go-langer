package main

import (
	"net/http"
	"fmt"
)

func DefaultHandler(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", DefaultHandler)
	
	fmt.Println("Starting Web Server @ 80")
	http.ListenAndServe(":80", nil)
	
}