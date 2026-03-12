# UI Update Plan

This document captures all API contract changes that the UI must adapt to, following the backend updates to Threshold Config, Case Management Folders, and User-Folder Assignments.

---

## API Changes Summary

### 1. Threshold Config — New `reinvest` Field

The `reinvest` boolean field is now present on all threshold config records in every response.

**UI changes required:**

- Display a reinvest toggle (checkbox or switch) in the threshold create and edit forms
- Show a reinvest column in the threshold list/table view
- Include `reinvest` (boolean) in POST, PUT, and PATCH request bodies when creating or updating thresholds
- Default the toggle to `false` if not yet set

**Affected endpoints:**

- `GET /api/ecdd/thresholdconfig` — `reinvest` now appears in every object in the response array
- `GET /api/ecdd/thresholdconfig/{pk}` — `reinvest` now appears in the single-record response
- `POST /api/ecdd/thresholdconfig` — `reinvest` must be included in the request body
- `PUT /api/ecdd/thresholdconfig/{pk}` — `reinvest` must be included in the request body

---

### 2. Case Management Folders — New `region`, `country_id`, `state_id` Fields

Three new fields are present on all folder records. `region` is required; `country_id` and `state_id` are optional/nullable.

**UI changes required:**

- Pass `?region=XX` as a query parameter when fetching folders for a region-specific view:
  - `GET /api/ecdd/casemanagementfolder?region=MALTA`
  - `GET /api/ecdd/usercasemanagement/stats?region=MALTA`
- Display `region` in the folder list and folder detail views
- Include `region` (required), `country_id` (optional), and `state_id` (optional) in create and update folder forms
- Validate that `region` is one of: MALTA, GIBRALTAR, USA, AUSTRALIA

**Affected endpoints:**

- `GET /api/ecdd/casemanagementfolder` — response objects now include `region`, `country_id`, `state_id`; supports new `?region=` query param
- `GET /api/ecdd/casemanagementfolder/{pk}` — response object now includes `region`, `country_id`, `state_id`
- `POST /api/ecdd/casemanagementfolder` — request body must include `region`; `country_id` and `state_id` are optional
- `PUT /api/ecdd/casemanagementfolder/{pk}` — request body must include `region`; `country_id` and `state_id` are optional

---

### 3. Folder Stats Response — New `region` Field

The `FolderAndStats` response object returned by the stats endpoints now includes a `region` field.

**UI changes required:**

- Read and display `region` from stats response objects where relevant
- Do not assume the stats endpoint returns all folders — it now returns only folders matching the requested region

**Affected endpoints:**

- `GET /api/ecdd/usercasemanagement/stats` — each stats object now includes `region`
- `GET /api/ecdd/usercasemanagement/folder/{pk}/stats` — response now includes `region`

---

### 4. User-Folder Assignments — FK Field Renamed

The foreign key field referencing the case management folder has been renamed from `folder_pk` to `case_management_folder_pk` in all JSON request and response bodies. This is a breaking rename — the old field name is no longer accepted or returned.

**UI changes required:**

- Update all code that reads `folder_pk` from assignment response objects to read `case_management_folder_pk` instead
- Update all code that sends `folder_pk` in assignment request bodies to send `case_management_folder_pk` instead
- Update any filter query parameters that used `?folder_pk=` to `?case_management_folder_pk=`

**Affected endpoints:**

- `GET /api/ecdd/usercasemanagement` — response field `folder_pk` is now `case_management_folder_pk`; filter query param renamed accordingly
- `POST /api/ecdd/usercasemanagement` — request body field `folder_pk` is now `case_management_folder_pk`
- `DELETE /api/ecdd/usercasemanagement` — query parameter `folder_pk` is now `case_management_folder_pk`
- `GET /api/ecdd/usercasemanagement/folder/{pk}/stats` — response field `folder_pk` retains the name for the stats PK reference, but the FK field in assignment records is renamed

---

### 5. Region Filtering — 4 Regions, UK Maps to MALTA

Region filtering on folder and stats endpoints is now correct. There are 4 ECDD regions: MALTA, GIBRALTAR, USA, AUSTRALIA. UK is not a standalone region — UK customers (country_id=197) fall under the MALTA region as a licensed country.

**UI changes required:**

- Use `?region=MALTA` (not `?region=UK`) when fetching folders and stats for the Malta region, which includes UK-licensed customers
- Remove any UI dropdown option or hardcoded value for "UK" as a region
- Ensure the region selector presents the 4 valid values: MALTA, GIBRALTAR, USA, AUSTRALIA

---

## Call Pattern Changes

The region filter on the stats endpoint now filters both folders and their user counts by region. The UI should consolidate its calls accordingly.

### Before (3 calls per page load)

```
GET /api/ecdd/casemanagementfolder                   → all folders (no region filter)
GET /api/ecdd/usercasemanagement/stats?region=MALTA   → all folders returned, but user counts show MALTA users only
GET /api/ecdd/usercasemanagement/folder/{id}/users?region=MALTA  → folder users (on expand)
```

The first two calls had to be merged by the UI, and the folder list contained folders from all regions with potentially zero relevant users.

### After (1-2 calls per page load)

```
GET /api/ecdd/usercasemanagement/stats?region=MALTA   → MALTA folders only, with MALTA-user counts (single source of truth)
GET /api/ecdd/usercasemanagement/folder/{id}/users?region=MALTA  → folder users (on expand)
```

The stats endpoint now returns only the folders that belong to the requested region, and user counts reflect only users in that region. The separate `GET /casemanagementfolder` call for listing is no longer needed when loading the folder stats view.

---

## Breaking Changes

The following changes are breaking and require UI code updates before the new backend is deployed:

| Change | Old value | New value | Impact |
|---|---|---|---|
| Assignment FK field name | `folder_pk` | `case_management_folder_pk` | Reads and writes break silently |
| Folder create/update | `region` not required | `region` is required | Create/update requests without `region` will be rejected (HTTP 400) |
| Stats endpoint folder scope | Returns all folders | Returns only region-matching folders | UI that expects all folders from stats will see fewer results — this is correct behaviour, not a bug |

---

## Non-Breaking Additions

The following changes add new fields that the UI can adopt incrementally without breaking existing functionality:

| Change | Notes |
|---|---|
| `reinvest` in threshold config responses | New field — safe to ignore until the UI adds the toggle |
| `region`, `country_id`, `state_id` in folder responses | New fields — safe to ignore until the UI adds display/edit support |
| `region` in folder stats responses | New field — safe to ignore until the UI adds display support |
| `?region=` query param on `GET /casemanagementfolder` | New param — existing calls without it continue to return all folders |
