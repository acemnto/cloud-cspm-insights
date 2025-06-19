# github.com/acemnto/cloud-cspm-insights

# ☁️ cloud-cspm-insights

> Evaluate AWS Config, GuardDuty, and CloudTrail findings using OPA policies to generate posture and compliance insights.

---

## 🔍 Overview

`cloud-cspm-insights` is a modular Golang tool designed for CSPM (Cloud Security Posture Management). It parses security data from:

- ✅ **AWS Config** (e.g. root MFA checks)
- 🔐 **GuardDuty** (e.g. threat severity analysis)
- 🧾 **CloudTrail** (e.g. logging enabled)

Each data type is evaluated against [OPA](https://www.openpolicyagent.org/) policies using custom Rego logic, with a final compliance posture report.

---

## 🧱 Project Structure

```
cloud-cspm-insights/
├── main.go                     # CLI entry point
├── parser/                    # Parses data from AWS config files
│   └── parser.go
├── engine/                    # OPA policy evaluation engine
│   └── evaluator.go
├── reports/                   # Report generator
│   └── report.go
├── policies/
│   └── cspm.rego              # Rego policies per service
├── data/                      # Example JSON input files
│   ├── config.json
│   ├── guardduty.json
│   └── cloudtrail.json
└── go.mod
```

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/acemnto/cloud-cspm-insights.git
cd cloud-cspm-insights
go mod tidy
```

### 2. Add Data
Add JSON data files from AWS Config, GuardDuty, and CloudTrail to the `data/` folder.

```json
[
  {
    "service": "config",
    "rule_id": "1.1",
    "details": {
      "root_mfa_enabled": false
    }
  }
]
```

### 3. Define Policies
Edit `policies/cspm.rego` with logic such as:

```rego
package cspm

allow {
  input.service == "config"
  input.rule_id == "1.1"
  input.details.root_mfa_enabled == true
}

allow {
  input.service == "guardduty"
  input.details.severity < 4.0
}

allow {
  input.service == "cloudtrail"
  input.details.logging_enabled == true
}
```

### 4. Run the Tool

```bash
go run main.go
```

---

## 📊 Example Output

```
--- Cloud CSPM Insights Report ---
Service: config       | Rule: 1.1                | Status: FAIL ❌
Service: guardduty    | Rule: gd1                | Status: FAIL ❌
Service: cloudtrail   | Rule: ct1                | Status: PASS ✅
```

