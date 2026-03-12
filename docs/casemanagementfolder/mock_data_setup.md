# ECDD Case Management Folder - Mock Data Setup

This document describes the data schema and mock data for the ECDDCaseManagementFolder table. Use this as a reference when creating mock data for service development or testing.

## Spanner Schema

| Column Name | Column Type | Index | Description | Example |
|---|---|---|---|---|
| ecdd_case_management_folder_pk | STRING | Primary Index | Spanner generated random UUID | a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d |
| folder_name | STRING(255) | | Name of the case management folder | MT - Unusual Activity Reports |
| region | STRING(50) | | Region name for the folder (MALTA, GIBRALTAR, USA, AUSTRALIA) | MALTA |
| country_id | INT64 | | Country identifier (nullable) | null |
| state_id | INT64 | | State identifier (nullable, populated for USA configs only) | null |
| logged_at | TIMESTAMP | | Timestamp when the row was last logged/modified | 2026-01-15T10:00:00Z |
| updated_by | STRING(255) | | Username of the person who last updated/added/modified this folder record | admin@company.com |

## Mock Data Location

- Mock folder data: `data/case_folders.json`

## Mock Folders

The mock data set contains 120 folders total (68 Malta + 36 USA + 12 Australia + 4 Gibraltar). The table below shows a sample of 5 folders per region.

| ecdd_case_management_folder_pk | folder_name | region | country_id | state_id | logged_at | updated_by |
|---|---|---|---|---|---|---|
| a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d | MT - Unusual Activity Reports | MALTA | null | null | 2026-01-15T10:00:00Z | admin@company.com |
| b2c3d4e5-f6a7-5b6c-9d0e-1f2a3b4c5d6e | MT - NEW Population - ECDD | MALTA | null | null | 2026-01-15T10:15:00Z | admin@company.com |
| c3d4e5f6-a7b8-6c7d-0e1f-2a3b4c5d6e7f | MT - Report Required - ECDD | MALTA | null | null | 2026-01-15T10:30:00Z | admin@company.com |
| d4e5f6a7-b8c9-7d8e-1f2a-3b4c5d6e7f8a | MT - ROW - ISO Required (TL) | MALTA | null | null | 2026-01-16T09:00:00Z | admin@company.com |
| e5f6a7b8-c9d0-8e9f-2a3b-4c5d6e7f8a9b | MT - CX - Initial Contact | MALTA | null | null | 2026-01-17T14:20:00Z | admin@company.com |
| d6e7f8a9-b0c1-9d0e-3f4a-5b6c7d8e9f0a | US - Unusual Activity Reports | USA | null | null | 2026-01-26T10:00:00Z | usa.compliance@company.com |
| e7f8a9b0-c1d2-0e1f-4a5b-6c7d8e9f0a1b | US - NEW Population - ECDD | USA | null | null | 2026-01-26T12:00:00Z | usa.compliance@company.com |
| f8a9b0c1-d2e3-1f2a-5b6c-7d8e9f0a1b2c | US - Report Required – ECDD | USA | null | null | 2026-01-27T08:00:00Z | usa.compliance@company.com |
| a9b0c1d2-e3f4-2a3b-6c7d-8e9f0a1b2c3d | US - ISO Required (Specialist) | USA | null | null | 2026-01-27T14:30:00Z | usa.compliance@company.com |
| b0c1d2e3-f4a5-3b4c-7d8e-9f0a1b2c3d4e | US - Pending Closure | USA | null | null | 2026-01-28T09:00:00Z | usa.compliance@company.com |
| c1d2e3f4-a5b6-4c5d-8e9f-0a1b2c3d4e5f | AU - Unusual Activity Reports | AUSTRALIA | null | null | 2026-01-29T10:00:00Z | aus.compliance@company.com |
| d2e3f4a5-b6c7-5d6e-9f0a-1b2c3d4e5f6a | AU - NEW Population - ECDD | AUSTRALIA | null | null | 2026-01-29T12:00:00Z | aus.compliance@company.com |
| e3f4a5b6-c7d8-6e7f-0a1b-2c3d4e5f6a7b | AU - Report Required - ECDD | AUSTRALIA | null | null | 2026-01-30T08:30:00Z | aus.compliance@company.com |
| f4a5b6c7-d8e9-7f8a-1b2c-3d4e5f6a7b8c | AU - ISO Required | AUSTRALIA | null | null | 2026-01-30T11:00:00Z | aus.compliance@company.com |
| a5b6c7d8-e9f0-8a9b-2c3d-4e5f6a7b8c9d | AU - Monitoring | AUSTRALIA | null | null | 2026-01-31T09:00:00Z | aus.compliance@company.com |
| e1f2a3b4-c5d6-4e5f-8a9b-0c1d2e3f4a5b | GIB - Unusual Activity Reports | GIBRALTAR | null | null | 2026-01-23T09:00:00Z | gib.compliance@company.com |
| f2a3b4c5-d6e7-5f6a-9b0c-1d2e3f4a5b6c | ECDD Auto Suspend | GIBRALTAR | null | null | 2026-01-23T11:00:00Z | gib.compliance@company.com |
| a3b4c5d6-e7f8-6a7b-0c1d-2e3f4a5b6c7d | Details Changed | GIBRALTAR | null | null | 2026-01-24T08:30:00Z | gib.compliance@company.com |
| b4c5d6e7-f8a9-7b8c-1d2e-3f4a5b6c7d8e | NEW Population - AML Reports | GIBRALTAR | null | null | 2026-01-24T10:00:00Z | gib.compliance@company.com |

## Sample Mock Record

```json
{
    "ecdd_case_management_folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
    "folder_name": "MT - Unusual Activity Reports",
    "region": "MALTA",
    "country_id": null,
    "state_id": null,
    "logged_at": "2026-01-15T10:00:00Z",
    "updated_by": "admin@company.com"
}
```
