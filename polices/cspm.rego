package cspm

default allow = false

# AWS Config: Ensure root account MFA is enabled
allow {
  input.service == "config"
  input.rule_id == "1.1"
  input.details.root_mfa_enabled == true
}

# GuardDuty: Medium and High severity only
allow {
  input.service == "guardduty"
  input.details.severity < 4.0
}

# CloudTrail: Logging must be enabled
allow {
  input.service == "cloudtrail"
  input.details.logging_enabled == true
}