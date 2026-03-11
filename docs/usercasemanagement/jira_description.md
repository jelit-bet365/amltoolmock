h2. ECDD User Case Management API - Implement User Case Management Endpoints

Implement the ECDD User Case Management API endpoints for managing user-to-folder assignments in the ECDD system. The API provides CRUD operations on the ECDDUserCaseManagementFolder Spanner table, folder-scoped user management (list users, add/remove single and bulk), and folder assignment statistics.

h3. Endpoints

||Method||Path||Description||
|GET|/api/ecdd/usercasemanagement|List all assignments with optional filtering, sorting, and pagination|
|POST|/api/ecdd/usercasemanagement|Create a new user-folder assignment|
|DELETE|/api/ecdd/usercasemanagement/\{usercasemanagementpk\}|Delete a single assignment by its primary key|
|DELETE|/api/ecdd/usercasemanagement?folder_pk=\{folder_pk\}|Delete all assignments for a given folder|
|GET|/api/ecdd/usercasemanagement/folder/\{folderpk\}/users|Get users assigned to a folder (always paginated)|
|DELETE|/api/ecdd/usercasemanagement/folder/\{folderpk\}/users/\{userstatuspk\}|Remove a single user from a folder|
|POST|/api/ecdd/usercasemanagement/folder/\{folderpk\}/users/bulk-delete|Bulk remove users from a folder|
|POST|/api/ecdd/usercasemanagement/folder/\{folderpk\}/users/bulk-add|Bulk add users to a folder|
|GET|/api/ecdd/usercasemanagement/folder/\{folderpk\}/stats|Get stats for a single folder|
|GET|/api/ecdd/usercasemanagement/stats|Get stats for all folders with optional filtering, sorting, and pagination|
