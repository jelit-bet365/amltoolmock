
# ECDD Mock API - Developer Documentation

> **Base URL:** `http://localhost:3001`  
> **Version:** v1  
> **Content-Type:** `application/json`

## Table of Contents
- [Region & Country Mapping](#region--country-mapping)
- [Authentication](#authentication)
- [Error Handling](#error-handling)
- [API Endpoints](#api-endpoints)
  - [Health Check](#health-check)
  - [User Status](#user-status-endpoints)
  - [Threshold Configuration](#threshold-configuration-endpoints)
  - [Case Management Folders](#case-management-folder-endpoints)
  - [Multiplier Configuration](#multiplier-configuration-endpoints)
  - [Business Profiles](#business-profile-endpoints)
  - [User-Folder Assignments](#user-folder-assignment-endpoints)
- [React Integration Examples](#react-integration-examples)

---

## Region & Country Mapping

### Active Regions

**Important:** Each region maps to **multiple country IDs** for regulatory and jurisdictional purposes.

| Region | Country IDs | Total Users | State Support | Regulatory Body |
|--------|------------|-------------|---------------|-----------------|
| **Malta** | 134, 135, 136 | 30 | No | Malta Gaming Authority (MGA) |
| **USA** | 231, 232, 233 | 30 | Yes | State Gaming Commissions |
| **Australia** | 13, 14, 15 | 30 | No | ACMA |
| **Gibraltar** | 83, 84, 85 | 30 | No | Gibraltar Gambling Commission |

### React Constants File Example

```javascript
// src/constants/regions.js
export const REGIONS = {
  MALTA: {
    name: 'Malta',
    countryIds: [134, 135, 136],
    code: 'MT',
    currency: 'EUR'
  },
  USA: {
    name: 'United States',
    countryIds: [231, 232, 233],
    code: 'US',
    currency: 'USD'
  },
  AUSTRALIA: {
    name: 'Australia',
    countryIds: [13, 14, 15],
    code: 'AU',
    currency: 'AUD'
  },
  GIBRALTAR: {
    name: 'Gibraltar',
    countryIds: [83, 84, 85],
    code: 'GI',
    currency: 'GBP'
  }
};

// Helper function to get region by country ID
export const getRegionByCountryId = (countryId) => {
  for (const [key, region] of Object.entries(REGIONS)) {
    if (region.countryIds.includes(countryId)) {
      return { ...region, key };
    }
  }
  return null;
};
```

---

## Authentication

**Current Implementation:** No authentication required (mock API)

---

## Error Handling

### Error Response Format

```json
{
  "error": "Error message description"
}
```

### HTTP Status Codes

| Code | Meaning | When It Occurs |
|------|---------|----------------|
| 200 | OK | Successful GET/PUT request |
| 201 | Created | Successful POST request |
| 400 | Bad Request | Invalid request body or parameters |
| 404 | Not Found | Resource not found |
| 405 | Method Not Allowed | Invalid HTTP method for endpoint |
| 500 | Internal Server Error | Server-side error |

---

## API Endpoints

### Health Check

#### GET /health

Check API health status.

**Response:** `200 OK`
```json
{
  "status": "healthy",
  "service": "ECDD Mock API"
}
```

---

### User Status Endpoints

#### GET /api/v1/users

Get list of all users (120 total - 30 per region).

**Query Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `country_ids` | string | Comma-separated list of country IDs to filter by (e.g. `134,135`) |
| `ecdd_status` | integer | Filter by ECDD status value |
| `search` | string | Case-insensitive substring match on `user_name` |
| `page` | integer | Page number (1-based) for pagination |
| `page_size` | integer | Number of results per page |
| `sort_by` | string | Field to sort by (default: `user_id`) |
| `sort_order` | string | Sort direction: `asc` or `desc` |

**Response:** `200 OK`
```json
[
  {
    "ecdd_user_status_pk": "u001-malta-134-001",
    "user_id": 100001,
    "user_name": "Matthew Borg",
    "country_id": 134,
    "state_id": null,
    "ecdd_status": 1,
    "ecdd_threshold": 12345.67,
    "ecdd_review_trigger": 3,
    "ecdd_suspension_due_date": null,
    "ecdd_multiplier": 0.50,
    "ecdd_multiplier_rg_flag": true,
    "user_lt_net_deposit_threshold_gbp": 2500.00,
    "user_lt_deposit_threshold_gbp": 5000.00,
    "user_12month_net_deposit_threshold_gbp": 1250.00,
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
]
```

**React Example:**
```javascript
const fetchUsers = async () => {
  const response = await fetch('http://localhost:3001/api/v1/users');
  if (!response.ok) throw new Error('Failed to fetch users');
  return await response.json();
};
```

---

#### GET /api/v1/users/:id

Get a specific user by their UUID.

**Response:** `200 OK` - Returns single user object

**Error Response:** `404 Not Found`
```json
{
  "error": "User not found"
}
```

---

#### POST /api/v1/users

Create a new user.

**Response:** `201 Created` - Returns created user object

---

#### PUT /api/v1/users/:id

Update an existing user.

**Response:** `200 OK` - Returns updated user object

---

#### DELETE /api/v1/users/:id

Delete a user. Hard deletes the user and cascade-deletes all related folder assignments.

**Response:** `200 OK`
```json
{
  "message": "User deleted successfully. 2 folder assignment(s) cascade-deleted."
}
```

---

#### GET /api/v1/users/:id/folders

Get the list of folders that a user is assigned to.

**Response:** `200 OK`
```json
[
  {
    "ecdd_case_management_folder_pk": "f001",
    "folder_name": "Malta High Risk",
    "logged_at": "2026-01-15T10:00:00Z",
    "updated_by": "admin@company.com"
  }
]
```

**Error Response:** `404 Not Found` — if the user does not exist.

---

### Threshold Configuration Endpoints

#### GET /api/v1/thresholds

Get list of all threshold configurations.

**Response:** `200 OK`
```json
[
  {
    "ecdd_threshold_config_pk": "t1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c",
    "title": "UK 28 day NDL 50K",
    "is_active": true,
    "country_id": 1,
    "state_id": null,
    "type": 2,
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
]
```

---

#### GET /api/v1/thresholds/:id

Get a specific threshold configuration.

**Response:** `200 OK` - Returns single threshold object

---

#### POST /api/v1/thresholds

Create a new threshold configuration.

**Response:** `201 Created` - Returns created threshold object

---

#### PUT /api/v1/thresholds/:id

Update an existing threshold configuration.

**Response:** `200 OK` - Returns updated threshold object

---

#### DELETE /api/v1/thresholds/:id

Soft delete a threshold configuration (sets deleted_at timestamp).

**Response:** `200 OK`
```json
{
  "message": "Threshold soft-deleted successfully"
}
```

---

### Case Management Folder Endpoints

#### GET /api/v1/folders

Get list of all case management folders.

**Response:** `200 OK`
```json
[
  {
    "ecdd_case_management_folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
    "folder_name": "High Risk Customers",
    "logged_at": "2026-01-15T10:00:00Z",
    "updated_by": "admin@company.com"
  }
]
```

---

#### GET /api/v1/folders/:id

Get a specific folder by UUID.

**Response:** `200 OK` - Returns single folder object

---

#### POST /api/v1/folders

Create a new case management folder.

**Request Body:**
```json
{
  "folder_name": "Malta - Under Investigation",
  "updated_by": "malta.compliance@company.com"
}
```

**Response:** `201 Created` - Returns created folder object

---

#### PUT /api/v1/folders/:id

Update an existing folder.

**Response:** `200 OK` - Returns updated folder object

---

### Multiplier Configuration Endpoints

#### GET /api/v1/multipliers

Get list of all multiplier configurations.

**Response:** `200 OK`
```json
[
  {
    "ecdd_multiplier_config_pk": "m1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c",
    "country_id": 134,
    "state_id": null,
    "age_multipliers": [18, 19, 20, 21, 22, 23, 24, 25],
    "status_multiplier": true,
    "is_active": true,
    "created_at": "2026-01-05T10:00:00Z",
    "updated_at": "2026-01-05T10:00:00Z",
    "updated_by": "system@company.com"
  }
]
```

---

#### GET /api/v1/multipliers/:id

Get a specific multiplier configuration.

**Response:** `200 OK` - Returns single multiplier object

---

#### POST /api/v1/multipliers

Create a new multiplier configuration.

**Response:** `201 Created` - Returns created multiplier object

---

#### PUT /api/v1/multipliers/:id

Update an existing multiplier configuration.

**Response:** `200 OK` - Returns updated multiplier object

---

### Business Profile Endpoints

#### GET /api/v1/business-profiles

Get list of all business profiles.

**Response:** `200 OK`
```json
[
  {
    "ecdd_business_profile_pk": "b1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c",
    "country_id": 134,
    "state_id": null,
    "risk_status_id": 2,
    "average_deposit": 5000.00,
    "deposit_multiplier": 2.00,
    "time_period_days": 28,
    "enabled": true,
    "created_at": "2026-01-01T00:00:00Z",
    "updated_at": "2026-01-01T00:00:00Z",
    "updated_by": "system@company.com"
  }
]
```

---

#### GET /api/v1/business-profiles/:id

Get a specific business profile.

**Response:** `200 OK` - Returns single business profile object

---

#### POST /api/v1/business-profiles

Create a new business profile.

**Response:** `201 Created` - Returns created business profile object

---

#### PUT /api/v1/business-profiles/:id

Update an existing business profile.

**Response:** `200 OK` - Returns updated business profile object

---

### User-Folder Assignment Endpoints

#### GET /api/v1/user-folders

Get list of all user-folder assignments.

**Response:** `200 OK`
```json
[
  {
    "ecdd_user_case_management_folder_pk": "uf1a2b3c-d4e5-6f7a-8b9c-0d1e2f3a4b5c",
    "case_management_folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
    "user_status_pk": "u001-malta-134-001",
    "logged_at": "2026-01-20T10:00:00Z",
    "updated_by": "compliance@company.com"
  }
]
```

---

#### POST /api/v1/user-folders

Assign a user to a folder.

**Request Body:**
```json
{
  "case_management_folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
  "user_status_pk": "u001-malta-134-001",
  "updated_by": "admin@company.com"
}
```

**Response:** `201 Created` - Returns created assignment object

---

#### DELETE /api/v1/user-folders/:id

Remove a user-folder assignment.

**Response:** `200 OK`
```json
{
  "message": "Assignment deleted successfully"
}
```

---

## React Integration Examples

### Complete React Service Module

```javascript
// src/services/ecddApi.js
const BASE_URL = 'http://localhost:3001';

// Generic fetch wrapper
const apiFetch = async (endpoint, options = {}) => {
  const url = `${BASE_URL}${endpoint}`;
  const config = {
    headers: {
      'Content-Type': 'application/json',
      ...options.headers
    },
    ...options
  };

  const response = await fetch(url, config);
  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || `HTTP ${response.status}`);
  }
  return await response.json();
};

// Health Check
export const checkHealth = () => apiFetch('/health');

// User Operations
export const userAPI = {
  getAll: () => apiFetch('/api/v1/users'),
  getById: (id) => apiFetch(`/api/v1/users/${id}`),
  create: (userData) => apiFetch('/api/v1/users', {
    method: 'POST',
    body: JSON.stringify(userData)
  }),
  update: (id, updates) => apiFetch(`/api/v1/users/${id}`, {
    method: 'PUT',
    body: JSON.stringify(updates)
  }),
  getByRegion: async (countryIds) => {
    const users = await apiFetch('/api/v1/users');
    return users.filter(user => countryIds.includes(user.country_id));
  }
};

// Threshold Operations
export const thresholdAPI = {
  getAll: () => apiFetch('/api/v1/thresholds'),
  getById: (id) => apiFetch(`/api/v1/thresholds/${id}`),
  create: (data) => apiFetch('/api/v1/thresholds', {
    method: 'POST',
    body: JSON.stringify(data)
  }),
  update: (id, updates) => apiFetch(`/api/v1/thresholds/${id}`, {
    method: 'PUT',
    body: JSON.stringify(updates)
  }),
  delete: (id) => apiFetch(`/api/v1/thresholds/${id}`, { method: 'DELETE' })
};

// Folder Operations
export const folderAPI = {
  getAll: () => apiFetch('/api/v1/folders'),
  getById: (id) => apiFetch(`/api/v1/folders/${id}`),
  create: (folderName, updatedBy
) => apiFetch('/api/v1/folders', {
    method: 'POST',
    body: JSON.stringify({ folder_name: folderName, updated_by: updatedBy })
  }),
  update: (id, updates) => apiFetch(`/api/v1/folders/${id}`, {
    method: 'PUT',
    body: JSON.stringify(updates)
  })
};

// Multiplier Operations
export const multiplierAPI = {
  getAll: () => apiFetch('/api/v1/multipliers'),
  getById: (id) => apiFetch(`/api/v1/multipliers/${id}`),
  create: (data) => apiFetch('/api/v1/multipliers', {
    method: 'POST',
    body: JSON.stringify(data)
  }),
  update: (id, updates) => apiFetch(`/api/v1/multipliers/${id}`, {
    method: 'PUT',
    body: JSON.stringify(updates)
  })
};

// Business Profile Operations
export const businessProfileAPI = {
  getAll: () => apiFetch('/api/v1/business-profiles'),
  getById: (id) => apiFetch(`/api/v1/business-profiles/${id}`),
  create: (data) => apiFetch('/api/v1/business-profiles', {
    method: 'POST',
    body: JSON.stringify(data)
  }),
  update: (id, updates) => apiFetch(`/api/v1/business-profiles/${id}`, {
    method: 'PUT',
    body: JSON.stringify(updates)
  })
};

// User-Folder Assignment Operations
export const userFolderAPI = {
  getAll: () => apiFetch('/api/v1/user-folders'),
  assign: (folderId, userId, updatedBy) => apiFetch('/api/v1/user-folders', {
    method: 'POST',
    body: JSON.stringify({
      case_management_folder_pk: folderId,
      user_status_pk: userId,
      updated_by: updatedBy
    })
  }),
  remove: (id) => apiFetch(`/api/v1/user-folders/${id}`, { method: 'DELETE' })
};
```

---

## Field Reference Guide

### User Status Fields

| Field | Type | Description | Values |
|-------|------|-------------|--------|
| `ecdd_status` | integer | Current ECDD status | 1=Active, 2=Under Review, 3=Flagged, 4=Suspended, 5=Blocked |
| `ecdd_threshold` | decimal | Monitoring threshold amount | GBP value |
| `ecdd_review_trigger` | integer | Review trigger count | Number of triggers |
| `ecdd_multiplier` | decimal | Applied multiplier | 0.5, 1.0, 1.5, etc. |
| `ecdd_multiplier_rg_flag` | boolean | RG multiplier applied | true/false |
| `info_source` | integer | Information source | 1-10 (various sources) |
| `sign_off_status` | integer | Sign-off status | 1=Pending, 2=Approved, 3=Rejected |
| `ecdd_rg_review_status` | integer | RG review status | 1=Not Required, 2=Pending, 3=In Progress, 4=Complete |
| `ecdd_report_status` | integer | Report status | 1=Not Started, 2=In Progress, 3=Complete, 4=Submitted |
| `ecdd_review_status` | integer | Review status | 1=Not Required, 2=Pending, 3=In Progress, 4=Complete |
| `ecdd_document_status` | integer | Documentation status | 1=Not Required, 2=Pending, 3=Received, 4=Verified |
| `ecdd_escalation_status` | integer | Escalation status | 1=None, 2=Level 1, 3=Level 2 |
| `uar_status` | integer | UAR status | 1=Current, 2=Pending Update, 3=Overdue |

### Threshold Type Values

| Type ID | Description |
|---------|-------------|
| 1 | Deposit |
| 2 | Net Deposit |
| 3 | Stakes |

### Period Values

| Period ID | Description |
|-----------|-------------|
| 1 | 24 hours |
| 2 | 28 days |
| 3 | 84 days |
| 4 | 91 days |
| 5 | 182 days |
| 6 | 365 days |

### Customer Risk Level Values

| Risk Level ID | Description |
|---------------|-------------|
| 1 | Low |
| 2 | Medium |
| 3 | Medium-High |
| 4 | High |

---

## TypeScript Types

```typescript
// src/types/ecdd.ts

export interface ECDDUserStatus {
  ecdd_user_status_pk: string;
  user_id: number;
  user_name: string;
  country_id: number;
  state_id: number | null;
  ecdd_status: number;
  ecdd_threshold: number;
  ecdd_review_trigger: number;
  ecdd_suspension_due_date: string | null;
  ecdd_multiplier: number;
  ecdd_multiplier_rg_flag: boolean;
  user_lt_net_deposit_threshold_gbp: number;
  user_lt_deposit_threshold_gbp: number;
  user_12month_net_deposit_threshold_gbp: number;
  info_source: number;
  sign_off_status: number;
  date_last_ecdd_sign_off: string | null;
  ecdd_rg_review_status: number;
  date_last_ecdd_rg_sign_off: string | null;
  ecdd_report_status: number;
  ecdd_review_status: number;
  ecdd_document_status: number;
  ecdd_escalation_status: number;
  uar_status: number;
  logged_at: string;
  updated_by: string;
}

export interface ECDDThresholdConfig {
  ecdd_threshold_config_pk: string;
  title: string;
  is_active: boolean;
  country_id: number;
  state_id: number | null;
  type: number;
  value: number;
  currency_id: number;
  period: number;
  use_multipliers: boolean;
  use_rg_flag: boolean;
  apply_all_statuses: boolean;
  backfill: boolean;
  hierarchy: number;
  ecdd_status: number;
  ecdd_review_status: number;
  ecdd_report_status: number;
  sign_off_status: number;
  customer_risk_level: number;
  ndl_28_day_gbp: number;
  ndl_monthly_gbp: number;
  case_management_folder_pk: string | null;
  logged_at: string;
  updated_by: string;
}

export interface ECDDCaseManagementFolder {
  ecdd_case_management_folder_pk: string;
  folder_name: string;
  logged_at: string;
  updated_by: string;
}

export interface ECDDUserCaseManagementFolder {
  ecdd_user_case_management_folder_pk: string;
  case_management_folder_pk: string;
  user_status_pk: string;
  logged_at: string;
  updated_by: string;
}

export interface Region {
  name: string;
  countryIds: number[];
  code: string;
  currency: string;
}

export const REGIONS: Record<string, Region> = {
  MALTA: {
    name: 'Malta',
    countryIds: [134, 135, 136],
    code: 'MT',
    currency: 'EUR'
  },
  USA: {
    name: 'United States',
    countryIds: [231, 232, 233],
    code: 'US',
    currency: 'USD'
  },
  AUSTRALIA: {
    name: 'Australia',
    countryIds: [13, 14, 15],
    code: 'AU',
    currency: 'AUD'
  },
  GIBRALTAR: {
    name: 'Gibraltar',
    countryIds: [83, 84, 85],
    code: 'GI',
    currency: 'GBP'
  }
};
```

---

## Support & Resources

- **Mock Data Location**: `data/*.json`
- **Server Port**: 3001
- **Protocol**: HTTP (no HTTPS in mock)
- **Authentication**: None (mock environment)

---

**Last Updated**: February 12, 2026  
**API Version**: v1  
**Documentation Version**: 2.0.0 (No Pagination)