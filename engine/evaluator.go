package engine

import (
	"context"

	"github.com/acemnto/cloud-cspm-insights/parser"

	"github.com/open-policy-agent/opa/v1/rego"
)

type EvaluationResult struct {
	RuleID  string
	Service string
	Passed  bool
}

func Evaluate(findings []parser.Finding, policyPath string) ([]EvaluationResult, error) {
	var results []EvaluationResult
	ctx := context.Background()

	for _, f := range findings {
		query := rego.New(
			rego.Query("data.cspm.allow"),
			rego.Load([]string{policyPath}, nil),
			rego.Input(f),
		)

		rs, err := query.Eval(ctx)
		if err != nil || len(rs) == 0 {
			results = append(results, EvaluationResult{RuleID: f.RuleID, Service: f.Service, Passed: false})
			continue
		}
		pass, ok := rs[0].Expressions[0].Value.(bool)
		results = append(results, EvaluationResult{RuleID: f.RuleID, Service: f.Service, Passed: pass && ok})
	}
	return results, nil
}
