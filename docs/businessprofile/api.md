# ECDD Business Profile API

The ECDD Business Profile API provides endpoints for managing ECDD business profiles. It supports listing business profiles with filtering, sorting, and pagination, as well as update, patch, and delete operations.

- [Get All Business Profiles](#get-all-business-profiles)
- [Get Business Profile By PK](#get-business-profile-by-pk)
- [Create Business Profile](#create-business-profile)
- [Update Business Profile](#update-business-profile)
- [Patch Business Profile](#patch-business-profile)
- [Delete Business Profile](#delete-business-profile)

---

## Get All Business Profiles

### Request

GET

> http://`<ecddapi>`/api/ecdd/businessprofile

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### URL Params

| query string parameter | description | type | example | required |
|----|----|----|----|----|
| enabled | filter by enabled status | bool | true | false |
| country_id | country ID to filter by | int | 1 | false |
| risk_status_id | risk status ID to filter by | int | 2 | false |
| sort_by | field to sort results by | string | country_id | false |
| sort_dir | sort direction | string | asc | false |
| page | page number for pagination (starts at 1) | int | 1 | false |
| page_size | number of results per page (max 100, default 20) | int | 20 | false |

**sort_by valid values:** ecdd_business_profile_pk, country_id, risk_status_id, average_deposit, deposit_multiplier, logged_at

**sort_dir valid values:** asc, desc (default: asc)

### Response 200

> HTTP status 200 OK

#### Body (without pagination)

content-type: application/json

Returns an array of business profile objects.

| parameter | type | description | example | required |
|----|----|----|----|----|
| ecdd_business_profile_pk | string | spanner generated UUID primary key | "b1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c" | true |
| country_id | int | country identifier reference | 1 | true |
| state_id | int | state identifier reference (nullable) | null | false |
| risk_status_id | int | risk status identifier reference | 2 | true |
| average_deposit | float | average deposit value | 5000.00 | true |
| deposit_multiplier | float | multiplier applied to the deposit threshold | 2.00 | true |
| time_period_days | int | time period in days for the profile | 28 | true |
| enabled | bool | whether the business profile is active | true | true |
| logged_at | string | timestamp when the row was last logged/modified (RFC3339) | "2026-01-28T14:25:00Z" | true |
| updated_by | string | username of the person who last updated this record | "admin@company.com" | true |

##### Example

```json
[
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
]
```

#### Body (with pagination)

content-type: application/json

When `page` or `page_size` query parameters are provided, the response is wrapped in a pagination envelope.

| parameter | type | description | example | required |
|----|----|----|----|----|
| data | array | array of business profile objects (see above) | [...] | true |
| page | int | current page number | 1 | true |
| page_size | int | number of results per page | 20 | true |
| total_count | int | total number of matching records | 50 | true |
| total_pages | int | total number of pages | 3 | true |

##### Example

```json
{
    "data": [
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

## Get Business Profile By PK

### Request

GET

> http://`<ecddapi>`/api/ecdd/businessprofile/{businessprofilepk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| businessprofilepk | the ecdd_business_profile_pk of the business profile record | string | b1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c | true |

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns a single business profile object (see [Get All Business Profiles](#get-all-business-profiles) response body for field definitions).

##### Example

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

### Response 404

> HTTP status 404 Not Found

#### Body

content-type: application/json

##### Example

```json
{"error": "Business profile not found"}
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

## Create Business Profile

### Request

POST

> http://`<ecddapi>`/api/ecdd/businessprofile

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Body

content-type: application/json

New business profile object. `ecdd_business_profile_pk` and `logged_at` are server-generated and must not be provided.

| parameter | type | description | example | required |
|----|----|----|----|----|
| country_id | int | country identifier reference | 1 | true |
| state_id | int | state identifier reference (nullable) | null | false |
| risk_status_id | int | risk status identifier reference | 2 | true |
| average_deposit | float | average deposit value | 5000.00 | true |
| deposit_multiplier | float | multiplier applied to the deposit threshold | 2.00 | true |
| time_period_days | int | time period in days for the profile | 28 | true |
| enabled | bool | whether the business profile is active | true | true |
| updated_by | string | username of the person creating this record | "admin@company.com" | true |

##### Example

```json
{
    "country_id": 1,
    "state_id": null,
    "risk_status_id": 2,
    "average_deposit": 5000.00,
    "deposit_multiplier": 2.00,
    "time_period_days": 28,
    "enabled": true,
    "updated_by": "admin@company.com"
}
```

### Response 201

> HTTP status 201 Created

#### Body

content-type: application/json

Returns the created business profile object, including the server-generated `ecdd_business_profile_pk` and `logged_at`.

##### Example

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

## Update Business Profile

### Request

PUT

> http://`<ecddapi>`/api/ecdd/businessprofile/{businessprofilepk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| businessprofilepk | the ecdd_business_profile_pk of the business profile record | string | b1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c | true |

#### Body

content-type: application/json

Full business profile object. All fields are replaced.

| parameter | type | description | example | required |
|----|----|----|----|----|
| country_id | int | country identifier reference | 1 | true |
| state_id | int | state identifier reference (nullable) | null | false |
| risk_status_id | int | risk status identifier reference | 2 | true |
| average_deposit | float | average deposit value | 5000.00 | true |
| deposit_multiplier | float | multiplier applied to the deposit threshold | 2.00 | true |
| time_period_days | int | time period in days for the profile | 28 | true |
| enabled | bool | whether the business profile is active | true | true |
| updated_by | string | username of the person updating this record | "admin@company.com" | true |

##### Example

```json
{
    "country_id": 1,
    "state_id": null,
    "risk_status_id": 3,
    "average_deposit": 7500.00,
    "deposit_multiplier": 2.50,
    "time_period_days": 30,
    "enabled": true,
    "updated_by": "admin@company.com"
}
```

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns the updated business profile object.

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
{"error": "Business profile not found"}
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

## Patch Business Profile

### Request

PATCH

> http://`<ecddapi>`/api/ecdd/businessprofile/{businessprofilepk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| businessprofilepk | the ecdd_business_profile_pk of the business profile record | string | b1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c | true |

#### Body

content-type: application/json

Partial business profile object. Only provided fields are updated; omitted fields retain their existing values. See [Update Business Profile](#update-business-profile) request body for field definitions.

##### Example

```json
{
    "average_deposit": 7500.00,
    "enabled": false
}
```

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns the full updated business profile object.

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
{"error": "Business profile not found"}
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

## Delete Business Profile

### Request

DELETE

> http://`<ecddapi>`/api/ecdd/businessprofile/{businessprofilepk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| businessprofilepk | the ecdd_business_profile_pk of the business profile record | string | b1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c | true |

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

| parameter | type | description | example | required |
|----|----|----|----|----|
| message | string | confirmation message | "Business profile soft-deleted successfully" | true |

##### Example

```json
{
    "message": "Business profile soft-deleted successfully"
}
```

### Response 404

> HTTP status 404 Not Found

#### Body

content-type: application/json

##### Example

```json
{"error": "Business profile not found"}
```

### Response 405

> HTTP status 405 Method Not Allowed

#### Body

content-type: application/json

##### Example

```json
{"error": "Method not allowed"}
```
