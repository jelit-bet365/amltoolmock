# ECDD Business Profile - Mock Data Setup

This document describes the data schema and mock data generation patterns for the ECDDBusinessProfile table. Use this as a reference when creating mock data for service development or testing.

## Spanner Schema

| Column Name | Column Type | Index | Description | Example |
|---|---|---|---|---|
| ECDDBusinessProfile_PK | STRING | Primary Index | Spanner generated random UUID | b1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c |
| country_id | INT64 | | Country identifier reference | 1 |
| state_id | INT64 | | State identifier reference NULLABLE | null |
| risk_status_id | INT64 | | Risk status identifier reference | 2 |
| average_deposit | NUMERIC | | Average deposit value | 5000.00 |
| deposit_multiplier | NUMERIC | | Multiplier applied to the deposit threshold | 2.00 |
| time_period_days | INT64 | | Time period in days for the profile | 28 |
| enabled | BOOL | | Whether the business profile is active | TRUE |
| logged_at | TIMESTAMP | | Timestamp when the row was last logged/modified | 2026-01-28T14:25:00Z |
| updated_by | STRING(255) | | Username of the person who last updated/added/modified this business profile record | admin@company.com |

## Mock Data Generation Pattern

The mock data is stored in `data/business_profiles.json`.

### Field Value Distributions

| Field | Pattern | Range |
|---|---|---|
| country_id | cycled | 1-10 |
| state_id | set only for applicable country entries | int or null |
| risk_status_id | cycled | 1-5 |
| average_deposit | base 1000 + increments + random | 1000 - 20000 |
| deposit_multiplier | cycled | 1.00, 1.50, 2.00, 2.50 |
| time_period_days | cycled | 7, 14, 28, 30, 60, 90 |
| enabled | alternating | true/false |
| logged_at | incremented per record | Jan-Feb 2026 timestamps |
| updated_by | static | admin@company.com |

### Numeric Precision

All float/numeric values are rounded to 2 decimal places.

## Sample Mock Record

```json
{
    "ecdd_business_profile_pk": "b1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c",
    "country_id": 1,
    "state_id": null,
    "risk_status_id": 2,
    "average_deposit": 5000.00,
    "deposit_multiplier": 2.00,
    "time_period_days": 28,
    "enabled": true,
    "logged_at": "2026-01-28T14:25:00Z",
    "updated_by": "admin@company.com"
}
```

## Data File Location

- Mock data: `data/business_profiles.json`
