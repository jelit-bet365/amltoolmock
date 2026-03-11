h2. ECDD User Status API - Implement User Status Endpoints

Implement the ECDD User Status API endpoints for managing user ECDD status records. The API provides read, update, patch, delete operations and folder assignment lookups for user statuses stored in the ECDDUserStatus Spanner table.

h3. Endpoints

||Method||Path||Description||
|GET|/api/ecdd/userstatus|List all user statuses with optional filtering, sorting, and pagination|
|GET|/api/ecdd/userstatus/\{userstatuspk\}|Get a single user status by its primary key|
|PUT|/api/ecdd/userstatus/\{userstatuspk\}|Full update of a user status record|
|PATCH|/api/ecdd/userstatus/\{userstatuspk\}|Partial update of a user status record|
|DELETE|/api/ecdd/userstatus/\{userstatuspk\}|Delete a user status record (cascade-deletes folder assignments)|
|GET|/api/ecdd/userstatus/\{userstatuspk\}/folders|Get case management folders assigned to a user status|
