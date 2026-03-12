# ECDD User Case Management - Mock Data Setup

This document describes the data schema and mock data generation patterns for the ECDDUserCaseManagementFolder table. Use this as a reference when creating mock data for service development or testing.

## Spanner Schema

| Column Name | Column Type | Index | Description | Example |
|---|---|---|---|---|
| ecdd_user_case_management_folder_pk | STRING | Primary Index | Spanner generated random UUID | 604f0583-5016-48fb-8ed2-113121c204c4 |
| case_management_folder_pk | STRING | | Foreign key referencing ECDDCaseManagementFolder | a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d |
| user_status_pk | STRING | | Foreign key referencing ECDDUserStatus | u001-malta-134-001 |
| logged_at | TIMESTAMP | | Timestamp when the row was last logged/modified | 2026-02-01T09:00:00Z |
| updated_by | STRING(255) | | Username of the person who last updated/added/modified this record | malta.compliance@company.com |

## Relationship

The ECDDUserCaseManagementFolder table represents a many-to-many relationship between users (ECDDUserStatus) and folders (ECDDCaseManagementFolder). A single user status may be assigned to multiple folders, and a single folder may contain multiple user statuses.

## Mock Data Generation Pattern

The mock data is stored in `data/user_case_folders.json` and follows these rules:

### Volume

- Each user is assigned to 1-3 folders based on their index
- Users at index 0 (mod 3) receive 1 folder assignment
- Users at index 1 (mod 3) receive 2 folder assignments
- Users at index 2 (mod 3) receive 3 folder assignments

### Primary Key Format (Mock)

Each record uses a randomly generated UUID (e.g. `604f0583-5016-48fb-8ed2-113121c204c4`).

**Note:** In production (Spanner), this is also a randomly generated UUID.

### Field Value Distributions

| Field | Pattern |
|---|---|
| ecdd_user_case_management_folder_pk | random UUID per record |
| case_management_folder_pk | cycled from available ECDDCaseManagementFolder PKs in `data/case_folders.json` |
| user_status_pk | references the `ecdd_user_status_pk` from `data/user_statuses.json` |
| logged_at | matches the `logged_at` of the corresponding user status record |
| updated_by | matches the `updated_by` of the corresponding user status record |

### Numeric Precision

Not applicable — no numeric fields in this table.

## Sample Mock Record

```json
{
    "ecdd_user_case_management_folder_pk": "604f0583-5016-48fb-8ed2-113121c204c4",
    "case_management_folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
    "user_status_pk": "u001-malta-134-001",
    "logged_at": "2026-02-01T09:00:00Z",
    "updated_by": "malta.compliance@company.com"
}
```

## Data File Location

- Mock data: `data/user_case_folders.json`
