# UI Migration Context — Backend Breaking Changes

This document lists every backend change that affects the frontend. Use it as a checklist when updating the UI codebase.

---

## 1. API Route Renames (ALL endpoints changed)

Every `/api/v1/` prefix is now `/api/ecdd/`. The path segments after the prefix also changed to match the Spanner table names.

| Old Frontend Path | New Backend Path | Endpoint File |
|---|---|---|
| `/api/v1/users` | `/api/ecdd/userstatus` | `usersApi.ts` |
| `/api/v1/users/{id}` | `/api/ecdd/userstatus/{userstatuspk}` | `usersApi.ts` |
| `/api/v1/thresholds` | `/api/ecdd/thresholdconfig` | `thresholdsApi.ts` |
| `/api/v1/thresholds/{id}` | `/api/ecdd/thresholdconfig/{thresholdconfigpk}` | `thresholdsApi.ts` |
| `/api/v1/folders` | `/api/ecdd/casemanagementfolder` | `foldersApi.ts` |
| `/api/v1/folders/{id}` | `/api/ecdd/casemanagementfolder/{folderpk}` | `foldersApi.ts` |
| `/api/v1/multipliers` | `/api/ecdd/multiplierconfig` | `multipliersApi.ts` |
| `/api/v1/multipliers/{id}` | `/api/ecdd/multiplierconfig/{multiplierconfigpk}` | `multipliersApi.ts` |
| `/api/v1/business-profiles` | `/api/ecdd/businessprofile` | `businessProfilesApi.ts` |
| `/api/v1/business-profiles/{id}` | `/api/ecdd/businessprofile/{businessprofilepk}` | `businessProfilesApi.ts` |
| `/api/v1/folder-assignments` | `/api/ecdd/usercasemanagement` | `folderAssignmentsApi.ts` |
| `/api/v1/folder-assignments/{id}` | `/api/ecdd/usercasemanagement/{usercasemanagementpk}` | `folderAssignmentsApi.ts` |
| `/api/v1/folders/{id}/users` | `/api/ecdd/usercasemanagement/folder/{folderpk}/users` | `foldersApi.ts` |
| `/api/v1/folders/{id}/users/{userId}` | `/api/ecdd/usercasemanagement/folder/{folderpk}/users/{userstatuspk}` | `folderAssignmentsApi.ts` |
| `/api/v1/folders/{id}/users/bulk-add` | `/api/ecdd/usercasemanagement/folder/{folderpk}/users/bulk-add` | `folderAssignmentsApi.ts` |
| `/api/v1/folders/{id}/users/bulk-delete` | `/api/ecdd/usercasemanagement/folder/{folderpk}/users/bulk-delete` | `folderAssignmentsApi.ts` |
| `/api/v1/folders/{id}/stats` | `/api/ecdd/usercasemanagement/folder/{folderpk}/stats` | `foldersApi.ts` |
| `/api/v1/folders/stats` | `/api/ecdd/usercasemanagement/stats` | `foldersApi.ts` |

### Files to Change

Update the `BASE_PATH` constant in each endpoint file:

| File | Old `BASE_PATH` | New `BASE_PATH` |
|---|---|---|
| `src/hooks/api/endpoints/usersApi.ts` | `/api/v1/users` | `/api/ecdd/userstatus` |
| `src/hooks/api/endpoints/thresholdsApi.ts` | `/api/v1/thresholds` | `/api/ecdd/thresholdconfig` |
| `src/hooks/api/endpoints/foldersApi.ts` | `/api/v1/folders` | `/api/ecdd/casemanagementfolder` |
| `src/hooks/api/endpoints/multipliersApi.ts` | `/api/v1/multipliers` | `/api/ecdd/multiplierconfig` |
| `src/hooks/api/endpoints/businessProfilesApi.ts` | `/api/v1/business-profiles` | `/api/ecdd/businessprofile` |
| `src/hooks/api/endpoints/folderAssignmentsApi.ts` | `/api/v1/folder-assignments` | `/api/ecdd/usercasemanagement` |

