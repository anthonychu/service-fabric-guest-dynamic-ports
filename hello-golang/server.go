package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go "+runtime.Version()+"!</h1><img src=\"https://raw.githubusercontent.com/ashleymcnamara/gophers/master/Azure_Gophers.png\" width=810 height=600 />")
}
func main() {
	portPtr := flag.Int("port", 8000, "port to listen on")
	flag.Parse()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+strconv.Itoa(*portPtr), nil)
}
