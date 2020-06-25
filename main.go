package main

import (
	"fmt"
	"net/http"
	"os"
)

func webHandlerFunc(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>Hello, World</h1>")
}

func main() {
	port:=os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", webHandlerFunc)
	http.ListenAndServe(":" + port, nil)
}
