# Case Assignment Schema Proposal

## 1. Overview

This document proposes the database schema for case assignment -- the ability to assign ECDD user accounts (people being investigated) to analysts (people using the AML tool). This covers the `ecdd_analyst` table, the `ecdd_case_assignment` table, and a separate `ecdd_case_assignment_history` audit table.

### Key Terminology

| Term | Meaning |
|------|---------|
| **User** | The person being investigated -- an `ECDDUserStatus` record |
| **Analyst** | The person using the AML tool, authenticated via Okta |
| **Assignment** | A link between a user and an analyst for review purposes |

---

## 2. Research Findings (Existing Codebase Patterns)

### 2.1 Primary Key Convention

All existing tables use a string UUID primary key, named as `ecdd_{entity}_pk`. The UUID is generated server-side using `github.com/google/uuid` via `uuid.New().String()`.

**Examples:**
- `ecdd_user_status_pk`
- `ecdd_case_management_folder_pk`
- `ecdd_threshold_config_pk`
- `ecdd_multiplier_config_pk`
- `ecdd_business_profile_pk`
- `ecdd_user_case_management_folder_pk`

### 2.2 Common Fields

Every table includes:
- `logged_at` (`time.Time`, JSON: `logged_at`) -- server-set timestamp on create/update
- `updated_by` (`string`, JSON: `updated_by`) -- freeform string, currently stores email addresses (e.g., `admin@company.com`, `malta.compliance@company.com`)

### 2.3 Data Storage

The mock backend uses **in-memory maps** (`map[string]*Model`) with a singleton `DataService`, thread-safe via `sync.RWMutex`. Data is loaded from JSON files in `data/`. The production backend targets **Google Cloud Spanner**.

### 2.4 JSON Tag Conventions

- All JSON tags use `snake_case`
- Optional fields use `omitempty`
- Nullable fields use pointer types (`*int64`, `*time.Time`, `*string`)

### 2.5 Foreign Key Pattern

Foreign keys are stored as plain strings referencing the PK of another table. There is no formal FK enforcement in the mock -- validation is done at the application level. Example from `ECDDUserCaseManagementFolder`:
```go
FolderPK     string `json:"folder_pk"`       // FK to ECDDCaseManagementFolder
UserStatusPK string `json:"user_status_pk"`  // FK to ECDDUserStatus
```

### 2.6 Delete Patterns

- **Soft delete**: Used for config tables (`is_active = false` or `enabled = false`)
- **Hard delete**: Used for junction/assignment tables (removed from map) with cascade deletion of related records

### 2.7 API Path Convention

All endpoints follow `/api/v1/{resource}` and `/api/v1/{resource}/{id}`. Nested resources use `/api/v1/{parent}/{id}/{child}`.

### 2.8 Existing Relationship: Folders and Users

The `ecdd_user_case_management_folder` table is a many-to-many junction table linking users to folders. A user can be in multiple folders; a folder can contain multiple users. This is the closest analog to the case assignment pattern.

---

## 3. Schema Design

### 3.1 Entity-Relationship Description

```
ECDDAnalyst (1) ----< (M) ECDDCaseAssignment (M) >---- (1) ECDDUserStatus
                                |
                                | assigned_by_analyst_pk (FK to ECDDAnalyst)
                                |
ECDDCaseAssignmentHistory ------+  (audit trail, one row per state change)
```

**Cardinality rules:**
- A user (ECDDUserStatus) can have **at most one active assignment** at any time (1:1 active constraint)
- An analyst can have **many active assignments** (1:M)
- Reassignment creates a new history record and updates the existing assignment (not a new assignment row)
- Historical/completed assignments are preserved on the assignment record via status field

### 3.2 Design Decisions

| Decision | Choice | Rationale |
|----------|--------|-----------|
| One active assignment per user | Yes | Prevents conflicting reviews. A user account should have a single responsible analyst at any time. |
| Reassignment model | Update existing record + insert history row | Keeps the assignment table lean (one row per user-analyst pair) while preserving full audit trail in a dedicated history table. |
| Analyst auto-provisioning | Yes, on first Okta login | Analysts should not need manual pre-provisioning. The system creates an analyst record when an Okta-authenticated user first accesses the tool. |
| Analyst soft delete | Yes (`is_active = false`) | Analysts may return; hard delete would break FK references in assignment history. |
| Assignment status lifecycle | `assigned` -> `in_progress` -> `completed` (also: `reassigned` terminal state for old assignments) | Matches typical case workflow. `reassigned` marks the previous assignment as superseded. |
| Separate history table | Yes | Keeps the assignment table query-efficient for "who is assigned what right now" while providing full audit capability separately. |

