# ECDD Case Management Folder - Mock Data Setup

This document describes the data schema and mock data for the ECDDCaseManagementFolder table. Use this as a reference when creating mock data for service development or testing.

## Spanner Schema

| Column Name | Column Type | Index | Description | Example |
|---|---|---|---|---|
| ecdd_case_management_folder_pk | STRING | Primary Index | Spanner generated random UUID | a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d |
| folder_name | STRING(255) | | Name of the case management folder | High Risk Customers |
| logged_at | TIMESTAMP | | Timestamp when the row was last logged/modified | 2026-01-15T10:00:00Z |
| updated_by | STRING(255) | | Username of the person who last updated/added/modified this folder record | admin@company.com |

## Mock Data Location

- Mock folder data: `data/case_folders.json`

## Mock Folders

The mock data set contains 10 folders:

| ecdd_case_management_folder_pk | folder_name | logged_at | updated_by |
|---|---|---|---|
| a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d | High Risk Customers | 2026-01-15T10:00:00Z | admin@company.com |
| b2c3d4e5-f6a7-5b6c-9d0e-1f2a3b4c5d6e | Pending Review | 2026-01-15T10:15:00Z | admin@company.com |
| c3d4e5f6-a7b8-6c7d-0e1f-2a3b4c5d6e7f | Suspended Accounts | 2026-01-15T10:30:00Z | admin@company.com |
| d4e5f6a7-b8c9-7d8e-1f2a-3b4c5d6e7f8a | VIP Customers | 2026-01-16T09:00:00Z | manager@company.com |
| e5f6a7b8-c9d0-8e9f-2a3b-4c5d6e7f8a9b | New Accounts Under Review | 2026-01-17T14:20:00Z | compliance@company.com |
| f6a7b8c9-d0e1-9f0a-3b4c-5d6e7f8a9b0c | Medium Risk Profile | 2026-01-18T11:45:00Z | risk.officer@company.com |
| a7b8c9d0-e1f2-0a1b-4c5d-6e7f8a9b0c1d | Escalated Cases | 2026-01-19T08:30:00Z | senior.manager@company.com |
| b8c9d0e1-f2a3-1b2c-5d6e-7f8a9b0c1d2e | Document Verification Queue | 2026-01-20T13:00:00Z | compliance@company.com |
| c9d0e1f2-a3b4-2c3d-6e7f-8a9b0c1d2e3f | Responsible Gambling Concerns | 2026-01-21T10:10:00Z | rg.team@company.com |
| d0e1f2a3-b4c5-3d4e-7f8a-9b0c1d2e3f4a | Threshold Breach Monitoring | 2026-01-22T15:25:00Z | monitoring@company.com |

## Sample Mock Record

```json
{
    "ecdd_case_management_folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
    "folder_name": "High Risk Customers",
    "logged_at": "2026-01-15T10:00:00Z",
    "updated_by": "admin@company.com"
}
```
