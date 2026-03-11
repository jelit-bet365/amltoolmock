# ECDD Threshold Config API

The ECDD Threshold Config API provides endpoints for managing ECDD threshold configuration records. It supports listing threshold configs with filtering, sorting, and pagination, as well as update, patch, and delete operations.

- [Get All Threshold Configs](#get-all-threshold-configs)
- [Get Threshold Config By PK](#get-threshold-config-by-pk)
- [Create Threshold Config](#create-threshold-config)
- [Update Threshold Config](#update-threshold-config)
- [Patch Threshold Config](#patch-threshold-config)
- [Delete Threshold Config](#delete-threshold-config)

---

## Get All Threshold Configs

### Request

GET

> http://`<ecddapi>`/api/ecdd/thresholdconfig

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
| type | threshold type to filter by | int | 2 | false |
| sort_by | field to sort results by | string | hierarchy | false |
| sort_dir | sort direction | string | asc | false |
| page | page number for pagination (starts at 1) | int | 1 | false |
| page_size | number of results per page (max 100, default 20) | int | 20 | false |

**sort_by valid values:** ecdd_threshold_config_pk, title, country_id, type, value, hierarchy, logged_at

**sort_dir valid values:** asc, desc (default: asc)

### Response 200

> HTTP status 200 OK

#### Body (without pagination)

content-type: application/json

Returns an array of threshold config objects.

| parameter | type | description | example | required |
|----|----|----|----|----|
| ecdd_threshold_config_pk | string | spanner generated UUID primary key | "t1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c" | true |
| title | string | descriptive name for the threshold rule | "UK 28 day NDL 50K" | true |
| is_active | bool | activation status of the threshold rule | true | true |
| country_id | int | country identifier where threshold applies | 1 | true |
| state_id | int | state identifier where threshold applies (nullable for country-wide) | null | false |
| type | int | threshold type (1=Deposit, 2=Net Deposit, 3=Stakes) | 2 | true |
| reinvest | bool | whether to look at if there is any reinvestment on the account | true | true |
| value | float | monetary threshold amount that triggers the rule | 50000.00 | true |
| currency_id | int | currency identifier for the threshold value | 1 | true |
| period | int | period (1=24hrs, 2=28days, 3=84days, 4=91days, 5=182days, 6=365days) | 2 | true |
| use_multipliers | bool | whether to apply ECDD multiplier adjustments | true | true |
| use_rg_flag | bool | whether to consider ECDD Multiplier RG Flag | true | true |
| apply_all_statuses | bool | apply to all ECDD statuses (true) or only "Not Required" (false) | false | true |
| backfill | bool | apply to existing accounts (true) or new accounts only (false) | false | true |
| hierarchy | int | priority order for threshold evaluation (lower number = higher priority) | 10 | true |
| ecdd_status | int | ECDD Status to set when threshold is breached | 2 | true |
| ecdd_review_status | int | ECDD Review Status to set when threshold is breached | 4 | true |
| ecdd_report_status | int | ECDD Report Status to set when threshold is breached | 2 | true |
| sign_off_status | int | Sign Off Status to set when threshold is breached | 1 | true |
| customer_risk_level | int | customer risk level (1=Low, 2=Medium, 3=Medium-High, 4=High) | 3 | true |
| ndl_28_day_gbp | float | 28 Day Net Deposit Limit in GBP to be applied | 5000.00 | true |
| ndl_monthly_gbp | float | Monthly Net Deposit Limit in GBP to be applied | 20000.00 | true |
| case_management_folder_pk | string | foreign key on Case Management folder table for user assignment (nullable) | "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d" | false |
| logged_at | string | timestamp when the row was last logged/modified (RFC3339) | "2026-01-28T14:25:00Z" | true |
| updated_by | string | username of the person who last updated this record | "john.doe@company.com" | true |

##### Example

```json
[
    {
        "ecdd_threshold_config_pk": "t1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c",
        "title": "UK 28 day NDL 50K",
        "is_active": true,
        "country_id": 1,
        "state_id": null,
        "type": 2,
        "reinvest": true,
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

#### Body (with pagination)

content-type: application/json

When `page` or `page_size` query parameters are provided, the response is wrapped in a pagination envelope.

| parameter | type | description | example | required |
|----|----|----|----|----|
| data | array | array of threshold config objects (see above) | [...] | true |
| page | int | current page number | 1 | true |
| page_size | int | number of results per page | 20 | true |
| total_count | int | total number of matching records | 50 | true |
| total_pages | int | total number of pages | 3 | true |

##### Example

```json
{
    "data": [
        {
            "ecdd_threshold_config_pk": "t1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c",
            "title": "UK 28 day NDL 50K",
            "is_active": true,
            "country_id": 1,
            "state_id": null,
            "type": 2,
            "reinvest": true,
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

## Get Threshold Config By PK

### Request

GET

> http://`<ecddapi>`/api/ecdd/thresholdconfig/{thresholdconfigpk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| thresholdconfigpk | the ecdd_threshold_config_pk of the threshold config record | string | t1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c | true |

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns a single threshold config object (see [Get All Threshold Configs](#get-all-threshold-configs) response body for field definitions).

##### Example

```json
{
    "ecdd_threshold_config_pk": "t1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c",
    "title": "UK 28 day NDL 50K",
    "is_active": true,
    "country_id": 1,
    "state_id": null,
    "type": 2,
    "reinvest": true,
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
```

### Response 404

> HTTP status 404 Not Found

#### Body

content-type: application/json

##### Example

```json
{"error": "Threshold not found"}
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

## Create Threshold Config

### Request

POST

> http://`<ecddapi>`/api/ecdd/thresholdconfig

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Body

content-type: application/json

New threshold config object. The `ecdd_threshold_config_pk` and `logged_at` fields are server-generated and must not be included in the request.

| parameter | type | description | example | required |
|----|----|----|----|----|
| title | string | descriptive name for the threshold rule | "UK 28 day NDL 50K" | true |
| is_active | bool | activation status of the threshold rule | true | true |
| country_id | int | country identifier where threshold applies | 1 | true |
| state_id | int | state identifier where threshold applies (nullable for country-wide) | null | false |
| type | int | threshold type (1=Deposit, 2=Net Deposit, 3=Stakes) | 2 | true |
| reinvest | bool | whether to look at if there is any reinvestment on the account | true | true |
| value | float | monetary threshold amount that triggers the rule | 50000.00 | true |
| currency_id | int | currency identifier for the threshold value | 1 | true |
| period | int | period (1=24hrs, 2=28days, 3=84days, 4=91days, 5=182days, 6=365days) | 2 | true |
| use_multipliers | bool | whether to apply ECDD multiplier adjustments | true | true |
| use_rg_flag | bool | whether to consider ECDD Multiplier RG Flag | true | true |
| apply_all_statuses | bool | apply to all ECDD statuses (true) or only "Not Required" (false) | false | true |
| backfill | bool | apply to existing accounts (true) or new accounts only (false) | false | true |
| hierarchy | int | priority order for threshold evaluation (lower number = higher priority) | 10 | true |
| ecdd_status | int | ECDD Status to set when threshold is breached | 2 | true |
| ecdd_review_status | int | ECDD Review Status to set when threshold is breached | 4 | true |
| ecdd_report_status | int | ECDD Report Status to set when threshold is breached | 2 | true |
| sign_off_status | int | Sign Off Status to set when threshold is breached | 1 | true |
| customer_risk_level | int | customer risk level (1=Low, 2=Medium, 3=Medium-High, 4=High) | 3 | true |
| ndl_28_day_gbp | float | 28 Day Net Deposit Limit in GBP to be applied | 5000.00 | true |
| ndl_monthly_gbp | float | Monthly Net Deposit Limit in GBP to be applied | 20000.00 | true |
| case_management_folder_pk | string | foreign key on Case Management folder table for user assignment (nullable) | "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d" | false |
| updated_by | string | username of the person creating this record | "john.doe@company.com" | true |

##### Example

```json
{
    "title": "UK 28 day NDL 50K",
    "is_active": true,
    "country_id": 1,
    "state_id": null,
    "type": 2,
    "reinvest": true,
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
    "updated_by": "john.doe@company.com"
}
```

### Response 201

> HTTP status 201 Created

#### Body

content-type: application/json

Returns the created threshold config object, including the server-generated `ecdd_threshold_config_pk` and `logged_at`.

##### Example

```json
{
    "ecdd_threshold_config_pk": "t1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c",
    "title": "UK 28 day NDL 50K",
    "is_active": true,
    "country_id": 1,
    "state_id": null,
    "type": 2,
    "reinvest": true,
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

## Update Threshold Config

### Request

PUT

> http://`<ecddapi>`/api/ecdd/thresholdconfig/{thresholdconfigpk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| thresholdconfigpk | the ecdd_threshold_config_pk of the threshold config record | string | t1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c | true |

#### Body

content-type: application/json

Full threshold config object. All fields are replaced.

| parameter | type | description | example | required |
|----|----|----|----|----|
| title | string | descriptive name for the threshold rule | "UK 28 day NDL 50K" | true |
| is_active | bool | activation status of the threshold rule | true | true |
| country_id | int | country identifier where threshold applies | 1 | true |
| state_id | int | state identifier where threshold applies (nullable for country-wide) | null | false |
| type | int | threshold type (1=Deposit, 2=Net Deposit, 3=Stakes) | 2 | true |
| reinvest | bool | whether to look at if there is any reinvestment on the account | true | true |
| value | float | monetary threshold amount that triggers the rule | 50000.00 | true |
| currency_id | int | currency identifier for the threshold value | 1 | true |
| period | int | period (1=24hrs, 2=28days, 3=84days, 4=91days, 5=182days, 6=365days) | 2 | true |
| use_multipliers | bool | whether to apply ECDD multiplier adjustments | true | true |
| use_rg_flag | bool | whether to consider ECDD Multiplier RG Flag | true | true |
| apply_all_statuses | bool | apply to all ECDD statuses (true) or only "Not Required" (false) | false | true |
| backfill | bool | apply to existing accounts (true) or new accounts only (false) | false | true |
| hierarchy | int | priority order for threshold evaluation (lower number = higher priority) | 10 | true |
| ecdd_status | int | ECDD Status to set when threshold is breached | 2 | true |
| ecdd_review_status | int | ECDD Review Status to set when threshold is breached | 4 | true |
| ecdd_report_status | int | ECDD Report Status to set when threshold is breached | 2 | true |
| sign_off_status | int | Sign Off Status to set when threshold is breached | 1 | true |
| customer_risk_level | int | customer risk level (1=Low, 2=Medium, 3=Medium-High, 4=High) | 3 | true |
| ndl_28_day_gbp | float | 28 Day Net Deposit Limit in GBP to be applied | 5000.00 | true |
| ndl_monthly_gbp | float | Monthly Net Deposit Limit in GBP to be applied | 20000.00 | true |
| case_management_folder_pk | string | foreign key on Case Management folder table for user assignment (nullable) | "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d" | false |
| updated_by | string | username of the person updating this record | "john.doe@company.com" | true |

##### Example

```json
{
    "title": "UK 28 day NDL 50K",
    "is_active": true,
    "country_id": 1,
    "state_id": null,
    "type": 2,
    "reinvest": true,
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
    "updated_by": "john.doe@company.com"
}
```

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns the updated threshold config object.

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
{"error": "Threshold not found"}
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

## Patch Threshold Config

### Request

PATCH

> http://`<ecddapi>`/api/ecdd/thresholdconfig/{thresholdconfigpk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| thresholdconfigpk | the ecdd_threshold_config_pk of the threshold config record | string | t1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c | true |

#### Body

content-type: application/json

Partial threshold config object. Only provided fields are updated; omitted fields retain their existing values. See [Update Threshold Config](#update-threshold-config) request body for field definitions.

##### Example

```json
{
    "is_active": false,
    "value": 75000.00
}
```

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

Returns the full updated threshold config object.

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
{"error": "Threshold not found"}
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

## Delete Threshold Config

### Request

DELETE

> http://`<ecddapi>`/api/ecdd/thresholdconfig/{thresholdconfigpk}

#### Headers

| header | description | possibleValues | required |
|----|----|----|-----|
| bet365-applicationname | the name of the application sending the request | Mobile | false |
| bet365-correlationid | the unique id for the request | a2e70fe9-616d-471f-b6df-5807a166bea0 | false |
| bet365-username | the bet365 username associated with the request | williamarmitage | false |

#### Path Params

| parameter | description | type | example | required |
|----|----|----|----|----|
| thresholdconfigpk | the ecdd_threshold_config_pk of the threshold config record | string | t1a2b3c4-d5e6-4f7a-8b9c-0d1e2f3a4b5c | true |

### Response 200

> HTTP status 200 OK

#### Body

content-type: application/json

| parameter | type | description | example | required |
|----|----|----|----|----|
| message | string | confirmation message | "Threshold soft-deleted successfully" | true |

##### Example

```json
{
    "message": "Threshold soft-deleted successfully"
}
```

### Response 404

> HTTP status 404 Not Found

#### Body

content-type: application/json

##### Example

```json
{"error": "Threshold not found"}
```

### Response 405

> HTTP status 405 Method Not Allowed

#### Body

content-type: application/json

##### Example

```json
{"error": "Method not allowed"}
```