---

## 4. Table Definitions

### 4.1 `ecdd_analyst`

Stores analyst (Okta user) records. Auto-created on first login.

| Column | Type | Nullable | Description |
|--------|------|----------|-------------|
| `ecdd_analyst_pk` | STRING(36) | No | UUID primary key |
| `username` | STRING(255) | No | Okta username (unique) |
| `email` | STRING(255) | No | Okta email address |
| `display_name` | STRING(255) | No | Full display name from Okta |
| `okta_groups` | STRING(MAX) | Yes | JSON array of Okta group names (determines feature access) |
| `role` | INT64 | No | Role level: 1=Analyst, 2=Senior Analyst, 3=Team Lead, 4=Manager, 5=Admin |
| `is_active` | BOOL | No | Whether the analyst account is active (soft delete) |
| `max_assignments` | INT64 | Yes | Optional cap on concurrent assignments (null = no limit) |
| `logged_at` | TIMESTAMP | No | Last modified timestamp |
| `last_login_at` | TIMESTAMP | Yes | Timestamp of most recent login |
| `updated_by` | STRING(255) | No | Who last modified this record |

**Unique constraint:** `username`

### 4.2 `ecdd_case_assignment`

Links a user account to an analyst. One active assignment per user at a time.

| Column | Type | Nullable | Description |
|--------|------|----------|-------------|
| `ecdd_case_assignment_pk` | STRING(36) | No | UUID primary key |
| `user_status_pk` | STRING(36) | No | FK to `ecdd_user_status.ecdd_user_status_pk` |
| `analyst_pk` | STRING(36) | No | FK to `ecdd_analyst.ecdd_analyst_pk` (assigned analyst) |
| `assigned_by_analyst_pk` | STRING(36) | No | FK to `ecdd_analyst.ecdd_analyst_pk` (who made the assignment) |
| `status` | INT64 | No | 1=Assigned, 2=In Progress, 3=Completed, 4=Reassigned |
| `priority` | INT64 | No | 1=Low, 2=Medium, 3=High, 4=Critical |
| `notes` | STRING(MAX) | Yes | Reason for assignment or review notes |
| `assigned_at` | TIMESTAMP | No | When the assignment was created |
| `started_at` | TIMESTAMP | Yes | When the analyst began working (status -> In Progress) |
| `completed_at` | TIMESTAMP | Yes | When review was completed (status -> Completed) |
| `logged_at` | TIMESTAMP | No | Last modified timestamp (follows existing convention) |
| `updated_by` | STRING(255) | No | Who last modified this record |

**Business rules:**
- Only one assignment with `status` IN (1, 2) per `user_status_pk` at any time (enforced at application level)
- When reassigning: old assignment status -> 4 (Reassigned), new assignment created with status 1

### 4.3 `ecdd_case_assignment_history`

Immutable audit log. One row inserted for every state change on an assignment.

| Column | Type | Nullable | Description |
|--------|------|----------|-------------|
| `ecdd_case_assignment_history_pk` | STRING(36) | No | UUID primary key |
| `case_assignment_pk` | STRING(36) | No | FK to `ecdd_case_assignment.ecdd_case_assignment_pk` |
| `action` | STRING(50) | No | Action performed: `created`, `status_changed`, `reassigned`, `priority_changed`, `notes_updated` |
| `old_status` | INT64 | Yes | Previous status value (null on creation) |
| `new_status` | INT64 | Yes | New status value |
| `old_analyst_pk` | STRING(36) | Yes | Previous analyst (for reassignment) |
| `new_analyst_pk` | STRING(36) | Yes | New analyst (for reassignment) |
| `notes` | STRING(MAX) | Yes | Additional context for this change |
| `performed_by` | STRING(255) | No | Username of person who performed the action |
| `performed_at` | TIMESTAMP | No | When the action was performed |

---