**CRITICAL**: `foldersApi.ts` and `folderAssignmentsApi.ts` have hardcoded `/api/v1/folders/` paths beyond just `BASE_PATH`. These must ALL be updated:

In `foldersApi.ts`:
- `getUsers`: path builds `${BASE_PATH}/${id}/users` — NOW this is a different service. Must change to `/api/ecdd/usercasemanagement/folder/${id}/users`
- `getStats`: path builds `${BASE_PATH}/${id}/stats` — must change to `/api/ecdd/usercasemanagement/folder/${id}/stats`
- `getAllStats`: path builds `${BASE_PATH}/stats` — must change to `/api/ecdd/usercasemanagement/stats`

In `folderAssignmentsApi.ts`:
- `bulkAssign`: hardcoded `` `/api/v1/folders/${folderId}/users/bulk-add` `` — must change to `` `/api/ecdd/usercasemanagement/folder/${folderId}/users/bulk-add` ``
- `removeFromFolder`: hardcoded `` `/api/v1/folders/${folderId}/users/${userId}` `` — must change to `` `/api/ecdd/usercasemanagement/folder/${folderId}/users/${userId}` ``
- `bulkRemoveFromFolder`: hardcoded `` `/api/v1/folders/${folderId}/users/bulk-delete` `` — must change to `` `/api/ecdd/usercasemanagement/folder/${folderId}/users/bulk-delete` ``

### Vite Proxy

If `vite.config.ts` has a proxy entry for `/api/v1`, update it to `/api/ecdd` (or just `/api`).

---

## 2. New Fields on Existing Types

### 2a. ECDDCaseManagementFolder — 3 new fields

Backend model now returns:

```json
{
  "ecdd_case_management_folder_pk": "uuid",
  "folder_name": "MT - Unusual Activity Reports",
  "region": "MALTA",
  "country_id": null,
  "state_id": null,
  "logged_at": "2026-01-15T10:00:00Z",
  "updated_by": "admin@company.com"
}
```

**Changes needed in `case.types.ts`** — `ECDDCaseManagementFolder` interface:

| Field | Type | Status |
|---|---|---|
| `region` | `string` | **ADD** — values: `"MALTA"`, `"GIBRALTAR"`, `"USA"`, `"AUSTRALIA"` |
| `country_id` | `number \| null` | **ADD** — nullable, usually null |
| `state_id` | `number \| null` | **ADD** — nullable, populated for US states only |
| `created_at` | — | **REMOVE** — backend uses `logged_at` not `created_at` |
| `updated_at` | — | **REMOVE** — backend uses `logged_at` not `updated_at` |
| `logged_at` | `string` | **ADD** (replaces `created_at`/`updated_at`) |

Current frontend type:
```typescript
export interface ECDDCaseManagementFolder {
  ecdd_case_management_folder_pk: string;
  folder_name: string;
  created_at: string;   // WRONG — backend sends logged_at
  updated_at: string;   // WRONG — backend sends logged_at
  updated_by: string;
}
```

Correct type:
```typescript
export interface ECDDCaseManagementFolder {
  ecdd_case_management_folder_pk: string;
  folder_name: string;
  region: string;
  country_id: number | null;
  state_id: number | null;
  logged_at: string;
  updated_by: string;
}
```

**Also update `CreateFolderInput` and `UpdateFolderInput`** in `foldersApi.ts`:
```typescript
export interface CreateFolderInput {
  folder_name: string;
  region: string;         // ADD — required
  country_id?: number;    // ADD — optional
  state_id?: number;      // ADD — optional
  updated_by: string;
}

export interface UpdateFolderInput {
  folder_name?: string;
  region?: string;        // ADD — optional
  country_id?: number;    // ADD — optional
  state_id?: number;      // ADD — optional
  updated_by: string;
}
```

### 2b. FolderStats / FolderAndStats — `region` field added

Backend stats response now includes `region`:

```json
{
  "folder_pk": "uuid",
  "folder_name": "MT - Unusual Activity Reports",
  "region": "MALTA",
  "user_count": 15,
  "oldest_user_date": "2025-09-20T00:00:00Z"
}
```

