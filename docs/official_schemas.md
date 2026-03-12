# Official Spanner Schemas (from Jira)

Source Jira tickets under project PRJFDD. All tables live in `pib-user-db`.

---

## 1. ECDDUserStatus (PRJFDD-32)

| Column | Type | Index | Nullable | Description |
|---|---|---|---|---|
| ECDDUserStatus_PK | STRING | Primary | No | Spanner generated UUID |
| user_id | INT64 | Unique (composite) | No | User identifier from core system |
| user_name | STRING(255) | Unique (composite) | No | Display name |
| country_id | INT64 | Unique (composite) | No | Country identifier |
| state_id | INT64 | Unique (composite) | Yes | State identifier |
| ecdd_status | INT64 | | No | 1=Not Required, 2=In progress, 3=Complete, 4=Suspended-Manual, 5=Suspended-Auto, 6=Closed, 7=Block Process |
| ecdd_threshold | NUMERIC | | No | Threshold value |
| ecdd_review_trigger | INT64 | | No | Review trigger type |
| ecdd_suspension_due_date | DATE | | Yes | Suspension due date |
| ecdd_multiplier | NUMERIC | | No | Multiplier value |
| ecdd_multiplier_rg_flag | BOOL | | No | RG flag |
| user_lt_net_deposit_threshold_gbp | NUMERIC | | No | Lifetime net deposit threshold GBP |
| user_lt_deposit_threshold_gbp | NUMERIC | | No | Lifetime deposit threshold GBP |
| user_12month_net_deposit_threshold_gbp | NUMERIC | | No | 12-month net deposit threshold GBP |
| info_source | INT64 | | No | Information source identifier |
| sign_off_status | INT64 | | No | Sign-off status |
| date_last_ecdd_sign_off | DATE | | Yes | Last ECDD sign-off date |
| ecdd_rg_review_status | INT64 | | No | RG review status |
| date_last_ecdd_rg_sign_off | DATE | | Yes | Last RG sign-off date |
| ecdd_report_status | INT64 | | No | Report status |
| ecdd_review_status | INT64 | | No | Review status |
| ecdd_document_status | INT64 | | No | Document status |
| ecdd_escalation_status | INT64 | | No | Escalation status |
| uar_status | INT64 | | No | UAR status |
| logged_at | TIMESTAMP | | No | Last modified timestamp |
| updated_by | STRING(255) | | No | Last updater username |

### Delta vs Current Mock Model (`models/user_status.go`)
- **Mock has `language` (STRING)** — NOT in official schema. Mock-only field for UI convenience.
- Otherwise matches.

---

## 2. ECDDThresholdConfig (PRJFDD-34)

| Column | Type | Index | Nullable | Description |
|---|---|---|---|---|
| ECDDThresholdConfig_PK | STRING | Primary | No | Spanner generated UUID |
| title | STRING(255) | Unique | No | Descriptive name |
| is_active | BOOL | | No | Activation status |
| country_id | INT64 | | No | Country identifier |
| state_id | INT64 | | Yes | State identifier |
| type | INT64 | | No | 1=Deposit, 2=Net Deposit, 3=Stakes |
| **reinvest** | **BOOL** | | **No** | **Whether to check reinvestment on the account** |
| value | NUMERIC | | No | Monetary threshold amount |
| currency_id | INT64 | | No | Currency identifier |
| period | INT64 | | No | 1=24hrs, 2=28days, 3=84days, 4=91days, 5=182days, 6=365days |
| use_multipliers | BOOL | | No | Apply ECDD multiplier adjustments |
| use_rg_flag | BOOL | | No | Consider ECDD Multiplier RG Flag |
| apply_all_statuses | BOOL | | No | Apply to all statuses or only "Not Required" |
| backfill | BOOL | | No | Apply to existing accounts |
| hierarchy | INT64 | Unique | No | Priority order (lower = higher priority) |
| ecdd_status | INT64 | | No | ECDD Status on breach |
| ecdd_review_status | INT64 | | No | Review Status on breach |
| ecdd_report_status | INT64 | | No | Report Status on breach |
| sign_off_status | INT64 | | No | Sign Off Status on breach |
| customer_risk_level | INT64 | | No | 1=Low, 2=Medium, 3=Medium-High, 4=High |
| ndl_28_day_gbp | NUMERIC | | No | 28 Day NDL in GBP |
| ndl_monthly_gbp | NUMERIC | | No | Monthly NDL in GBP |
| case_management_folder_pk | STRING | | Yes | FK to ECDDCaseManagementFolder |
| logged_at | TIMESTAMP | | No | Last modified timestamp |
| updated_by | STRING(255) | | No | Last updater username |

### Delta vs Current Mock Model (`models/threshold_config.go`)
- **MISSING `Reinvest bool`** — must add between `Type` and `Value`
- Otherwise matches.

---

## 3. ECDDCaseManagementFolder (PRJFDD-36)

| Column | Type | Index | Nullable | Description |
|---|---|---|---|---|
| ECDDCaseManagementFolder_PK | STRING | Primary | No | Spanner generated UUID |
| folder_name | STRING(255) | Unique | No | Folder name |
| **region** | **STRING** | | **No** | **Region name (MALTA, GIBRALTAR, USA, AUSTRALIA). UK falls under MALTA as a licensed country (country_id=197).** |
| **country_id** | **INT64** | | **Yes** | **Country identifier** |
| **state_id** | **INT64** | | **Yes** | **State identifier** |
| logged_at | TIMESTAMP | | No | Last modified timestamp |
| updated_by | STRING(255) | | No | Last updater username |

### Delta vs Current Mock Model (`models/case_management_folder.go`)
- **MISSING `Region string`** — region scoping for the folder
- **MISSING `CountryID *int64`** — optional country-level scoping
- **MISSING `StateID *int64`** — optional state-level scoping

