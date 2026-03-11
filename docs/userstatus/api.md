# ECDD User Status API

The ECDD User Status API provides endpoints for managing ECDD user statuses. It supports listing user statuses with filtering, sorting, and pagination, as well as update, patch, delete operations, and retrieving a user status's assigned folders.

- [Get All User Statuses](#get-all-user-statuses)
- [Get User Status By PK](#get-user-status-by-pk)
- [Update User Status](#update-user-status)
- [Patch User Status](#patch-user-status)
- [Delete User Status](#delete-user-status)
- [Get User Status Folders](#get-user-status-folders)

---

## Get All User Statuses

### Request

GET

> http://`<ecddapi>`/api/ecdd/userstatus

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### URL Params

| query string parameter | description | type | example | required |
|----|----|----|----|----|
| country_id | country ID to filter by | int | 1 | false |
| ecdd_status | ECDD status to filter by | int | 2 | false |
| user_name | filter by user name (case-insensitive match) | string | johnsmith | false |
| user_id | filter by user ID | int | 12345 | false |
| sort_by | field to sort results by | string | user_id | false |
| sort_dir | sort direction | string | asc | false |
| page | page number for pagination (starts at 1) | int | 1 | false |
| page_size | number of results per page (max 100, default 20) | int | 20 | false |

**sort_by valid values:** user_id, user_name, country_id, ecdd_status, ecdd_threshold, ecdd_multiplier, logged_at, ecdd_user_status_pk

**sort_dir valid values:** asc, desc (default: asc)

### Response 200

> HTTP status 200 OK

#### Body (without pagination)

content-type: application/json

Returns an array of user status objects.

| parameter | type | description | example | required |
|----|----|----|----|----|
| ecdd_user_status_pk | string | spanner generated UUID primary key | "f9e8d7c6-b5a4-3210-fedc-ba9876543210" | true |
| user_id | int | user identifier from the core user system | 100001 | true |
| user_name | string | display name of the user | "Matthew Borg" | true |
| country_id | int | country identifier reference | 134 | true |
| language | string | ISO 639-1 language code | "EN" | true |
| state_id | int | state identifier reference (nullable) | 5 | false |
| ecdd_status | int | ECDD status (1=Not Required, 2=In progress, 3=Complete, 4=Suspended-Manual, 5=Suspended-Auto, 6=Closed, 7=Block Process) | 2 | true |
| ecdd_threshold | float | threshold value | 14031.98 | true |
| ecdd_review_trigger | int | review trigger type identifier | 3 | true |
| ecdd_suspension_due_date | string | suspension due date (RFC3339, nullable) | "2026-02-15T00:00:00Z" | false |
| ecdd_multiplier | float | multiplier value applied to thresholds | 1.00 | true |
| ecdd_multiplier_rg_flag | bool | RG (Responsible Gambling) flag indicator | false | true |
| user_lt_engg_threshold_gbp | float | lifetime engagement threshold in GBP | 2047.05 | true |
| user_lt_deposit_threshold_gbp | float | lifetime deposit threshold in GBP | 4599.86 | true |
| user_12month_drop_threshold_gbp | float | 12-month drop threshold in GBP | 1104.47 | true |
| info_source | int | information source identifier | 1 | true |
| sign_off_status | int | sign-off status indicator | 1 | true |
| date_last_ecdd_sign_off | string | last ECDD sign-off date (RFC3339, nullable) | "2026-01-10T00:00:00Z" | false |
| ecdd_rg_review_status | int | RG review status indicator | 1 | true |
| date_last_ecdd_rg_sign_off | string | last RG sign-off date (RFC3339, nullable) | "2026-01-08T00:00:00Z" | false |
| ecdd_report_status | int | report status indicator | 1 | true |
| ecdd_review_status | int | review status indicator | 1 | true |
| ecdd_document_status | int | document status indicator | 1 | true |
| ecdd_escalation_status | int | escalation status indicator | 1 | true |
| uar_status | int | UAR (Unusual Activity Report) status indicator | 1 | true |
| logged_at | string | timestamp when the row was last logged/modified (RFC3339) | "2026-02-01T09:00:00Z" | true |
| updated_by | string | username of the person who last updated this record | "malta.compliance@company.com" | true |

##### Example

```json
[
    {
        "ecdd_user_status_pk": "u001-malta-134-001",
        "user_id": 100001,
        "user_name": "Matthew Borg",
        "country_id": 134,
        "language": "EN",
        "ecdd_status": 1,
        "ecdd_threshold": 14031.98,
        "ecdd_review_trigger": 3,
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
]
```

#### Body (with pagination)

content-type: application/json

When `page` or `page_size` query parameters are provided, the response is wrapped in a pagination envelope.

| parameter | type | description | example | required |
|----|----|----|----|----|
| data | array | array of user status objects (see above) | [...] | true |
| page | int | current page number | 1 | true |
| page_size | int | number of results per page | 20 | true |
| total_count | int | total number of matching records | 120 | true |
| total_pages | int | total number of pages | 6 | true |

##### Example

```json
{
    "data": [
        {
            "ecdd_user_status_pk": "u001-malta-134-001",
            "user_id": 100001,
            "user_name": "Matthew Borg",
            "country_id": 134,
            "language": "EN",
            "ecdd_status": 1,
            "ecdd_threshold": 14031.98,
            "ecdd_review_trigger": 3,
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
    ],
    "page": 1,
    "page_size": 20,
    "total_count": 120,
    "total_pages": 6
}
```

### Response 405

> HTTP status 405 Method Not Allowed

#### Body

content-type: application/json

##### Example

```json
{"error": "Method not allowed"}
```

---

## Get User Status By PK

### Request

GET

> http://`<ecddapi>`/api/ecdd/userstatus/{userstatuspk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| userstatuspk | the ecdd_user_status_pk of the user status record | string | u001-malta-134-001 | true |

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns a single user status object (see [Get All User Statuses](#get-all-user-statuses) response body for field definitions).

##### Example

```json
{
    "ecdd_user_status_pk": "u001-malta-134-001",
    "user_id": 100001,
    "user_name": "Matthew Borg",
    "country_id": 134,
    "language": "EN",
    "ecdd_status": 1,
    "ecdd_threshold": 14031.98,
    "ecdd_review_trigger": 3,
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

### Response 404

> HTTP status 404 Not Found

#### Body

content-type: application/json

##### Example

```json
{"error": "User not found"}
```

### Response 405

> HTTP status 405 Method Not Allowed

#### Body

content-type: application/json

##### Example

```json
{"error": "Method not allowed"}
```

---

## Update User Status

### Request

PUT

> http://`<ecddapi>`/api/ecdd/userstatus/{userstatuspk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| userstatuspk | the ecdd_user_status_pk of the user status record | string | u001-malta-134-001 | true |

#### Body

content-type: application/json

Full user status object. All fields are replaced.

| parameter | type | description | example | required |
|----|----|----|----|----|
| user_id | int | user identifier from the core user system | 100001 | true |
| user_name | string | display name of the user | "Matthew Borg" | true |
| country_id | int | country identifier reference | 134 | true |
| language | string | ISO 639-1 language code | "EN" | true |
| state_id | int | state identifier reference (nullable) | 5 | false |
| ecdd_status | int | ECDD status code | 3 | true |
| ecdd_threshold | float | threshold value | 20000.00 | true |
| ecdd_review_trigger | int | review trigger type identifier | 3 | true |
| ecdd_suspension_due_date | string | suspension due date (RFC3339, nullable) | "2026-02-15T00:00:00Z" | false |
| ecdd_multiplier | float | multiplier value applied to thresholds | 1.00 | true |
| ecdd_multiplier_rg_flag | bool | RG flag indicator | false | true |
| user_lt_engg_threshold_gbp | float | lifetime engagement threshold in GBP | 5000.00 | true |
| user_lt_deposit_threshold_gbp | float | lifetime deposit threshold in GBP | 10000.00 | true |
| user_12month_drop_threshold_gbp | float | 12-month drop threshold in GBP | 2500.00 | true |
| info_source | int | information source identifier | 1 | true |
| sign_off_status | int | sign-off status indicator | 1 | true |
| date_last_ecdd_sign_off | string | last ECDD sign-off date (RFC3339, nullable) | "2026-01-20T00:00:00Z" | false |
| ecdd_rg_review_status | int | RG review status indicator | 1 | true |
| date_last_ecdd_rg_sign_off | string | last RG sign-off date (RFC3339, nullable) | "2026-01-18T00:00:00Z" | false |
| ecdd_report_status | int | report status indicator | 1 | true |
| ecdd_review_status | int | review status indicator | 1 | true |
| ecdd_document_status | int | document status indicator | 1 | true |
| ecdd_escalation_status | int | escalation status indicator | 1 | true |
| uar_status | int | UAR status indicator | 1 | true |
| updated_by | string | username of the person updating this record | "malta.compliance@company.com" | true |

##### Example

```json
{
    "user_id": 100001,
    "user_name": "Matthew Borg",
    "country_id": 134,
    "language": "EN",
    "ecdd_status": 3,
    "ecdd_threshold": 20000.00,
    "ecdd_review_trigger": 3,
    "ecdd_multiplier": 1.00,
    "ecdd_multiplier_rg_flag": false,
    "user_lt_engg_threshold_gbp": 5000.00,
    "user_lt_deposit_threshold_gbp": 10000.00,
    "user_12month_drop_threshold_gbp": 2500.00,
    "info_source": 1,
    "sign_off_status": 1,
    "ecdd_rg_review_status": 1,
    "ecdd_report_status": 1,
    "ecdd_review_status": 1,
    "ecdd_document_status": 1,
    "ecdd_escalation_status": 1,
    "uar_status": 1,
    "updated_by": "malta.compliance@company.com"
}
```

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns the updated user status object.

### Response 400

> HTTP status 400 Bad Request

#### Body

content-type: application/json

##### Example

```json
{"error": "Invalid request body"}
```

### Response 404

> HTTP status 404 Not Found

#### Body

content-type: application/json

##### Example

```json
{"error": "User not found"}
```

### Response 405

> HTTP status 405 Method Not Allowed

#### Body

content-type: application/json

##### Example

```json
{"error": "Method not allowed"}
```

---

## Patch User Status

### Request

PATCH

> http://`<ecddapi>`/api/ecdd/userstatus/{userstatuspk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| userstatuspk | the ecdd_user_status_pk of the user status record | string | u001-malta-134-001 | true |

#### Body

content-type: application/json

Partial user status object. Only provided fields are updated; omitted fields retain their existing values. See [Update User Status](#update-user-status) request body for field definitions.

##### Example

```json
{
    "ecdd_status": 3,
    "ecdd_threshold": 20000.00
}
```

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns the full updated user status object.

### Response 400

> HTTP status 400 Bad Request

#### Body

content-type: application/json

##### Example

```json
{"error": "Invalid request body"}
```

### Response 404

> HTTP status 404 Not Found

#### Body

content-type: application/json

##### Example

```json
{"error": "User not found"}
```

### Response 405

> HTTP status 405 Method Not Allowed

#### Body

content-type: application/json

##### Example

```json
{"error": "Method not allowed"}
```

---

## Delete User Status

### Request

DELETE

> http://`<ecddapi>`/api/ecdd/userstatus/{userstatuspk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| userstatuspk | the ecdd_user_status_pk of the user status record | string | u001-malta-134-001 | true |

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

| parameter | type | description | example | required |
|----|----|----|----|----|
| message | string | confirmation message including cascade-deleted folder assignment count | "User deleted successfully. 3 folder assignment(s) cascade-deleted." | true |

##### Example

```json
{
    "message": "User deleted successfully. 3 folder assignment(s) cascade-deleted."
}
```

### Response 404

> HTTP status 404 Not Found

#### Body

content-type: application/json

##### Example

```json
{"error": "User not found"}
```

### Response 405

> HTTP status 405 Method Not Allowed

#### Body

content-type: application/json

##### Example

```json
{"error": "Method not allowed"}
```

---

## Get User Status Folders

### Request

GET

> http://`<ecddapi>`/api/ecdd/userstatus/{userstatuspk}/folders

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| userstatuspk | the ecdd_user_status_pk of the user status record | string | u001-malta-134-001 | true |

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns an array of case management folder objects assigned to the user status. Returns an empty array if the user status has no folder assignments.

| parameter | type | description | example | required |
|----|----|----|----|----|
| ecdd_case_management_folder_pk | string | primary key of the folder | "folder-high-risk-001" | true |
| folder_name | string | name of the folder | "High Risk" | true |
| logged_at | string | record logged date (RFC3339) | "2026-01-15T10:00:00Z" | true |
| updated_by | string | last updated by | "compliance.admin@company.com" | true |

##### Example

```json
[
    {
        "ecdd_case_management_folder_pk": "folder-high-risk-001",
        "folder_name": "High Risk",
        "logged_at": "2026-01-15T10:00:00Z",
        "updated_by": "compliance.admin@company.com"
    },
    {
        "ecdd_case_management_folder_pk": "folder-medium-risk-002",
        "folder_name": "Medium Risk",
        "logged_at": "2026-01-15T10:00:00Z",
        "updated_by": "compliance.admin@company.com"
    }
]
```

### Response 400

> HTTP status 400 Bad Request

#### Body

content-type: application/json

##### Example

```json
{"error": "User ID required"}
```

### Response 404

> HTTP status 404 Not Found

#### Body

content-type: application/json

##### Example

```json
{"error": "User not found"}
```

### Response 405

> HTTP status 405 Method Not Allowed

#### Body

content-type: application/json

##### Example

```json
{"error": "Method not allowed"}
```
