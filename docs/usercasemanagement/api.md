# ECDD User Case Management API

The ECDD User Case Management API provides endpoints for managing user-to-folder assignments in the ECDD system. It supports listing assignments with filtering, sorting, and pagination, creating and deleting assignments (single and bulk), retrieving users assigned to folders, retrieving folders assigned to a user, and folder assignment statistics. All operations target the ECDDUserCaseManagementFolder Spanner table.

- [Get All Assignments](#get-all-assignments)
- [Create Assignment](#create-assignment)
- [Delete Assignment By PK](#delete-assignment-by-pk)
- [Delete Assignments By Folder](#delete-assignments-by-folder)
- [Get Folder Users](#get-folder-users)
- [Get User Folders](#get-user-folders)
- [Remove User From Folder](#remove-user-from-folder)
- [Bulk Remove Users From Folder](#bulk-remove-users-from-folder)
- [Bulk Add Users To Folder](#bulk-add-users-to-folder)
- [Get Folder Stats](#get-folder-stats)
- [Get All Folder Stats](#get-all-folder-stats)

---

## Get All Assignments

### Request

GET

> http://`<ecddapi>`/api/ecdd/usercasemanagement

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### URL Params

| query string parameter | description | type | example | required |
|----|----|----|----|----|
| case_management_folder_pk | folder primary key to filter by | string | a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d | false |
| user_status_pk | user status primary key to filter by | string | u001-malta-134-001 | false |
| sort_by | field to sort results by | string | logged_at | false |
| sort_dir | sort direction | string | asc | false |
| page | page number for pagination (starts at 1) | int | 1 | false |
| page_size | number of results per page (max 100, default 20) | int | 20 | false |

**sort_by valid values:** ecdd_user_case_management_folder_pk, case_management_folder_pk, user_status_pk, logged_at

**sort_dir valid values:** asc, desc (default: asc)

**Filtering:** `case_management_folder_pk` and `user_status_pk` can be combined (AND logic).

### Response 200

> HTTP status 200 OK

#### Body (without pagination)

content-type: application/json

Returns an array of assignment objects. Returns an empty array when no results match.

| parameter | type | description | example | required |
|----|----|----|----|----|
| ecdd_user_case_management_folder_pk | string | spanner generated UUID primary key | "604f0583-5016-48fb-8ed2-113121c204c4" | true |
| case_management_folder_pk | string | foreign key referencing ECDDCaseManagementFolder | "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d" | true |
| user_status_pk | string | foreign key referencing ECDDUserStatus | "u001-malta-134-001" | true |
| logged_at | string | timestamp when the row was last logged/modified (RFC3339) | "2026-02-01T09:00:00Z" | true |
| updated_by | string | username of the person who last updated this record | "malta.compliance@company.com" | true |

##### Example

```json
[
    {
        "ecdd_user_case_management_folder_pk": "604f0583-5016-48fb-8ed2-113121c204c4",
        "case_management_folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
        "user_status_pk": "u001-malta-134-001",
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
| data | array | array of assignment objects (see above) | [...] | true |
| page | int | current page number | 1 | true |
| page_size | int | number of results per page | 20 | true |
| total_count | int | total number of matching records | 50 | true |
| total_pages | int | total number of pages | 3 | true |

##### Example

```json
{
    "data": [
        {
            "ecdd_user_case_management_folder_pk": "604f0583-5016-48fb-8ed2-113121c204c4",
            "case_management_folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
            "user_status_pk": "u001-malta-134-001",
            "logged_at": "2026-02-01T09:00:00Z",
            "updated_by": "malta.compliance@company.com"
        }
    ],
    "page": 1,
    "page_size": 20,
    "total_count": 50,
    "total_pages": 3
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

## Create Assignment

### Request

POST

> http://`<ecddapi>`/api/ecdd/usercasemanagement

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Body

content-type: application/json

| parameter | type | description | example | required |
|----|----|----|----|----|
| case_management_folder_pk | string | foreign key referencing ECDDCaseManagementFolder | "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d" | true |
| user_status_pk | string | foreign key referencing ECDDUserStatus | "u001-malta-134-001" | true |
| updated_by | string | username of the person creating this record | "malta.compliance@company.com" | true |

##### Example

```json
{
    "case_management_folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
    "user_status_pk": "u001-malta-134-001",
    "updated_by": "malta.compliance@company.com"
}
```

### Response 201

> HTTP status 201 Created

#### Body

content-type: application/json

Returns the created assignment object with server-generated fields.

| parameter | type | description | example | required |
|----|----|----|----|----|
| ecdd_user_case_management_folder_pk | string | spanner generated UUID primary key | "604f0583-5016-48fb-8ed2-113121c204c4" | true |
| case_management_folder_pk | string | foreign key referencing ECDDCaseManagementFolder | "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d" | true |
| user_status_pk | string | foreign key referencing ECDDUserStatus | "u001-malta-134-001" | true |
| logged_at | string | timestamp when the row was created (RFC3339) | "2026-02-01T09:00:00Z" | true |
| updated_by | string | username of the person who created this record | "malta.compliance@company.com" | true |

##### Example

```json
{
    "ecdd_user_case_management_folder_pk": "604f0583-5016-48fb-8ed2-113121c204c4",
    "case_management_folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
    "user_status_pk": "u001-malta-134-001",
    "logged_at": "2026-02-01T09:00:00Z",
    "updated_by": "malta.compliance@company.com"
}
```

### Response 400

> HTTP status 400 Bad Request

#### Body

content-type: application/json

##### Example

```json
{"error": "Invalid request body"}
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

## Delete Assignment By PK

### Request

DELETE

> http://`<ecddapi>`/api/ecdd/usercasemanagement/{usercasemanagementpk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| usercasemanagementpk | the ecdd_user_case_management_folder_pk of the assignment record | string | 604f0583-5016-48fb-8ed2-113121c204c4 | true |

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

| parameter | type | description | example | required |
|----|----|----|----|----|
| message | string | confirmation message | "Assignment deleted successfully" | true |

##### Example

```json
{
    "message": "Assignment deleted successfully"
}
```

### Response 404

> HTTP status 404 Not Found

#### Body

content-type: application/json

##### Example

```json
{"error": "Assignment not found"}
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

## Delete Assignments By Folder

### Request

DELETE

> http://`<ecddapi>`/api/ecdd/usercasemanagement?case_management_folder_pk={case_management_folder_pk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### URL Params

| query string parameter | description | type | example | required |
|----|----|----|----|----|
| case_management_folder_pk | the folder primary key whose assignments should be deleted | string | a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d | true |

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

| parameter | type | description | example | required |
|----|----|----|----|----|
| message | string | confirmation message | "Assignments deleted successfully" | true |
| count | int | number of assignment records deleted | 3 | true |

##### Example

```json
{
    "message": "Assignments deleted successfully",
    "count": 3
}
```

### Response 400

> HTTP status 400 Bad Request

#### Body

content-type: application/json

##### Example

```json
{"error": "case_management_folder_pk query parameter is required"}
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

## Get Folder Users

### Request

GET

> http://`<ecddapi>`/api/ecdd/usercasemanagement/folder/{folderpk}/users

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| folderpk | the ecdd_case_management_folder_pk of the folder record | string | a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d | true |

#### URL Params

| query string parameter | description | type | example | required |
|----|----|----|----|----|
| language | filter by language code | string | EN | false |
| country_id | filter by country ID | int | 134 | false |
| region | filter by region name | string | MALTA | false |
| sort_by | field to sort results by | string | user_id | false |
| sort_dir | sort direction | string | asc | false |
| page | page number for pagination (starts at 1, default 1) | int | 1 | false |
| page_size | number of results per page (default 10) | int | 10 | false |

**sort_by valid values:** user_id, user_name, country_id, ecdd_status, ecdd_threshold, ecdd_multiplier, logged_at, ecdd_user_status_pk

**sort_dir valid values:** asc, desc (default: asc)

**region valid values:** MALTA, GIBRALTAR, USA, AUSTRALIA

**Note:** This endpoint always returns a paginated response (default page=1, pageSize=10).

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns a paginated list of user status objects assigned to the folder. See [ECDD User Status API](../userstatus/api.md) for user status field definitions.

| parameter | type | description | example | required |
|----|----|----|----|----|
| data | array | array of user status objects | [...] | true |
| page | int | current page number | 1 | true |
| page_size | int | number of results per page | 10 | true |
| total_count | int | total number of matching records | 42 | true |
| total_pages | int | total number of pages | 5 | true |

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
            "user_lt_net_deposit_threshold_gbp": 2047.05,
            "user_lt_deposit_threshold_gbp": 4599.86,
            "user_12month_net_deposit_threshold_gbp": 1104.47,
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
    "page_size": 10,
    "total_count": 42,
    "total_pages": 5
}
```

### Response 404

> HTTP status 404 Not Found

#### Body

content-type: application/json

##### Example

```json
{"error": "Folder not found"}
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

## Get User Folders

### Request

GET

> http://`<ecddapi>`/api/ecdd/usercasemanagement/user/{userstatuspk}/folders

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

## Remove User From Folder

### Request

DELETE

> http://`<ecddapi>`/api/ecdd/usercasemanagement/folder/{folderpk}/users/{userstatuspk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| folderpk | the ecdd_case_management_folder_pk of the folder record | string | a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d | true |
| userstatuspk | the ecdd_user_status_pk of the user status record | string | u001-malta-134-001 | true |

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

| parameter | type | description | example | required |
|----|----|----|----|----|
| message | string | confirmation message | "User removed from folder successfully" | true |

##### Example

```json
{
    "message": "User removed from folder successfully"
}
```

### Response 404

> HTTP status 404 Not Found

#### Body

content-type: application/json

##### Example

```json
{"error": "User-folder assignment not found"}
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

## Bulk Remove Users From Folder

### Request

POST

> http://`<ecddapi>`/api/ecdd/usercasemanagement/folder/{folderpk}/users/bulk-delete

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| folderpk | the ecdd_case_management_folder_pk of the folder record | string | a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d | true |

#### Body

content-type: application/json

| parameter | type | description | example | required |
|----|----|----|----|----|
| user_status_pks | array of string | list of ecdd_user_status_pk values to remove | ["pk1", "pk2"] | true |
| updated_by | string | username of the person performing the operation | "admin" | true |

##### Example

```json
{
    "user_status_pks": ["u001-malta-134-001", "u002-malta-134-002"],
    "updated_by": "admin"
}
```

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

| parameter | type | description | example | required |
|----|----|----|----|----|
| message | string | summary message | "Users removed from folder successfully" | true |
| deleted_count | int | number of assignments successfully deleted | 2 | true |
| failed_count | int | number of assignments that failed to delete | 0 | true |
| failed_user_status_pks | array of string | list of pks that failed | [] | true |

##### Example

```json
{
    "message": "Users removed from folder successfully",
    "deleted_count": 2,
    "failed_count": 0,
    "failed_user_status_pks": []
}
```

### Response 400

> HTTP status 400 Bad Request

#### Body

content-type: application/json

##### Example

```json
{"error": "Invalid request body"}
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

## Bulk Add Users To Folder

### Request

POST

> http://`<ecddapi>`/api/ecdd/usercasemanagement/folder/{folderpk}/users/bulk-add

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| folderpk | the ecdd_case_management_folder_pk of the folder record | string | a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d | true |

#### Body

content-type: application/json

| parameter | type | description | example | required |
|----|----|----|----|----|
| user_ids | array of string | list of user status PKs to add to the folder | ["pk1", "pk2"] | true |
| updated_by | string | username of the person performing the operation | "admin" | true |

##### Example

```json
{
    "user_ids": ["u003-malta-134-003", "u004-malta-134-004"],
    "updated_by": "admin"
}
```

### Response 201

> HTTP status 201 Created

#### Body

content-type: application/json

| parameter | type | description | example | required |
|----|----|----|----|----|
| message | string | summary message | "Users assigned to folder successfully" | true |
| created_count | int | number of assignments successfully created | 2 | true |
| skipped_count | int | number of assignments skipped (already exist) | 0 | true |
| skipped_user_ids | array of string | list of user ids that were skipped | [] | true |

##### Example

```json
{
    "message": "Users assigned to folder successfully",
    "created_count": 2,
    "skipped_count": 0,
    "skipped_user_ids": []
}
```

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
{"error": "Folder not found"}
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

## Get Folder Stats

### Request

GET

> http://`<ecddapi>`/api/ecdd/usercasemanagement/folder/{folderpk}/stats

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| folderpk | the ecdd_case_management_folder_pk of the folder record | string | a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d | true |

#### URL Params

| query string parameter | description | type | example | required |
|----|----|----|----|----|
| language | filter by language code | string | EN | false |
| country_id | filter by country ID | int | 134 | false |
| region | filter by region name | string | MALTA | false |

**region valid values:** MALTA, GIBRALTAR, USA, AUSTRALIA

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns a single folder stats object.

| parameter | type | description | example | required |
|----|----|----|----|----|
| folder_pk | string | primary key of the folder | "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d" | true |
| folder_name | string | name of the folder | "High Risk Customers" | true |
| region | string | region name for the folder | "MALTA" | true |
| user_count | int | number of users assigned to the folder | 42 | true |
| oldest_user_date | string | logged_at timestamp of the oldest user assignment (RFC3339, nullable) | "2026-01-15T10:00:00Z" | false |

##### Example

```json
{
    "folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
    "folder_name": "High Risk Customers",
    "region": "MALTA",
    "user_count": 42,
    "oldest_user_date": "2026-01-15T10:00:00Z"
}
```

### Response 404

> HTTP status 404 Not Found

#### Body

content-type: application/json

##### Example

```json
{"error": "Folder not found"}
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

## Get All Folder Stats

### Request

GET

> http://`<ecddapi>`/api/ecdd/usercasemanagement/stats

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### URL Params

| query string parameter | description | type | example | required |
|----|----|----|----|----|
| language | filter by language code | string | EN | false |
| country_id | filter by country ID | int | 134 | false |
| region | filter by region name | string | MALTA | false |
| sort_by | field to sort results by | string | folder_name | false |
| sort_dir | sort direction | string | asc | false |
| page | page number for pagination (starts at 1) | int | 1 | false |
| page_size | number of results per page (max 100, default 20) | int | 20 | false |

**sort_by valid values:** folder_name, folder_pk, user_count

**sort_dir valid values:** asc, desc (default: asc)

**region valid values:** MALTA, GIBRALTAR, USA, AUSTRALIA

**Note:** When `region` is specified, only folders belonging to that region are returned. Both folders and their user counts are filtered to the specified region.

### Response 200

> HTTP status 200 OK

#### Body (without pagination)

content-type: application/json

Returns an array of folder stats objects.

| parameter | type | description | example | required |
|----|----|----|----|----|
| folder_pk | string | primary key of the folder | "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d" | true |
| folder_name | string | name of the folder | "High Risk Customers" | true |
| region | string | region name for the folder | "MALTA" | true |
| user_count | int | number of users assigned to the folder | 42 | true |
| oldest_user_date | string | logged_at timestamp of the oldest user assignment (RFC3339, nullable) | "2026-01-15T10:00:00Z" | false |

##### Example

```json
[
    {
        "folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
        "folder_name": "High Risk Customers",
        "region": "MALTA",
        "user_count": 42,
        "oldest_user_date": "2026-01-15T10:00:00Z"
    }
]
```

#### Body (with pagination)

content-type: application/json

When `page` or `page_size` query parameters are provided, the response is wrapped in a pagination envelope.

| parameter | type | description | example | required |
|----|----|----|----|----|
| data | array | array of folder stats objects (see above) | [...] | true |
| page | int | current page number | 1 | true |
| page_size | int | number of results per page | 20 | true |
| total_count | int | total number of matching records | 10 | true |
| total_pages | int | total number of pages | 1 | true |

##### Example

```json
{
    "data": [
        {
            "folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
            "folder_name": "High Risk Customers",
            "region": "MALTA",
            "user_count": 42,
            "oldest_user_date": "2026-01-15T10:00:00Z"
        }
    ],
    "page": 1,
    "page_size": 20,
    "total_count": 10,
    "total_pages": 1
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
