package main

import (
	"fmt"
	// "io"
	// "net/http"
	"os"
)

func main() {
	// Get URL from cli args
	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL. Usage: browser https://example.com")
		os.Exit(1)
	}
}