**Update `FolderStats` interface** in `foldersApi.ts`:
```typescript
export interface FolderStats {
  folder_pk: string;
  folder_name: string;
  region: string;              // ADD
  user_count: number;
  oldest_user_date: string | null;
}
```

### 2c. ECDDThresholdConfig — `reinvest` + many missing fields

Backend model has these fields that the frontend type is MISSING:

| Field | Type | Description |
|---|---|---|
| `reinvest` | `boolean` | **NEW** — whether to check reinvestment |
| `use_multipliers` | `boolean` | Apply ECDD multiplier adjustments |
| `use_rg_flag` | `boolean` | Consider ECDD Multiplier RG Flag |
| `apply_all_statuses` | `boolean` | Apply to all statuses or only "Not Required" |
| `backfill` | `boolean` | Apply to existing accounts |
| `hierarchy` | `number` | Priority order (lower = higher) |
| `ecdd_status` | `number` | ECDD Status on breach |
| `ecdd_review_status` | `number` | Review Status on breach |
| `ecdd_report_status` | `number` | Report Status on breach |
| `sign_off_status` | `number` | Sign Off Status on breach |
| `customer_risk_level` | `number` | 1=Low, 2=Medium, 3=Medium-High, 4=High |
| `ndl_28_day_gbp` | `number` | 28 Day NDL in GBP |
| `ndl_monthly_gbp` | `number` | Monthly NDL in GBP |
| `case_management_folder_pk` | `string \| null` | FK to case management folder |

