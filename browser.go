package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Get URL from cli args
	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL. Usage: browser https://example.com")
		os.Exit(1)
	}

	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		os.Exit(1)
	}

	// 'close body' when function exits
	defer resp.Body.Close()

	// print headers in a list
	fmt.Println("--- Headers ---")
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("| %s: %s\n", key, value);
		}
	}
	fmt.Println("---------------")

	// body content
	fmt.Println("--- Body ---")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
}