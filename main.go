package main

import (
	"log"

	"github.com/acemnto/cloud-cspm-insights/parser"
)

func main() {
	_, err := parser.LoadFindings("./data")
	if err != nil {
		log.Fatalf("Failed to load findings: %v", err)
	}

}
