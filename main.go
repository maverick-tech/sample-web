package main

import (
	"fmt"
	"net/http"
)

func webHandlerFunc(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>Hello, World</h1>")
}

func main() {
	http.HandleFunc("/", webHandlerFunc)
	http.ListenAndServe(":3030", nil)
}
