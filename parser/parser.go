package parser

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
)

type Finding struct {
	Service string                 `json:"service"`  // config | guardduty | cloudtrail
	RuleID  string                 `json:"rule_id"`
	Details map[string]interface{} `json:"details"`
}

func LoadFindings(dir string) ([]Finding, error) {
	var results []Finding

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".json" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			var entries []Finding
			if err := json.Unmarshal(content, &entries); err != nil {
				return err
			}
			results = append(results, entries...)
		}
		return nil
	})
	return results, err
}