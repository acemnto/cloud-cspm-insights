package main

import (
	"log"

	"github.com/acemnto/cloud-cspm-insights/engine"
	"github.com/acemnto/cloud-cspm-insights/parser"
	"github.com/acemnto/cloud-cspm-insights/reports"
)

func main() {
	inputs, err := parser.LoadFindings("./data")
	if err != nil {
		log.Fatalf("Failed to load findings: %v", err)
	}

	results, err := engine.Evaluate(inputs, "./policies/cspm.rego")
	if err != nil {
		log.Fatalf("Policy evaluation error: %v", err)
	}

	reports.PrintReport(results)
}