## 5. Go Model Structs

```go
package models

import "time"

// ECDDAnalyst represents an analyst (Okta user) who reviews ECDD cases.
// Auto-created on first Okta login if the analyst does not already exist.
type ECDDAnalyst struct {
	ECDDAnalystPK  string     `json:"ecdd_analyst_pk"`
	Username       string     `json:"username"`
	Email          string     `json:"email"`
	DisplayName    string     `json:"display_name"`
	OktaGroups     []string   `json:"okta_groups,omitempty"`  // JSON array of Okta group names
	Role           int64      `json:"role"`                   // 1=Analyst, 2=Senior Analyst, 3=Team Lead, 4=Manager, 5=Admin
	IsActive       bool       `json:"is_active"`
	MaxAssignments *int64     `json:"max_assignments,omitempty"` // Optional cap on concurrent assignments
	LoggedAt       time.Time  `json:"logged_at"`
	LastLoginAt    *time.Time `json:"last_login_at,omitempty"`
	UpdatedBy      string     `json:"updated_by"`
}

// ECDDCaseAssignment represents an assignment of a user account to an analyst for review.
type ECDDCaseAssignment struct {
	ECDDCaseAssignmentPK string     `json:"ecdd_case_assignment_pk"`
	UserStatusPK         string     `json:"user_status_pk"`           // FK to ECDDUserStatus
	AnalystPK            string     `json:"analyst_pk"`               // FK to ECDDAnalyst (assigned analyst)
	AssignedByAnalystPK  string     `json:"assigned_by_analyst_pk"`   // FK to ECDDAnalyst (who assigned)
	Status               int64      `json:"status"`                   // 1=Assigned, 2=In Progress, 3=Completed, 4=Reassigned
	Priority             int64      `json:"priority"`                 // 1=Low, 2=Medium, 3=High, 4=Critical
	Notes                *string    `json:"notes,omitempty"`
	AssignedAt           time.Time  `json:"assigned_at"`
	StartedAt            *time.Time `json:"started_at,omitempty"`
	CompletedAt          *time.Time `json:"completed_at,omitempty"`
	LoggedAt             time.Time  `json:"logged_at"`
	UpdatedBy            string     `json:"updated_by"`
}

// ECDDCaseAssignmentHistory represents an immutable audit record for assignment changes.
type ECDDCaseAssignmentHistory struct {
	ECDDCaseAssignmentHistoryPK string     `json:"ecdd_case_assignment_history_pk"`
	CaseAssignmentPK            string     `json:"case_assignment_pk"` // FK to ECDDCaseAssignment
	Action                      string     `json:"action"`             // created, status_changed, reassigned, priority_changed, notes_updated
	OldStatus                   *int64     `json:"old_status,omitempty"`
	NewStatus                   *int64     `json:"new_status,omitempty"`
	OldAnalystPK                *string    `json:"old_analyst_pk,omitempty"`
	NewAnalystPK                *string    `json:"new_analyst_pk,omitempty"`
	Notes                       *string    `json:"notes,omitempty"`
	PerformedBy                 string     `json:"performed_by"`
	PerformedAt                 time.Time  `json:"performed_at"`
}
```

---

## 6. SQL CREATE TABLE Statements (Google Cloud Spanner)

