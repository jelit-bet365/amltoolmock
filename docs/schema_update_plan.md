# Schema Update Plan — reinvest + Region-Specific Folders

## 1. Schema Update: Threshold `reinvest` Field

### Current State
- `docs/thresholdconfig/api.md:64` — `reinvest` is already documented in all request/response examples
- `models/threshold_config.go` — **field is MISSING** from the struct
- `data/threshold_configs.json` — **field is MISSING** from all 11 records

### Impact: LOW

| Component | Change Needed | Breaking? |
|---|---|---|
| `models/threshold_config.go` | Add `Reinvest bool json:"reinvest"` between `Type` and `Value` | No |
| `handlers/threshold_handler.go` | **None** — all CRUD uses `json.Encode/Decode` on the struct | No |
| `PatchThreshold` handler | **None** — decodes onto existing struct, `reinvest` automatically works | No |
| `sortThresholds` | No sort needed (boolean) | N/A |
| `GetAllThresholds` filter | Optional `?reinvest=` filter, not in docs | Optional |
| Data service | **None** — model change propagates | No |
| UI | Add `reinvest` toggle in threshold create/edit forms and list columns | Minor |

### Silent Bug
Until the model is updated, the API silently drops `reinvest` on POST/PUT and never returns it on GET. The UI already expects it per the docs — meaning the UI receives data that doesn't match its spec. Any UI code reading `reinvest` gets `undefined`.

---

## 2. Schema Update: Region-Specific Case Folders

### Current State
- `models/case_management_folder.go` — **NO `region` field**
- `data/case_folders.json` — 10 folders, all "global" (no region scoping)
- `handlers/folder_users_handler.go:17-22` — hardcoded `RegionCountries` map with **WRONG IDs**:
  ```
  "MALTA":     {134, 135, 136}  → Actually Nepal, Netherlands, NetherlandsAntilles
  "USA":       {231, 232, 233}  → Actually VaticanCity, Wallis, ChristmasIsland
  "AUSTRALIA": {13, 14, 15}     → 13=Australia correct, 14=Austria, 15=Azerbaijan
  "GIBRALTAR": {83, 84, 85}     → Actually Guinea, GuineaBissau, Guyana
  ```
  Note: There are 4 ECDD regions — MALTA, GIBRALTAR, USA, AUSTRALIA. UK is not a standalone region; UK (country_id=197) falls under the MALTA region as a licensed country.
- `enums/region/region.go` — **already created** with correct mappings

### Impact: HIGH

#### Current (Broken) UI Call Pattern
```
Call 1: GET /api/v1/folders/stats?region=MALTA
        → Returns ALL 10 folders, user counts filtered by MALTA users
        → Non-MALTA folders show 0 users (noise)
        → RegionCountries map has wrong IDs, so filter is effectively broken

Call 2: GET /api/v1/folders/{id}/users?region=MALTA
        → Gets paginated users for one folder, filtered by MALTA
        → Same broken region filtering

Call 3: GET /api/v1/thresholds?country_id=<wrong_id>
        → Gets thresholds (with wrong country IDs)
```

#### Target UI Call Pattern (After Changes)
```
Call 1: GET /api/v1/folders/stats?region=MALTA
        → Returns ONLY MALTA-region folders with MALTA-user stats (includes UK country_id=197)
        → Single source of truth for folder sidebar

Call 2: GET /api/v1/folders/{id}/users?region=MALTA  (on folder expand)
        → Gets paginated users, correctly filtered by MALTA country IDs (includes UK)
```

### Handler Changes Required

| Handler | Function | Change |
|---|---|---|
| `folder_handler.go` | `GetAllFolders` | Add `?region=` filter on folder's `Region` field |
| `folder_handler.go` | `GetAllFolderAndStats` | Filter **folders** by region (not just users) |
| `folder_handler.go` | `FolderAndStats` struct | Add `Region string json:"region"` |
| `folder_handler.go` | `sortFolders` | Add `region` sort option |
| `folder_handler.go` | `CreateFolder` | Accept `region` in body |
| `folder_handler.go` | `UpdateFolder` | Accept `region` in body |
| `folder_users_handler.go` | `RegionCountries` var | **Delete**, replace with `region.GetCountriesForRegion()` |
| `folder_users_handler.go` | `FilterUsers` | Use enum package for region→country mapping |

