package main

import (
	"fmt"
	// "io"
	"net/http"
	"os"
	"log"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
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
	// // NOTE: if you do io.ReadAll right now, we can't read it later for the goquery. if we wanted to do this (we don't. we're using goquery), we could read the body into a byte slice, then uses bytes.newReader a couple times
	// fmt.Println("--- Body ---")
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("Error reading body: %v\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(body))


	///
	/// HTML parsing
	///

	fmt.Println("---------------")

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	summarizeTagsRecursively(doc.Selection, "")
}

func summarizeTagsRecursively(sel *goquery.Selection, indentation string) {
	for _, node := range sel.Nodes {

		// Only print element nodes (tags). ignore text, comments etc, for now
        if node.Type == html.ElementNode {
            fmt.Printf("%s%s %s\n", indentation, node.Data, node.Attr) // node.Data holds the tag name
        }
	}

	sel.Children().Each(func(i int, childSelection *goquery.Selection) {
		summarizeTagsRecursively(childSelection, indentation + "  ")
	})
}
