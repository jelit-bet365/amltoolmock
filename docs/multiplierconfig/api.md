# ECDD Multiplier Config API

The ECDD Multiplier Config API provides endpoints for managing ECDD multiplier configuration records. It supports listing multiplier configs with filtering, sorting, and pagination, as well as update, patch, and delete operations.

- [Get All Multiplier Configs](#get-all-multiplier-configs)
- [Get Multiplier Config By PK](#get-multiplier-config-by-pk)
- [Create Multiplier Config](#create-multiplier-config)
- [Update Multiplier Config](#update-multiplier-config)
- [Patch Multiplier Config](#patch-multiplier-config)
- [Delete Multiplier Config](#delete-multiplier-config)

---

## Get All Multiplier Configs

### Request

GET

> http://`<ecddapi>`/api/ecdd/multiplierconfig

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### URL Params

| query string parameter | description | type | example | required |
|----|----|----|----|----|
| is_active | filter by active status | bool | true | false |
| country_id | country ID to filter by | int | 1 | false |
| sort_by | field to sort results by | string | country_id | false |
| sort_dir | sort direction | string | asc | false |
| page | page number for pagination (starts at 1) | int | 1 | false |
| page_size | number of results per page (max 100, default 20) | int | 20 | false |

**sort_by valid values:** ecdd_multiplier_config_pk, country_id, logged_at

**sort_dir valid values:** asc, desc (default: asc)

### Response 200

> HTTP status 200 OK

#### Body (without pagination)

content-type: application/json

Returns an array of multiplier config objects.

| parameter | type | description | example | required |
|----|----|----|----|----|
| ecdd_multiplier_config_pk | string | spanner generated UUID primary key | "m1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c" | true |
| country_id | int | country identifier reference | 1 | true |
| state_id | int | state identifier reference (nullable) | 5 | false |
| age_multipliers | array | array of ages where 0.5 multiplier applies | [18, 19, 20, 21] | true |
| status_multiplier | bool | whether a status multiplier is applied | true | true |
| is_active | bool | whether this config is active | true | true |
| logged_at | string | timestamp when the row was last logged/modified (RFC3339) | "2026-01-28T14:25:00Z" | true |
| updated_by | string | username of the person who last updated this record | "admin@company.com" | true |

##### Example

```json
[
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
]
```

#### Body (with pagination)

content-type: application/json

When `page` or `page_size` query parameters are provided, the response is wrapped in a pagination envelope.

| parameter | type | description | example | required |
|----|----|----|----|----|
| data | array | array of multiplier config objects (see above) | [...] | true |
| page | int | current page number | 1 | true |
| page_size | int | number of results per page | 20 | true |
| total_count | int | total number of matching records | 50 | true |
| total_pages | int | total number of pages | 3 | true |

##### Example

```json
{
    "data": [
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

## Get Multiplier Config By PK

### Request

GET

> http://`<ecddapi>`/api/ecdd/multiplierconfig/{multiplierconfigpk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| multiplierconfigpk | the ecdd_multiplier_config_pk of the multiplier config record | string | m1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c | true |

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns a single multiplier config object (see [Get All Multiplier Configs](#get-all-multiplier-configs) response body for field definitions).

##### Example

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

### Response 404

> HTTP status 404 Not Found

#### Body

content-type: application/json

##### Example

```json
{"error": "Multiplier not found"}
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

## Create Multiplier Config

### Request

POST

> http://`<ecddapi>`/api/ecdd/multiplierconfig

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Body

content-type: application/json

New multiplier config object. The `ecdd_multiplier_config_pk` and `logged_at` fields are server-generated and must not be included in the request.

| parameter | type | description | example | required |
|----|----|----|----|----|
| country_id | int | country identifier reference | 1 | true |
| state_id | int | state identifier reference (nullable) | 5 | false |
| age_multipliers | array | array of ages where 0.5 multiplier applies | [18, 19, 20, 21] | true |
| status_multiplier | bool | whether a status multiplier is applied | true | true |
| is_active | bool | whether this config is active | true | true |
| updated_by | string | username of the person creating this record | "admin@company.com" | true |

##### Example

```json
{
    "country_id": 1,
    "state_id": null,
    "age_multipliers": [18, 19, 20, 21, 22, 23, 24, 25],
    "status_multiplier": true,
    "is_active": true,
    "updated_by": "admin@company.com"
}
```

### Response 201

> HTTP status 201 Created

#### Body

content-type: application/json

Returns the created multiplier config object, including the server-generated `ecdd_multiplier_config_pk` and `logged_at`.

##### Example

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

## Update Multiplier Config

### Request

PUT

> http://`<ecddapi>`/api/ecdd/multiplierconfig/{multiplierconfigpk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| multiplierconfigpk | the ecdd_multiplier_config_pk of the multiplier config record | string | m1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c | true |

#### Body

content-type: application/json

Full multiplier config object. All fields are replaced.

| parameter | type | description | example | required |
|----|----|----|----|----|
| country_id | int | country identifier reference | 1 | true |
| state_id | int | state identifier reference (nullable) | 5 | false |
| age_multipliers | array | array of ages where 0.5 multiplier applies | [18, 19, 20, 21] | true |
| status_multiplier | bool | whether a status multiplier is applied | true | true |
| is_active | bool | whether this config is active | true | true |
| updated_by | string | username of the person updating this record | "admin@company.com" | true |

##### Example

```json
{
    "country_id": 1,
    "state_id": null,
    "age_multipliers": [18, 19, 20, 21, 22, 23, 24, 25],
    "status_multiplier": true,
    "is_active": true,
    "updated_by": "admin@company.com"
}
```

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns the updated multiplier config object.

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
{"error": "Multiplier not found"}
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

## Patch Multiplier Config

### Request

PATCH

> http://`<ecddapi>`/api/ecdd/multiplierconfig/{multiplierconfigpk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| multiplierconfigpk | the ecdd_multiplier_config_pk of the multiplier config record | string | m1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c | true |

#### Body

content-type: application/json

Partial multiplier config object. Only provided fields are updated; omitted fields retain their existing values. See [Update Multiplier Config](#update-multiplier-config) request body for field definitions.

##### Example

```json
{
    "is_active": false
}
```

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns the full updated multiplier config object.

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
{"error": "Multiplier not found"}
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

## Delete Multiplier Config

### Request

DELETE

> http://`<ecddapi>`/api/ecdd/multiplierconfig/{multiplierconfigpk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| multiplierconfigpk | the ecdd_multiplier_config_pk of the multiplier config record | string | m1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c | true |

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

| parameter | type | description | example | required |
|----|----|----|----|----|
| message | string | confirmation message | "Multiplier soft-deleted successfully" | true |

##### Example

```json
{
    "message": "Multiplier soft-deleted successfully"
}
```

### Response 404

> HTTP status 404 Not Found

#### Body

content-type: application/json

##### Example

```json
{"error": "Multiplier not found"}
```

### Response 405

> HTTP status 405 Method Not Allowed

#### Body

content-type: application/json

##### Example

```json
{"error": "Method not allowed"}
```