---

## 3. Cross-Cutting: All country_ids in Mock Data Are Wrong

Every data file uses sequential `country_id` values (1, 2, 3...) instead of production enum values:

| Data File | Current IDs | Correct IDs |
|---|---|---|
| `threshold_configs.json` | 1-7 | UK=197, US=198, Ireland=95, Canada=36, Australia=13, Germany=75, Spain=171 |
| `user_statuses.json` | Sequential per region | Real enum values per region |
| `multiplier_configs.json` | 1-5 | Real enum values |
| `business_profiles.json` | 1-6 | Real enum values |

---

## 4. Threshold → Folder FK Consistency

With region-specific folders, each threshold's `case_management_folder_pk` must reference a folder in the matching region:

| Threshold | Country | Region | Folder Must Be |
|---|---|---|---|
| UK 28 day NDL 50K | 197 (UK) | MALTA | MALTA-region folder |
| US-NY High Stakes | 198 (US) | USA | USA-region folder |
| Ireland 28 day NDL | 95 (Ireland) | MALTA | MALTA-region folder |
| Canada 182 day Deposit | 36 (Canada) | GIBRALTAR | GIBRALTAR-region folder |
| Australia 84 day NDL | 13 (Australia) | AUSTRALIA | AUSTRALIA-region folder |
| Germany High Risk 28 day | 75 (Germany) | MALTA | MALTA-region folder |
| Spain 365 day Stakes | 171 (Spain) | GIBRALTAR | GIBRALTAR-region folder |

Folders like "High Risk Customers" can no longer be shared across regions. Each region needs its own set of folders.

---

## 5. What Breaks If We Do Nothing

| Issue | Severity | Impact |
|---|---|---|
| `reinvest` missing from model | **Medium** | UI shows wrong toggle state for all thresholds |
| No `region` on folders | **High** | UI can't scope folder views by region |
| `RegionCountries` has wrong IDs | **Critical** | Every region-filtered API call returns incorrect data |
| All `country_id` values wrong | **High** | Thresholds, multipliers, profiles reference nonexistent countries |

---

## 6. Recommended Implementation Order

### Step 1: Model Updates
- Add `Reinvest bool` to `ECDDThresholdConfig`
- Add `Region string` to `ECDDCaseManagementFolder`

### Step 2: Fix Region Filtering
- Delete `RegionCountries` hardcoded map
- Update `FilterUsers` to use `enums/region` package

### Step 3: Add Region Filter to Folder Endpoints
- `GetAllFolders`: add `?region=` filter
- `GetAllFolderAndStats`: filter folders by region, add `Region` to `FolderAndStats` struct

### Step 4: Regenerate Mock Data
- All data files with correct country IDs from enums
- Region field on all folders (region-specific folder names)
- `reinvest` field on all thresholds
- Consistent FK relationships (threshold → folder region match)

### Step 5: Update Documentation
- `docs/thresholdconfig/api.md` — already correct (has reinvest)
- `docs/thresholdconfig/mock_data_setup.md` — update with reinvest
- `docs/casemanagementfolder/api.md` — add region field, add ?region= filter
- `docs/casemanagementfolder/mock_data_setup.md` — update with region
- `docs/usercasemanagement/api.md` — no schema changes, but document region filter fix

### Step 6: UI Update Plan
- Document all API contract changes for the UI agent

---

## 7. Optional Enhancements

### A. Add `?region=` to Thresholds Endpoint
Currently thresholds filter by `?country_id=`. Adding `?region=MALTA` would map region to country IDs via `region.GetCountriesForRegion()`, eliminating the need for the UI to know country-to-region mapping.

### B. Folder Region Validation on User Assignment
When bulk-adding users to a folder, optionally validate that users' countries match the folder's region. For mock purposes, keep this as display-only (Option A — no validation).
