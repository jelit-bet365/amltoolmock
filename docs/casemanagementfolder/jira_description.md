h2. ECDD Case Management Folder API - Implement Case Management Folder Endpoints

Implement the ECDD Case Management Folder API endpoints for managing ECDD case management folders. The API provides list, read, create, update, and delete operations for folders stored in the ECDDCaseManagementFolder Spanner table.

h3. Endpoints

||Method||Path||Description||
|GET|/api/ecdd/casemanagementfolder|List all folders with optional filtering, sorting, and pagination|
|GET|/api/ecdd/casemanagementfolder/\{folderpk\}|Get a single folder by its primary key|
|POST|/api/ecdd/casemanagementfolder|Create a new case management folder|
|PUT|/api/ecdd/casemanagementfolder/\{folderpk\}|Full update of a folder record|
|DELETE|/api/ecdd/casemanagementfolder/\{folderpk\}|Delete a folder record (cascade-deletes user-folder assignments)|
