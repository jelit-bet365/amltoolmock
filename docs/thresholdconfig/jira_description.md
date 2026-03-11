h2. ECDD Threshold Config API - Implement Threshold Config Endpoints

Implement the ECDD Threshold Config API endpoints for managing ECDD threshold configuration records. The API provides read, update, patch, and delete operations for threshold configs stored in the ECDDThresholdConfig Spanner table.

h3. Endpoints

||Method||Path||Description||
|GET|/api/ecdd/thresholdconfig|List all threshold configs with optional filtering, sorting, and pagination|
|GET|/api/ecdd/thresholdconfig/\{thresholdconfigpk\}|Get a single threshold config by its primary key|
|POST|/api/ecdd/thresholdconfig|Create a new threshold config record|
|PUT|/api/ecdd/thresholdconfig/\{thresholdconfigpk\}|Full update of a threshold config record|
|PATCH|/api/ecdd/thresholdconfig/\{thresholdconfigpk\}|Partial update of a threshold config record|
|DELETE|/api/ecdd/thresholdconfig/\{thresholdconfigpk\}|Soft delete a threshold config record|
