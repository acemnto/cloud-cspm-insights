package reports

import (
	"fmt"

	"github.com/acemnto/cloud-cspm-insights/engine"
)

func PrintReport(results []engine.EvaluationResult) {
	fmt.Println("\n--- Cloud CSPM Insights Report ---")
	for _, r := range results {
		status := "FAIL ❌"
		if r.Passed {
			status = "PASS ✅"
		}
		fmt.Printf("Service: %-12s | Rule: %-20s | Status: %s\n", r.Service, r.RuleID, status)
	}
}