```sql
-- ============================================================
-- Table: ecdd_analyst
-- Stores analysts (Okta users) who use the AML tool.
-- ============================================================
CREATE TABLE ecdd_analyst (
    ecdd_analyst_pk         STRING(36)   NOT NULL,
    username                STRING(255)  NOT NULL,
    email                   STRING(255)  NOT NULL,
    display_name            STRING(255)  NOT NULL,
    okta_groups             STRING(MAX),           -- JSON array, e.g. '["ECDD_Analyst","ECDD_Admin"]'
    role                    INT64        NOT NULL,  -- 1=Analyst, 2=Senior Analyst, 3=Team Lead, 4=Manager, 5=Admin
    is_active               BOOL         NOT NULL DEFAULT (true),
    max_assignments         INT64,                  -- NULL = no limit
    logged_at               TIMESTAMP    NOT NULL,
    last_login_at           TIMESTAMP,
    updated_by              STRING(255)  NOT NULL,
) PRIMARY KEY (ecdd_analyst_pk);

-- Unique index on username (one analyst record per Okta user)
CREATE UNIQUE INDEX idx_analyst_username ON ecdd_analyst (username);

-- Index for querying active analysts
CREATE INDEX idx_analyst_active ON ecdd_analyst (is_active);


-- ============================================================
-- Table: ecdd_case_assignment
-- Links a user account (ECDDUserStatus) to an analyst for review.
-- ============================================================
CREATE TABLE ecdd_case_assignment (
    ecdd_case_assignment_pk   STRING(36)   NOT NULL,
    user_status_pk            STRING(36)   NOT NULL,  -- FK to ecdd_user_status
    analyst_pk                STRING(36)   NOT NULL,  -- FK to ecdd_analyst (assigned analyst)
    assigned_by_analyst_pk    STRING(36)   NOT NULL,  -- FK to ecdd_analyst (who assigned)
    status                    INT64        NOT NULL,  -- 1=Assigned, 2=In Progress, 3=Completed, 4=Reassigned
    priority                  INT64        NOT NULL,  -- 1=Low, 2=Medium, 3=High, 4=Critical
    notes                     STRING(MAX),
    assigned_at               TIMESTAMP    NOT NULL,
    started_at                TIMESTAMP,
    completed_at              TIMESTAMP,
    logged_at                 TIMESTAMP    NOT NULL,
    updated_by                STRING(255)  NOT NULL,
) PRIMARY KEY (ecdd_case_assignment_pk);

-- Query pattern: "all active assignments for a given user"
CREATE INDEX idx_assignment_user_status ON ecdd_case_assignment (user_status_pk, status);

-- Query pattern: "all assignments for a given analyst"
CREATE INDEX idx_assignment_analyst ON ecdd_case_assignment (analyst_pk, status);

-- Query pattern: "all assignments by status"
CREATE INDEX idx_assignment_status ON ecdd_case_assignment (status, assigned_at);

-- Query pattern: "all assignments by priority"
CREATE INDEX idx_assignment_priority ON ecdd_case_assignment (priority, status);


-- ============================================================
-- Table: ecdd_case_assignment_history
-- Immutable audit log of all assignment state changes.
-- ============================================================
CREATE TABLE ecdd_case_assignment_history (
    ecdd_case_assignment_history_pk  STRING(36)   NOT NULL,
    case_assignment_pk               STRING(36)   NOT NULL,  -- FK to ecdd_case_assignment
    action                           STRING(50)   NOT NULL,  -- created, status_changed, reassigned, priority_changed, notes_updated
    old_status                       INT64,
    new_status                       INT64,
    old_analyst_pk                   STRING(36),
    new_analyst_pk                   STRING(36),
    notes                            STRING(MAX),
    performed_by                     STRING(255)  NOT NULL,
    performed_at                     TIMESTAMP    NOT NULL,
) PRIMARY KEY (ecdd_case_assignment_history_pk);

-- Query pattern: "all history for a given assignment"
CREATE INDEX idx_history_assignment ON ecdd_case_assignment_history (case_assignment_pk, performed_at);

-- Query pattern: "all actions by a given user"
CREATE INDEX idx_history_performed_by ON ecdd_case_assignment_history (performed_by, performed_at);
```

---

## 7. API Endpoint Design

### 7.1 Analyst Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/v1/analysts` | List all analysts (supports `?is_active=true`, `?role=1`, sorting, pagination) |
| `GET` | `/api/v1/analysts/{id}` | Get analyst by PK |
| `GET` | `/api/v1/analysts/{id}/assignments` | Get all assignments for an analyst (supports `?status=1`, sorting, pagination) |
| `POST` | `/api/v1/analysts` | Create analyst (or auto-created on first Okta login) |
| `PUT` | `/api/v1/analysts/{id}` | Full update of analyst record |
| `PATCH` | `/api/v1/analysts/{id}` | Partial update of analyst record (e.g., update role, max_assignments) |
| `DELETE` | `/api/v1/analysts/{id}` | Soft delete (set `is_active = false`) |
| `POST` | `/api/v1/analysts/login` | Upsert analyst on Okta login (creates if new, updates `last_login_at` and groups if existing) |

