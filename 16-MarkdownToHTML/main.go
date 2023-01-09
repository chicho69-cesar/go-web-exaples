package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gomarkdown/markdown"
)

func main() {
	file := "test.md"
	
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("%s file not found", file)
	}

	html := markdown.ToHTML(content, nil, nil)
	fmt.Println(string(html))

	fileOut := "test.html"
	err = os.WriteFile(fileOut, html, 0644)
	if err != nil {
		log.Fatalf("Could not write to %s", fileOut)
	}

	fmt.Printf("HTML outputted to %s", fileOut)
}