---

## 4. ECDDMultiplierConfig (PRJFDD-38)

| Column | Type | Index | Nullable | Description |
|---|---|---|---|---|
| ECDDMultiplierConfig_PK | STRING | Primary | No | Spanner generated UUID |
| country_id | INT64 | Unique (composite) | No | Country identifier |
| state_id | INT64 | Unique (composite) | Yes | State identifier |
| age_multipliers | ARRAY<INT64> | | No | Ages where 0.5 multiplier applies |
| status_multiplier | BOOL | | No | Status-based 0.5 multiplier flag |
| is_active | BOOL | | No | Activation status |
| logged_at | TIMESTAMP | | No | Last modified timestamp |
| updated_by | STRING(255) | | No | Last updater username |

### Delta vs Current Mock Model (`models/multiplier_config.go`)
- No schema changes. Model matches.

---

## 5. ECDDBusinessProfile (PRJFDD-40)

| Column | Type | Index | Nullable | Description |
|---|---|---|---|---|
| ECDDBusinessProfile_PK | STRING | Primary | No | Spanner generated UUID |
| country_id | INT64 | Unique (composite) | No | Country identifier |
| state_id | INT64 | Unique (composite) | Yes | State identifier |
| risk_status_id | INT64 | | No | Risk status identifier |
| average_deposit | NUMERIC | | No | Average deposit baseline |
| deposit_multiplier | NUMERIC | | No | Multiplier for breach threshold |
| time_period_days | INT64 | | No | Evaluation period in days |
| enabled | BOOL | | No | Whether rule is active |
| logged_at | TIMESTAMP | | No | Last modified timestamp |
| updated_by | STRING(255) | | No | Last updater username |

### Delta vs Current Mock Model (`models/business_profile.go`)
- No schema changes. Model matches.

---

## 6. ECDDUAR (PRJFDD-42) — NEW TABLE

| Column | Type | Index | Nullable | Description |
|---|---|---|---|---|
| ECDDUAR_PK | STRING | Primary | No | Spanner generated UUID |
| user_status_pk | STRING | | No | FK to ECDDUserStatus_PK |
| user_id | INT64 | | No | User identifier |
| username | STRING(255) | | No | Customer account username |
| country_id | INT64 | | No | Country identifier |
| state_id | INT64 | | Yes | State identifier |
| identified_by | STRING(255) | | No | Back-office user who identified activity |
| job_title_department | STRING(500) | | No | Job title/dept of identifier |
| date_activity_identified | DATE | | No | Date activity was identified |
| activity_identified | INT64 | | No | Activity type from dropdown |
| report_reason | STRING(MAX) | | No | Free text reason for UAR |
| attachment_urls | ARRAY<STRING(MAX)> | | Yes | File attachment URLs |
| date_reviewed | DATE | | Yes | Date UAR was reviewed |
| reviewed_by | STRING(255) | | Yes | Reviewer name |
| reviewer_job_title_department | STRING(500) | | Yes | Reviewer job title/dept |
| secondary_review | STRING(MAX) | | Yes | Secondary review notes |
| action_taken | INT64 | | Yes | 1=ECDD Required, 2=Customer Contact Required, 3=Escalated, 4=No Action Required |
| submitted_at | TIMESTAMP | | No | UAR submission timestamp |
| reviewed_at | TIMESTAMP | | Yes | Review completion timestamp |
| logged_at | TIMESTAMP | | No | Last modified timestamp |
| updated_by | STRING(255) | | No | Last updater username |

### Delta vs Current Mock
- **Entirely new table** — no model, handler, data, or API exists yet.
- FK: `user_status_pk` → `ECDDUserStatus.ECDDUserStatus_PK`

---

## 7. ECDDUserCaseManagementFolder (PRJFDD-44)

| Column | Type | Index | Nullable | Description |
|---|---|---|---|---|
| ECDDUserCaseManagementFolder_PK | STRING | Primary | No | Spanner generated UUID |
| case_management_folder_pk | STRING | | No | FK to ECDDCaseManagementFolder_PK |
| user_status_pk | STRING | | No | FK to ECDDUserStatus_PK |
| logged_at | TIMESTAMP | | No | Last modified timestamp |
| updated_by | STRING(255) | | No | Last updater username |

### Delta vs Current Mock Model (`models/user_case_folder.go`)
- Mock JSON tag now matches official: `case_management_folder_pk`.
- Otherwise matches.

---

## Summary of Required Changes

| Table | Status | Changes Needed |
|---|---|---|
| ECDDUserStatus | Exists | `language` is mock-only (keep for UI) |
| ECDDThresholdConfig | Exists | Add `reinvest bool` field |
| ECDDCaseManagementFolder | Exists | Add `region`, `country_id`, `state_id` fields |
| ECDDMultiplierConfig | Exists | No changes |
| ECDDBusinessProfile | Exists | No changes |
| ECDDUAR | **NEW** | Create model, handler, data, routes |
| ECDDUserCaseManagementFolder | Exists | `folder_pk` renamed to `case_management_folder_pk` (done) |

## Entity Relationships

```
ECDDThresholdConfig.case_management_folder_pk  →  ECDDCaseManagementFolder.PK
ECDDUserCaseManagementFolder.case_management_folder_pk  →  ECDDCaseManagementFolder.PK
ECDDUserCaseManagementFolder.user_status_pk  →  ECDDUserStatus.PK
ECDDUAR.user_status_pk  →  ECDDUserStatus.PK
```

All config tables (ThresholdConfig, MultiplierConfig, BusinessProfile) share `country_id` + `state_id` as jurisdiction keys. CaseManagementFolder now also has `region` + `country_id` + `state_id` for jurisdiction scoping.
