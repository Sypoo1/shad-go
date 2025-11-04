//go:build !solution

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func makeRequest(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	response, err := http.Get(url)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return 
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Error reading body from %s: %v\n", url, err)
		return
	}

	fmt.Printf("%.2fs    %d  %s\n", elapsed.Seconds(), len(body), url)
}

func main() {

	var wg sync.WaitGroup
	urls := os.Args[1:]
	
	start := time.Now()

	for _, url := range urls {
		wg.Add(1)
		go makeRequest(url, &wg)
	}
	wg.Wait()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
