# ECDD Case Management Folder API

The ECDD Case Management Folder API provides endpoints for managing ECDD case management folders. It supports listing folders with filtering, sorting, and pagination, as well as create, update, and delete operations for folders stored in the ECDDCaseManagementFolder Spanner table.

- [Get All Folders](#get-all-folders)
- [Get Folder By PK](#get-folder-by-pk)
- [Create Case Management Folder](#create-case-management-folder)
- [Update Folder](#update-folder)
- [Delete Folder](#delete-folder)

---

## Get All Folders

### Request

GET

> http://`<ecddapi>`/api/ecdd/casemanagementfolder

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### URL Params

| query string parameter | description | type | example | required |
|----|----|----|----|----|
| search | filter by folder name (case-insensitive substring match) | string | High Risk | false |
| sort_by | field to sort results by | string | folder_name | false |
| sort_dir | sort direction | string | asc | false |
| page | page number for pagination (starts at 1) | int | 1 | false |
| page_size | number of results per page (max 100, default 20) | int | 20 | false |

**sort_by valid values:** folder_name, ecdd_case_management_folder_pk, logged_at

**sort_dir valid values:** asc, desc (default: asc)

### Response 200

> HTTP status 200 OK

#### Body (without pagination)

content-type: application/json

Returns an array of case management folder objects.

| parameter | type | description | example | required |
|----|----|----|----|----|
| ecdd_case_management_folder_pk | string | spanner generated UUID primary key | "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d" | true |
| folder_name | string | name of the folder | "High Risk Customers" | true |
| logged_at | string | timestamp when the row was last logged/modified (RFC3339) | "2026-01-15T10:00:00Z" | true |
| updated_by | string | username of the person who last updated this record | "admin@company.com" | true |

##### Example

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

#### Body (with pagination)

content-type: application/json

When `page` or `page_size` query parameters are provided, the response is wrapped in a pagination envelope.

| parameter | type | description | example | required |
|----|----|----|----|----|
| data | array | array of case management folder objects (see above) | [...] | true |
| page | int | current page number | 1 | true |
| page_size | int | number of results per page | 20 | true |
| total_count | int | total number of matching records | 10 | true |
| total_pages | int | total number of pages | 1 | true |

##### Example

```json
{
    "data": [
        {
            "ecdd_case_management_folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
            "folder_name": "High Risk Customers",
            "logged_at": "2026-01-15T10:00:00Z",
            "updated_by": "admin@company.com"
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

---

## Get Folder By PK

### Request

GET

> http://`<ecddapi>`/api/ecdd/casemanagementfolder/{folderpk}

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

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns a single case management folder object (see [Get All Folders](#get-all-folders) response body for field definitions).

##### Example

```json
{
    "ecdd_case_management_folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
    "folder_name": "High Risk Customers",
    "logged_at": "2026-01-15T10:00:00Z",
    "updated_by": "admin@company.com"
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

## Create Case Management Folder

### Request

POST

> http://`<ecddapi>`/api/ecdd/casemanagementfolder

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
| folder_name | string | name of the folder | "High Risk Customers" | true |
| updated_by | string | username of the person creating this record | "admin@company.com" | true |

##### Example

```json
{
    "folder_name": "High Risk Customers",
    "updated_by": "admin@company.com"
}
```

### Response 201

> HTTP status 201 Created

#### Body

content-type: application/json

Returns the created case management folder object with the server-generated primary key and logged_at timestamp.

| parameter | type | description | example | required |
|----|----|----|----|----|
| ecdd_case_management_folder_pk | string | spanner generated UUID primary key | "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d" | true |
| folder_name | string | name of the folder | "High Risk Customers" | true |
| logged_at | string | timestamp when the row was created (RFC3339) | "2026-01-15T10:00:00Z" | true |
| updated_by | string | username of the person who created this record | "admin@company.com" | true |

##### Example

```json
{
    "ecdd_case_management_folder_pk": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
    "folder_name": "High Risk Customers",
    "logged_at": "2026-01-15T10:00:00Z",
    "updated_by": "admin@company.com"
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

## Update Folder

### Request

PUT

> http://`<ecddapi>`/api/ecdd/casemanagementfolder/{folderpk}

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

Full folder object. All fields are replaced.

| parameter | type | description | example | required |
|----|----|----|----|----|
| folder_name | string | name of the folder | "High Risk Customers" | true |
| updated_by | string | username of the person updating this record | "admin@company.com" | true |

##### Example

```json
{
    "folder_name": "High Risk Customers",
    "updated_by": "admin@company.com"
}
```

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns the updated case management folder object.

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

## Delete Folder

### Request

DELETE

> http://`<ecddapi>`/api/ecdd/casemanagementfolder/{folderpk}

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

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

| parameter | type | description | example | required |
|----|----|----|----|----|
| message | string | confirmation message including cascade-deleted user-folder assignment count | "Folder deleted successfully. 3 user-folder assignment(s) cascade-deleted." | true |

##### Example

```json
{
    "message": "Folder deleted successfully. 3 user-folder assignment(s) cascade-deleted."
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