Also, the backend uses `logged_at` (not `created_at`/`updated_at`), and has no `risk_id` field (frontend has `risk_id` which doesn't exist in backend — replace with `customer_risk_level`).

Full correct type:
```typescript
export interface ECDDThresholdConfig {
  ecdd_threshold_config_pk: string;
  title: string;
  is_active: boolean;
  country_id: number;
  state_id: number | null;
  type: number;                          // 1=Deposit, 2=Net Deposit, 3=Stakes
  reinvest: boolean;                     // ADD
  value: number;
  currency_id: number;
  period: number;                        // 1=24hrs, 2=28days, 3=84days, 4=91days, 5=182days, 6=365days
  use_multipliers: boolean;              // ADD
  use_rg_flag: boolean;                  // ADD
  apply_all_statuses: boolean;           // ADD
  backfill: boolean;                     // ADD
  hierarchy: number;                     // ADD
  ecdd_status: number;                   // ADD
  ecdd_review_status: number;            // ADD
  ecdd_report_status: number;            // ADD
  sign_off_status: number;               // ADD
  customer_risk_level: number;           // ADD (replaces risk_id)
  ndl_28_day_gbp: number;               // ADD
  ndl_monthly_gbp: number;              // ADD
  case_management_folder_pk: string | null; // ADD
  logged_at: string;                     // RENAME from created_at
  updated_by: string;
}
```

### 2d. ECDDUserCaseManagementFolder — FK key renamed

Backend JSON tag changed from `folder_pk` to `case_management_folder_pk`:

```json
{
  "ecdd_user_case_management_folder_pk": "uuid",
  "case_management_folder_pk": "folder-uuid",
  "user_status_pk": "user-uuid",
  "logged_at": "2026-01-15T10:00:00Z",
  "updated_by": "admin@company.com"
}
```

**Changes needed in `case.types.ts`**:
```typescript
// BEFORE
export interface ECDDUserCaseManagementFolder {
  ecdd_user_case_management_folder_pk: string;
  folder_pk: string;              // OLD key name
  user_status_pk: string;
  created_at: string;             // WRONG — backend sends logged_at
  updated_at: string;             // WRONG — backend sends logged_at
  updated_by: string;
}

// AFTER
export interface ECDDUserCaseManagementFolder {
  ecdd_user_case_management_folder_pk: string;
  case_management_folder_pk: string;  // RENAMED from folder_pk
  user_status_pk: string;
  logged_at: string;                  // RENAMED from created_at/updated_at
  updated_by: string;
}
```

**Also update `AssignUserToFolderInput`** in `folderAssignmentsApi.ts`:
```typescript
// BEFORE
export interface AssignUserToFolderInput {
  folder_pk: string;          // OLD
  user_status_pk: string;
  updated_by: string;
}

// AFTER
export interface AssignUserToFolderInput {
  case_management_folder_pk: string;  // RENAMED
  user_status_pk: string;
  updated_by: string;
}
```

**Also update `getByFolderId`** in `folderAssignmentsApi.ts`:
```typescript
// BEFORE
getByFolderId: (folderId: string) =>
  client.get(`${BASE_PATH}?folder_pk=${folderId}`)

// AFTER
getByFolderId: (folderId: string) =>
  client.get(`${BASE_PATH}?case_management_folder_pk=${folderId}`)
```

---

## 3. Region Config — Wrong Country IDs

`src/config/regions.ts` has completely wrong country IDs. These must be corrected:

| Region | Old (Wrong) IDs | Correct IDs |
|---|---|---|
| MALTA | `[134, 135, 136]` (Nepal, Netherlands, NetherlandsAntilles) | See full list below |
| USA | `[231, 232, 233]` (VaticanCity, Wallis, ChristmasIsland) | `[198]` (US only) |
| AUSTRALIA | `[13, 14, 15]` (Australia, Austria, Azerbaijan) | `[13]` (Australia only) |
| GIBRALTAR | `[83, 84, 85]` (Guinea, GuineaBissau, Guyana) | See full list below |

### Malta Country IDs (43 total = 23 regulated + 20 licensed)

**Regulated (23):**
Austria=14, Croatia=49, Finland=68, Ghana=76, Hungary=88, Iceland=89, IrelandRepOf=95, IvoryCoast=50, India=90, Kenya=103, Latvia=108, Liechtenstein=111, Lithuania=112, Luxembourg=113, Malta=120, NewZealand=139, Nicaragua=141, Norway=147, Slovakia=163, Slovenia=164, Tanzania=184, Tonga=192, WesternSamoa=207

**Licensed (20):**
UK=197, Spain=171, Italy=97, Denmark=54, Greenland=218, Sweden=181, Estonia=64, Bulgaria=31, Cyprus=51, Greece=78, Germany=75, Netherlands=135, CzechRepublic=52, Mexico=126, CanadaOntario=272, BuenosAiresCity=270, BuenosAiresProvince=271, France=70, Switzerland=174, Japan=99

**USA:** `[198]`

**Australia:** `[13]`

**Gibraltar (~108 countries):** Albania=2, Algeria=3, Andorra=5, Anguilla=6, AntiguaAndBarbuda=7, Argentina=10, Armenia=11, Aruba=12, Azerbaijan=15, Bahamas=16, Bahrain=17, Bangladesh=18, Barbados=19, Belarus=20, Belize=22, Benin=23, Bermuda=24, Bolivia=26, BosniaHerzegovina=27, Botswana=28, Brazil=29, BritishVirginIslands=30, BruneiDarussalam=225, BurkinaFaso=32, Cameroon=35, Canada=36, CapeVerdeIslands=37, CaymanIslands=38, CentralAfricanRepublic=39, Chile=41, CookIslands=47, CostaRica=48, Djibouti=56, Dominica=57, DominicanRepublic=58, Ecuador=60, Egypt=61, ElSalvador=62, Ethiopia=65, FaroeIslands=66, Fiji=67, FrenchPolynesia=71, Gabon=72, Gambia=73, Georgia=74, Gibraltar=77, Grenada=79, Guatemala=80, Guinea=83, Guyana=85, Honduras=87, Indonesia=91, Jamaica=98, Jordan=100, Kazakhstan=101, SouthKorea=167, Kuwait=105, Kyrgyzstan=106, Laos=107, Lebanon=109, Lesotho=110, Liberia=222, Macedonia=117, Madagascar=118, Malawi=119, Malaysia=260, Maldives=121, Mali=122, Mauritania=124, Mauritius=125, Moldova=128, Mongolia=129, Montenegro=131, Montserrat=223, Morocco=132, Mozambique=133, Namibia=134, Nepal=136, NetherlandsAntilles=137, NewCaledonia=138, Niger=142, Nigeria=143, Oman=148, Pakistan=149, Palestine=151, Panama=152, PapuaNewGuinea=153, Paraguay=154, Peru=155, Qatar=157, Rwanda=159, SanMarino=161, SaoTomeEPrincipe=224, SaudiArabia=162, Senegal=259, Serbia=261, Seychelles=262, SierraLeone=263, SolomonIslands=165, SriLanka=172, StKittsAndNevis=264, StLucia=265, StVincentAndTheGrenadines=266, Suriname=176, Swaziland=178, Taiwan=183, Thailand=186, Togo=190, TrinidadAndTobago=193, Tunisia=195, TurksAndCaicosIslands=267, Uganda=196, Ukraine=268, UnitedArabEmirates=269, Uruguay=200, Vanuatu=202, Vietnam=203, Zambia=205

### Decision: Remove `countryIds` from frontend regions

The frontend should NOT maintain its own country-to-region mapping. The backend now handles all region filtering server-side via the `?region=` query parameter. The `countryIds` array in `regions.ts` is:
1. Currently wrong (all IDs are wrong)
2. Unnecessary (backend filters by region, not the frontend)
3. Unmaintainable (Malta alone has 43 countries)

**Recommended approach**: Simplify `regions.ts` to only define region keys, names, and display info. Remove `countryIds` entirely. All filtering should use `?region=MALTA` (etc.) and let the backend resolve countries.

```typescript
export interface Region {
  name: string;
  code: string;
  currency: string;
}

export const REGIONS: Record<string, Region> = {
  MALTA: { name: 'Malta', code: 'MT', currency: 'EUR' },
  USA: { name: 'United States', code: 'US', currency: 'USD' },
  AUSTRALIA: { name: 'Australia', code: 'AU', currency: 'AUD' },
  GIBRALTAR: { name: 'Gibraltar', code: 'GI', currency: 'GBP' },
};
```

Remove these functions (no longer needed):
- `getRegionByCountryId()` — backend does this
- `isCountryInRegion()` — backend does this
- `getCountryIdsForRegion()` — backend does this
- `getPrimaryCountryId()` — backend does this

Keep: `REGION_KEYS`, `REGION_NAMES`, `getRegionKeyByName()` — UI display helpers.

---

## 4. Country Name Display — Wrong Mapping

`useFolderUsersColumns.tsx` has a hardcoded country name map with wrong IDs:

```typescript
// CURRENT (WRONG)
UK = 1, Germany = 2, ...

// CORRECT
UK = 197, Germany = 75, Spain = 171, Italy = 97, ...
```

**Action**: Create a proper country name lookup (or fetch from backend). At minimum, fix the hardcoded map to use correct IDs from the backend's `enums/country/country.go` package.

---

## 5. Enum Value Changes

### 5a. ThresholdType enum is wrong

```typescript
// CURRENT (WRONG)
export enum ThresholdType {
  DEPOSIT = 1,
  STAKE = 2,
  COMBINED = 3,
  LOSS = 4
}

// CORRECT (from backend)
// 1=Deposit, 2=Net Deposit, 3=Stakes
// No "COMBINED" or "LOSS"
```

### 5b. Period enum is wrong

```typescript
// CURRENT (WRONG)
export enum Period {
  DAILY = 1, WEEKLY = 2, MONTHLY = 3, QUARTERLY = 4, ANNUAL = 5
}

// CORRECT (from backend)
// 1=24hrs, 2=28days, 3=84days, 4=91days, 5=182days, 6=365days
```

### 5c. ECDDStatus enum is wrong

```typescript
// CURRENT (WRONG)
export enum ECDDStatus {
  ACTIVE = 1, UNDER_REVIEW = 2, FLAGGED = 3, SUSPENDED = 4, BLOCKED = 5
}

// CORRECT (from backend)
// 1=Not Required, 2=In Progress, 3=Complete, 4=Suspended-Manual,
// 5=Suspended-Auto, 6=Closed, 7=Block Process
```

### 5d. RiskStatus enum is wrong

```typescript
// CURRENT (WRONG)
export enum RiskStatus { LOW = 1, MEDIUM = 2, HIGH = 3, VERY_HIGH = 4 }

// CORRECT (from backend)
// 1=Low, 2=Medium, 3=Medium-High, 4=High
// "VERY_HIGH" should be "MEDIUM_HIGH" at position 3, "HIGH" at position 4
```

---

## 6. Behavioral Changes

### 6a. Stats endpoint now filters folders by region

**Before**: `GET /api/v1/folders/stats?region=MALTA` returned ALL folders, with user counts filtered by the region's countries. Non-region folders showed `user_count: 0`.

**After**: `GET /api/ecdd/usercasemanagement/stats?region=MALTA` returns ONLY folders whose `region` field is `"MALTA"`. Non-MALTA folders are excluded from the response entirely.

**Impact**: The frontend will receive fewer results (68 Malta folders instead of all 120). If the UI renders "No users in folder" for zero-count items, that noise is now gone.

### 6b. Folder volume increased: 10 → 120

Mock data now has 120 folders (68 Malta, 36 USA, 12 Australia, 4 Gibraltar) instead of the previous ~10. The `FolderListTab` component renders all folders with `.map()` without pagination or virtualization.

**Action needed**: Consider adding:
- Client-side search/filter within the folder list
- Virtualized scrolling (e.g., `react-virtualized` or `@tanstack/react-virtual`)
- Or rely on the `?region=` filter to get a regional subset (max 68 for Malta)

### 6c. UK users now appear under Malta

UK (country_id=197) is a Malta-licensed country. When filtering by `?region=MALTA`, UK users are included in the results. The frontend region selector should NOT offer "UK" as a separate option (it never did — good). But any UK-specific display logic must account for UK being part of Malta.

---

## 7. Timestamp Field Name Changes

The backend consistently uses `logged_at` (single timestamp), NOT `created_at`/`updated_at` (separate timestamps). All entity types in the frontend that reference `created_at` or `updated_at` must be updated to `logged_at`.

| Entity | Frontend Has | Backend Sends |
|---|---|---|
| ECDDCaseManagementFolder | `created_at`, `updated_at` | `logged_at` |
| ECDDUserCaseManagementFolder | `created_at`, `updated_at` | `logged_at` |
| ECDDThresholdConfig | `created_at`, `updated_at` | `logged_at` |
| ECDDMultiplierConfig | `created_at`, `updated_at` | `logged_at` |
| ECDDBusinessProfile | `created_at`, `updated_at` | `logged_at` |
| ECDDUserStatus | `logged_at` | `logged_at` (already correct) |

---

## 8. Summary Checklist

### Must Fix (Breaking)
- [ ] Update all 6 endpoint files with new `BASE_PATH` values
- [ ] Fix hardcoded `/api/v1/folders/` paths in `foldersApi.ts` and `folderAssignmentsApi.ts`
- [ ] Rename `folder_pk` to `case_management_folder_pk` in `ECDDUserCaseManagementFolder` type and all query params
- [ ] Fix `created_at`/`updated_at` → `logged_at` in all entity types
- [ ] Add `region`, `country_id`, `state_id` to `ECDDCaseManagementFolder` type
- [ ] Add `region` to `FolderStats` interface
- [ ] Fix region `countryIds` in `regions.ts` (or remove them entirely)
- [ ] Update vite proxy if applicable

### Should Fix (Data Correctness)
- [ ] Fix country name mapping in `useFolderUsersColumns.tsx` (wrong IDs)
- [ ] Fix `ECDDStatus` enum values (wrong labels)
- [ ] Fix `ThresholdType` enum values (wrong types)
- [ ] Fix `Period` enum values (wrong periods)
- [ ] Fix `RiskStatus` enum values (wrong label at position 3)
- [ ] Add missing fields to `ECDDThresholdConfig` type (including `reinvest`)
- [ ] Replace `risk_id` with `customer_risk_level` in threshold type

### Should Consider (UX)
- [ ] Handle 120 folders in `FolderListTab` (pagination or virtualization)
- [ ] Add `region` column/badge to folder list display
- [ ] Update `CreateFolderInput` / `UpdateFolderInput` to include `region`
- [ ] Remove `getRegionByCountryId()` and related functions from `regions.ts`