### 7.2 Case Assignment Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/v1/case-assignments` | List all assignments (supports `?analyst_pk=`, `?user_status_pk=`, `?status=`, `?priority=`, sorting, pagination) |
| `GET` | `/api/v1/case-assignments/{id}` | Get assignment by PK |
| `GET` | `/api/v1/case-assignments/{id}/history` | Get audit history for an assignment |
| `POST` | `/api/v1/case-assignments` | Create new assignment |
| `PATCH` | `/api/v1/case-assignments/{id}` | Update assignment (status change, notes, priority) |
| `POST` | `/api/v1/case-assignments/{id}/reassign` | Reassign to a different analyst |
| `POST` | `/api/v1/case-assignments/bulk-assign` | Bulk assign multiple users to an analyst |
| `GET` | `/api/v1/case-assignments/stats` | Assignment statistics (count by status, by analyst, by priority) |

### 7.3 Nested Endpoints on Users

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/v1/users/{id}/assignment` | Get the current active assignment for a user (singular -- at most one active) |

---

## 8. Request/Response Examples

### 8.1 Create Assignment

**Request:** `POST /api/v1/case-assignments`

```json
{
    "user_status_pk": "u001-malta-134-001",
    "analyst_pk": "a1b2c3d4-0001-4a5b-8c9d-analyst00001",
    "assigned_by_analyst_pk": "a1b2c3d4-0001-4a5b-8c9d-manager00001",
    "priority": 3,
    "notes": "High-risk threshold breach detected. Requires immediate review.",
    "updated_by": "team.lead@company.com"
}
```

**Response:** `201 Created`

```json
{
    "ecdd_case_assignment_pk": "ca-uuid-generated-by-server",
    "user_status_pk": "u001-malta-134-001",
    "analyst_pk": "a1b2c3d4-0001-4a5b-8c9d-analyst00001",
    "assigned_by_analyst_pk": "a1b2c3d4-0001-4a5b-8c9d-manager00001",
    "status": 1,
    "priority": 3,
    "notes": "High-risk threshold breach detected. Requires immediate review.",
    "assigned_at": "2026-03-10T14:30:00Z",
    "logged_at": "2026-03-10T14:30:00Z",
    "updated_by": "team.lead@company.com"
}
```

### 8.2 Reassign

**Request:** `POST /api/v1/case-assignments/{id}/reassign`

```json
{
    "new_analyst_pk": "a1b2c3d4-0001-4a5b-8c9d-analyst00002",
    "reassigned_by_analyst_pk": "a1b2c3d4-0001-4a5b-8c9d-manager00001",
    "notes": "Previous analyst on leave. Reassigning to senior analyst.",
    "updated_by": "team.lead@company.com"
}
```

**Response:** `200 OK`

```json
{
    "old_assignment": {
        "ecdd_case_assignment_pk": "ca-uuid-original",
        "status": 4,
        "notes": "Reassigned: Previous analyst on leave."
    },
    "new_assignment": {
        "ecdd_case_assignment_pk": "ca-uuid-new",
        "user_status_pk": "u001-malta-134-001",
        "analyst_pk": "a1b2c3d4-0001-4a5b-8c9d-analyst00002",
        "assigned_by_analyst_pk": "a1b2c3d4-0001-4a5b-8c9d-manager00001",
        "status": 1,
        "priority": 3,
        "assigned_at": "2026-03-10T15:00:00Z",
        "logged_at": "2026-03-10T15:00:00Z",
        "updated_by": "team.lead@company.com"
    }
}
```

### 8.3 Analyst Login (Upsert)

**Request:** `POST /api/v1/analysts/login`

```json
{
    "username": "jane.smith",
    "email": "jane.smith@company.com",
    "display_name": "Jane Smith",
    "okta_groups": ["ECDD_Analyst", "ECDD_Malta_Region"]
}
```

**Response:** `200 OK` (existing analyst) or `201 Created` (new analyst)

```json
{
    "ecdd_analyst_pk": "a1b2c3d4-0001-4a5b-8c9d-analyst00001",
    "username": "jane.smith",
    "email": "jane.smith@company.com",
    "display_name": "Jane Smith",
    "okta_groups": ["ECDD_Analyst", "ECDD_Malta_Region"],
    "role": 1,
    "is_active": true,
    "logged_at": "2026-03-10T08:00:00Z",
    "last_login_at": "2026-03-10T08:00:00Z",
    "updated_by": "system"
}
```

---

## 9. Relationship to Existing Tables

### 9.1 Relationship to `ecdd_user_case_management_folder`

Folders and assignments serve **different purposes** and coexist:

| Concept | Folder Assignment | Case Assignment |
|---------|-------------------|-----------------|
| Purpose | Organizational grouping (categorize users) | Work assignment (who reviews whom) |
| Cardinality | User in many folders | User assigned to one analyst (active) |
| Owner | Folder is a bucket | Analyst is a person |
| Workflow | No workflow states | Has status lifecycle |
| Audit | No history table | Full history table |

A user can simultaneously be in a folder (e.g., "High Risk Customers") AND assigned to an analyst for review. These are independent relationships.

### 9.2 Integration with `updated_by`

The existing `updated_by` field on all tables stores freeform email strings. With the analyst table, the flow becomes:

1. Analyst logs in via Okta -> `ecdd_analyst` record created/updated
2. Analyst performs actions -> `updated_by` is set to the analyst's email/username
3. For case assignments specifically, the `analyst_pk` and `assigned_by_analyst_pk` fields provide proper FK-based tracking rather than freeform strings

### 9.3 Cascade Behavior

When an `ecdd_user_status` record is deleted:
- All related `ecdd_case_assignment` records should be cascade-deleted
- All related `ecdd_case_assignment_history` records should be preserved (orphaned but kept for audit -- the `user_status_pk` on the assignment can still be used for reference)

When an `ecdd_analyst` record is soft-deleted:
- Active assignments should be flagged for reassignment (not auto-deleted)
- The analyst record persists for FK integrity in history

---

## 10. Mock Data Service Extensions

The `DataService` struct in `services/data_service.go` would need the following additions:

```go
type DataService struct {
    // ... existing fields ...
    Analysts           map[string]*models.ECDDAnalyst
    CaseAssignments    map[string]*models.ECDDCaseAssignment
    AssignmentHistory  map[string]*models.ECDDCaseAssignmentHistory
    mu                 sync.RWMutex
}
```

New JSON data files needed:
- `data/analysts.json`
- `data/case_assignments.json`
- `data/case_assignment_history.json`

---

## 11. Status Enum Reference

### Assignment Status

| Value | Name | Description |
|-------|------|-------------|
| 1 | Assigned | Case has been assigned to analyst but work has not started |
| 2 | In Progress | Analyst has begun reviewing the case |
| 3 | Completed | Review is finished |
| 4 | Reassigned | This assignment was superseded by a reassignment (terminal state) |

### Analyst Role

| Value | Name | Description |
|-------|------|-------------|
| 1 | Analyst | Standard analyst |
| 2 | Senior Analyst | Experienced analyst, may handle escalations |
| 3 | Team Lead | Can assign/reassign cases within their team |
| 4 | Manager | Full assignment and configuration access |
| 5 | Admin | System administrator |

### Assignment Priority

| Value | Name | Description |
|-------|------|-------------|
| 1 | Low | Routine review |
| 2 | Medium | Standard priority |
| 3 | High | Requires prompt attention |
| 4 | Critical | Requires immediate attention (threshold breach, escalation) |

---

## 12. Open Questions for Discussion

1. **Assignment capacity**: Should the system enforce `max_assignments` on analysts, or is it advisory only? (Proposal: enforce at application level, return 409 Conflict if exceeded.)

2. **Auto-assignment**: Should there be a future endpoint for automatic load-balanced assignment (round-robin across analysts)? (Proposal: not in v1, but schema supports it via `max_assignments` and assignment counts.)

3. **Region-based assignment**: Should analysts be restricted to reviewing users from specific regions/countries? (Proposal: use Okta groups like `ECDD_Malta_Region` for this, check at application level. No schema change needed.)

4. **Completed assignment retention**: How long should completed assignments be retained? (Proposal: indefinitely in Spanner, with a future TTL policy if needed.)

5. **Concurrent reviews**: If a user is in "Completed" status and a new threshold breach occurs, should a new assignment be auto-created? (Proposal: this is a business logic decision for the service layer, not schema. The schema supports it -- just create a new assignment row.)
