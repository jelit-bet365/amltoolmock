h2. ECDD Business Profile API - Implement Business Profile Endpoints

Implement the ECDD Business Profile API endpoints for managing ECDD business profile records. The API provides read, update, patch, and delete operations for business profiles stored in the ECDDBusinessProfile Spanner table.

h3. Endpoints

||Method||Path||Description||
|GET|/api/ecdd/businessprofile|List all business profiles with optional filtering, sorting, and pagination|
|GET|/api/ecdd/businessprofile/\{businessprofilepk\}|Get a single business profile by its primary key|
|POST|/api/ecdd/businessprofile|Create a new business profile record|
|PUT|/api/ecdd/businessprofile/\{businessprofilepk\}|Full update of a business profile record|
|PATCH|/api/ecdd/businessprofile/\{businessprofilepk\}|Partial update of a business profile record|
|DELETE|/api/ecdd/businessprofile/\{businessprofilepk\}|Soft delete a business profile record|
