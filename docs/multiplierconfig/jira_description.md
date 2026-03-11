h2. ECDD Multiplier Config API - Implement Multiplier Config Endpoints

Implement the ECDD Multiplier Config API endpoints for managing ECDD multiplier configuration records. The API provides read, update, patch, and delete operations for multiplier configs stored in the ECDDMultiplierConfig Spanner table.

h3. Endpoints

||Method||Path||Description||
|GET|/api/ecdd/multiplierconfig|List all multiplier configs with optional filtering, sorting, and pagination|
|GET|/api/ecdd/multiplierconfig/\{multiplierconfigpk\}|Get a single multiplier config by its primary key|
|POST|/api/ecdd/multiplierconfig|Create a new multiplier config record|
|PUT|/api/ecdd/multiplierconfig/\{multiplierconfigpk\}|Full update of a multiplier config record|
|PATCH|/api/ecdd/multiplierconfig/\{multiplierconfigpk\}|Partial update of a multiplier config record|
|DELETE|/api/ecdd/multiplierconfig/\{multiplierconfigpk\}|Soft delete a multiplier config record|
