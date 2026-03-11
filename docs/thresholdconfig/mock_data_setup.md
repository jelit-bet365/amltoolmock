# ECDD Threshold Config - Mock Data Setup

This document describes the data schema and mock data generation patterns for the ECDDThresholdConfig table. Use this as a reference when creating mock data for service development or testing.

## Spanner Schema

| Column Name | Column Type | Index | Description | Example |
|---|---|---|---|---|
| ecdd_threshold_config_pk | STRING | Primary Index | Spanner generated random UUID | t1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c |
| title | STRING(255) | Unique Index | Descriptive name for the threshold rule | UK 28 day NDL 50K |
| is_active | BOOL | | Activation status of the threshold rule | TRUE |
| country_id | INT64 | | Country identifier where threshold applies | 1 |
| state_id | INT64 | | State identifier where threshold applies NULLABLE (for country-wide) | 5 |
| type | INT64 | | Threshold type: 1=Deposit, 2=Net Deposit, 3=Stakes | 2 |
| reinvest | BOOL | | Whether to look at if there is any reinvestment on the account | TRUE |
| value | NUMERIC | | Monetary threshold amount that triggers the rule | 10000.00 |
| currency_id | INT64 | | Currency identifier for the threshold value | 1 |
| period | INT64 | | Time period: 1=24hrs, 2=28days, 3=84days, 4=91days, 5=182days, 6=365days | 2 |
| use_multipliers | BOOL | | Whether to apply ECDD multiplier adjustments | TRUE |
| use_rg_flag | BOOL | | Whether to consider ECDD Multiplier RG Flag | TRUE |
| apply_all_statuses | BOOL | | Apply to all ECDD statuses (TRUE) or only "Not Required" (FALSE) | FALSE |
| backfill | BOOL | | Apply to existing accounts (TRUE) or new accounts only (FALSE) | FALSE |
| hierarchy | INT64 | Unique Index | Priority order for threshold evaluation (lower number = higher priority) | 10 |
| ecdd_status | INT64 | | ECDD Status to set when threshold is breached | 2 |
| ecdd_review_status | INT64 | | ECDD Review Status to set when threshold is breached | 4 |
| ecdd_report_status | INT64 | | ECDD Report Status to set when threshold is breached | 2 |
| sign_off_status | INT64 | | Sign Off Status to set when threshold is breached | 1 |
| customer_risk_level | INT64 | | Customer Risk Level: 1=Low, 2=Medium, 3=Medium-High, 4=High | 2 |
| ndl_28_day_gbp | NUMERIC | | 28 Day Net Deposit Limit in GBP to be applied | 5000.00 |
| ndl_monthly_gbp | NUMERIC | | Monthly Net Deposit Limit in GBP to be applied | 20000.00 |
| case_management_folder_pk | STRING | | Foreign key on Case Management folder table for user assignment NULLABLE | 3 |
| logged_at | TIMESTAMP | | Timestamp when the row was last logged/modified | 2026-01-28T14:25:00Z |
| updated_by | STRING(255) | | Username of the person who last updated/added/modified this threshold configuration | john.doe |

## Mock Data Generation Pattern

The mock data is stored in `data/threshold_configs.json`.

### Types and Periods

| Type | Description |
|---|---|
| 1 | Deposit |
| 2 | Net Deposit |
| 3 | Stakes |

| Period | Description |
|---|---|
| 1 | 24hrs |
| 2 | 28days |
| 3 | 84days |
| 4 | 91days |
| 5 | 182days |
| 6 | 365days |

### Field Value Distributions

| Field | Pattern | Range |
|---|---|---|
| is_active | mostly true, some false | true/false |
| country_id | varies per config | 1, 83, 134, 231 |
| state_id | populated for USA configs only | int or null |
| type | cycled | 1-3 |
| reinvest | cycled | true/false |
| value | varies by type and period | 5000 - 100000 |
| currency_id | fixed | 1 |
| period | cycled | 1-6 |
| use_multipliers | cycled | true/false |
| use_rg_flag | cycled | true/false |
| apply_all_statuses | cycled | true/false |
| backfill | cycled | true/false |
| hierarchy | incremented | 10, 20, 30, ... |
| ecdd_status | cycled | 1-4 |
| ecdd_review_status | cycled | 1-4 |
| ecdd_report_status | cycled | 1-4 |
| sign_off_status | cycled | 1-3 |
| customer_risk_level | cycled | 1-4 |
| ndl_28_day_gbp | base 5000 + increments | 5000 - 25000 |
| ndl_monthly_gbp | base 20000 + increments | 20000 - 100000 |
| case_management_folder_pk | set on some records | UUID or null |
| logged_at | incremented per record | Jan-Feb 2026 timestamps |
| updated_by | fixed pattern | john.doe@company.com |

### Numeric Precision

All float/numeric values are rounded to 2 decimal places.

## Sample Mock Record

```json
{
    "ecdd_threshold_config_pk": "t1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c",
    "title": "UK 28 day NDL 50K",
    "is_active": true,
    "country_id": 1,
    "state_id": null,
    "type": 2,
    "reinvest": true,
    "value": 50000.00,
    "currency_id": 1,
    "period": 2,
    "use_multipliers": true,
    "use_rg_flag": true,
    "apply_all_statuses": false,
    "backfill": false,
    "hierarchy": 10,
    "ecdd_status": 2,
    "ecdd_review_status": 4,
    "ecdd_report_status": 2,
    "sign_off_status": 1,
    "customer_risk_level": 3,
    "ndl_28_day_gbp": 5000.00,
    "ndl_monthly_gbp": 20000.00,
    "case_management_folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
    "logged_at": "2026-01-28T14:25:00Z",
    "updated_by": "john.doe@company.com"
}
```

## Data File Location

- Mock data: `data/threshold_configs.json`
