//go:build !solution

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	urls := os.Args[1:]

	if len(urls) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s <url1> <url2> ...\n", os.Args[0])
		os.Exit(1)
	}

	for _, url := range urls {
		response, err := http.Get(url)
		errorHandler(err)
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		errorHandler(err)
		fmt.Printf("%s\n", body)
	}
}
