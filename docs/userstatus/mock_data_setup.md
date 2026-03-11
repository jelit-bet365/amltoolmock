# ECDD User Status - Mock Data Setup

This document describes the data schema and mock data generation patterns for the ECDDUserStatus table. Use this as a reference when creating mock data for service development or testing.

## Spanner Schema

| Column Name | Column Type | Index | Description | Example |
|---|---|---|---|---|
| ECDDUserStatus_PK | STRING | Primary Index | Spanner generated random UUID | f9e8d7c6-b5a4-3210-fedc-ba9876543210 |
| user_id | INT64 | Unique Index on (user_id, user_name, country_id, state_id) | User identifier from the core user system | 12345 |
| user_name | STRING(255) | Unique Index on (user_id, user_name, country_id, state_id) | Display name of the user | John Doe |
| country_id | INT64 | Unique Index on (user_id, user_name, country_id, state_id) | Country identifier reference | 1 |
| state_id | INT64 | Unique Index on (user_id, user_name, country_id, state_id) | State identifier reference NULLABLE | 5 |
| ecdd_status | INT64 | | ECDD status (1=Not Required, 2=In progress, 3=Complete, 4=Suspended - Manual, 5=Suspended - Auto, 6=Closed, 7=Block Process) | 2 |
| ecdd_threshold | NUMERIC | | Threshold value | 10000.00 |
| ecdd_review_trigger | INT64 | | Review trigger type identifier | 9 |
| ecdd_suspension_due_date | DATE | | Suspension due date for the user | 2026-02-15 |
| ecdd_multiplier | NUMERIC | | Multiplier value applied to thresholds | 1.00 |
| ecdd_multiplier_rg_flag | BOOL | | RG (Responsible Gambling) flag indicator | FALSE |
| user_lt_engg_threshold_gbp | NUMERIC | | Lifetime engagement threshold in GBP | 5000.00 |
| user_lt_deposit_threshold_gbp | NUMERIC | | Lifetime deposit threshold in GBP | 10000.00 |
| user_12month_drop_threshold_gbp | NUMERIC | | 12-month drop threshold in GBP | 2500.00 |
| info_source | INT64 | | Information source identifier | 4 |
| sign_off_status | INT64 | | Sign-off status indicator | 1 |
| date_last_ecdd_sign_off | DATE | | Last ECDD sign-off date | 2026-01-20 |
| ecdd_rg_review_status | INT64 | | RG review status indicator | 2 |
| date_last_ecdd_rg_sign_off | DATE | | Last RG sign-off date | 2026-01-18 |
| ecdd_report_status | INT64 | | Report status indicator | 3 |
| ecdd_review_status | INT64 | | Review status indicator | 4 |
| ecdd_document_status | INT64 | | Document status indicator | 3 |
| ecdd_escalation_status | INT64 | | Escalation status indicator | 1 |
| uar_status | INT64 | | UAR (Unusual Activity Report) status indicator | 2 |
| logged_at | TIMESTAMP | | Timestamp when the row was last logged/modified | 2026-01-28T14:25:00Z |
| updated_by | STRING(255) | | Username of the person who last updated/added/modified this user status record | john.doe |

**Note:** The mock codebase includes a `language` field (STRING, ISO 639-1 code e.g. "EN") which is not present in the spanner schema above. This field is used in the mock API responses and should be accounted for if the service requires it.

## Mock Data Generation Pattern

The mock data is generated via `data/scripts/generate_users.js` with the following rules:

### Regions and Country IDs

| Region | Country IDs |
|---|---|
| Malta | 134, 135, 136 |
| USA | 231, 232, 233 |
| Australia | 13, 14, 15 |
| Gibraltar | 83, 84, 85 |

### Volume

- 30 users per region (120 total)
- Users distributed evenly across country IDs within each region

### Primary Key Format (Mock)

Pattern: `u{index}-{region}-{country_id}-{sequence}`

Example: `u001-malta-134-001`

**Note:** In production (Spanner), this is a randomly generated UUID (e.g. `f9e8d7c6-b5a4-3210-fedc-ba9876543210`).

### User IDs

- Sequential starting from `100001`

### State ID

- Only populated for USA region users (cycles through state IDs: 5, 12, 36, 48, 6, 33, etc.)
- `null` for all other regions

### Field Value Distributions

| Field | Pattern | Range |
|---|---|---|
| ecdd_status | cycled | 1-5 |
| ecdd_threshold | base 10000 + increments + random | 10000 - 70000 |
| ecdd_review_trigger | offset from index | 3-17 |
| ecdd_suspension_due_date | set only when `ecdd_status` cycles to 5 | date in Feb 2026 or null |
| ecdd_multiplier | 0.50 every 4th user, 1.00 otherwise | 0.50, 1.00 |
| ecdd_multiplier_rg_flag | true every 4th user | true/false |
| user_lt_engg_threshold_gbp | base 2000 + increments + random | 2000 - 17000 |
| user_lt_deposit_threshold_gbp | base 4000 + increments + random | 4000 - 34000 |
| user_12month_drop_threshold_gbp | base 1000 + increments + random | 1000 - 8500 |
| info_source | cycled | 1-8 |
| sign_off_status | cycled | 1-3 |
| date_last_ecdd_sign_off | set every 3rd user | date in Jan-Feb 2026 or null |
| ecdd_rg_review_status | cycled | 1-4 |
| date_last_ecdd_rg_sign_off | set every 2nd user | date in Jan-Feb 2026 or null |
| ecdd_report_status | cycled | 1-4 |
| ecdd_review_status | cycled | 1-4 |
| ecdd_document_status | cycled | 1-4 |
| ecdd_escalation_status | cycled | 1-3 |
| uar_status | cycled | 1-3 |
| language | cycled per region | EN, DE, FR, etc. |
| logged_at | incremented per user | Feb 2026 timestamps |
| updated_by | region-based | `{region}.compliance@company.com` |

### Numeric Precision

All float/numeric values are rounded to 2 decimal places.

## Sample Mock Record

```json
{
    "ecdd_user_status_pk": "u001-malta-134-001",
    "user_id": 100001,
    "user_name": "Matthew Borg",
    "country_id": 134,
    "language": "EN",
    "state_id": null,
    "ecdd_status": 1,
    "ecdd_threshold": 14031.98,
    "ecdd_review_trigger": 3,
    "ecdd_suspension_due_date": null,
    "ecdd_multiplier": 0.5,
    "ecdd_multiplier_rg_flag": true,
    "user_lt_engg_threshold_gbp": 2047.05,
    "user_lt_deposit_threshold_gbp": 4599.86,
    "user_12month_drop_threshold_gbp": 1104.47,
    "info_source": 1,
    "sign_off_status": 1,
    "date_last_ecdd_sign_off": "2026-01-10T00:00:00Z",
    "ecdd_rg_review_status": 1,
    "date_last_ecdd_rg_sign_off": "2026-01-08T00:00:00Z",
    "ecdd_report_status": 1,
    "ecdd_review_status": 1,
    "ecdd_document_status": 1,
    "ecdd_escalation_status": 1,
    "uar_status": 1,
    "logged_at": "2026-02-01T09:00:00Z",
    "updated_by": "malta.compliance@company.com"
}
```

## Data File Location

- Mock data: `data/user_statuses.json`
- Generator script: `data/scripts/generate_users.js`
