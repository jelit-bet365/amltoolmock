# ECDD Multiplier Config - Mock Data Setup

This document describes the data schema and mock data generation patterns for the ECDDMultiplierConfig table. Use this as a reference when creating mock data for service development or testing.

## Spanner Schema

| Column Name | Column Type | Index | Description | Example |
|---|---|---|---|---|
| ECDDMultiplierConfig_PK | STRING | Primary Index | Spanner generated random UUID | m1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c |
| country_id | INT64 | | Country identifier reference | 1 |
| state_id | INT64 | | State identifier reference NULLABLE | 5 |
| age_multipliers | ARRAY<INT64> | | Array of ages where 0.5 multiplier applies | [18, 19, 20, 21] |
| status_multiplier | BOOL | | Whether a status multiplier is applied | TRUE |
| is_active | BOOL | | Whether this config is active | TRUE |
| logged_at | TIMESTAMP | | Timestamp when the row was last logged/modified | 2026-01-28T14:25:00Z |
| updated_by | STRING(255) | | Username of the person who last updated/added/modified this record | admin@company.com |

## Mock Data Generation Pattern

The mock data is stored in `data/multiplier_configs.json`.

### Primary Key Format (Mock)

Pattern: Spanner-style UUID string.

Example: `m1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c`

**Note:** In production (Spanner), this is a randomly generated UUID.

### State ID

- Only populated for country IDs that correspond to states-based regions (e.g. USA)
- `null` for all other countries

### Field Value Distributions

| Field | Pattern | Range |
|---|---|---|
| country_id | varied | country IDs from supported regions |
| state_id | set only for applicable countries | state ID integer or null |
| age_multipliers | array of contiguous ages | e.g. [18, 19, 20, 21, 22, 23, 24, 25] |
| status_multiplier | varied | true/false |
| is_active | mostly true | true/false |
| logged_at | incremented per record | Jan-Feb 2026 timestamps |
| updated_by | static | admin@company.com |

## Sample Mock Record

```json
{
    "ecdd_multiplier_config_pk": "m1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c",
    "country_id": 1,
    "state_id": null,
    "age_multipliers": [18, 19, 20, 21, 22, 23, 24, 25],
    "status_multiplier": true,
    "is_active": true,
    "logged_at": "2026-01-28T14:25:00Z",
    "updated_by": "admin@company.com"
}
```

## Data File Location

- Mock data: `data/multiplier_configs.json`
