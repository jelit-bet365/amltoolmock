# ECDD Mock API Service

A Go-based mock API service for Enhanced Customer Due Diligence (ECDD) system, providing RESTful endpoints for managing threshold configurations, user statuses, case folders, multiplier configs, business profiles, and user-folder assignments.

## Features

- ✅ **RESTful API** with 25+ endpoints
- ✅ **CORS enabled** for React frontend integration
- ✅ **In-memory data store** with thread-safe operations
- ✅ **Mock data** with realistic relationships (120 users, 10 folders, 240 user-folder mappings)
- ✅ **Pure Go** using standard library (no external dependencies except UUID)
- ✅ **Port 3001** for easy integration

## Project Structure

```
amltoolmock/
├── main.go                          # Entry point & HTTP server
├── go.mod                           # Go module definition
├── models/                          # Data models
│   ├── threshold_config.go
│   ├── user_status.go
│   ├── case_management_folder.go
│   ├── multiplier_config.go
│   ├── business_profile.go
│   └── user_case_folder.go
├── handlers/                        # HTTP request handlers
│   ├── user_handler.go
│   ├── threshold_handler.go
│   └── handlers.go
├── services/                        # Business logic
│   └── data_service.go
├── middleware/                      # HTTP middleware
│   └── cors.go
├── utils/                           # Utilities
│   └── pagination.go
└── data/                            # Mock JSON data
    ├── threshold_configs.json
    ├── user_statuses.json
    ├── case_folders.json
    ├── multiplier_configs.json
    ├── business_profiles.json
    ├── user_case_folders.json
    └── scripts/
        └── generate_users.js        # Script to regenerate mock data
```

## Quick Start

### Prerequisites

- Go 1.21 or higher

### Installation

```bash
# Clone or navigate to the project directory
cd amltoolmock

# Install dependencies
go mod tidy

# Run the server
go run main.go
```

The server will start on `http://localhost:3001`

## API Endpoints

### Health Check

```
GET /health
```

**Response:**
```json
{
  "status": "healthy",
  "service": "ECDD Mock API"
}
```

### User Status Endpoints

#### Get All Users
```
GET /api/v1/users
```

**Response:** Array of all user objects (120 users total)

#### Get User by ID
```
GET /api/v1/users/{id}
```

#### Create User
```
POST /api/v1/users
Content-Type: application/json

{
  "user_id": 100121,
  "user_name": "Jane Doe",
  "country_id": 134,
  "ecdd_status": 1,
  "ecdd_threshold": 25000.00,
  ...
}
```

#### Update User
```
PUT /api/v1/users/{id}
Content-Type: application/json

{
  "user_name": "Jane Updated",
  "ecdd_status": 2,
  ...
}
```

### Threshold Configuration Endpoints

#### Get All Thresholds
```
GET /api/v1/thresholds
```

#### Get Threshold by ID
```
GET /api/v1/thresholds/{id}
```

#### Create Threshold
```
POST /api/v1/thresholds
Content-Type: application/json

{
  "title": "New Threshold Rule",
  "is_active": true,
  "country_id": 134,
  "type": 2,
  "value": 30000.00,
  "currency_id": 1,
  "period": 2,
  ...
}
```

#### Update Threshold
```
PUT /api/v1/thresholds/{id}
Content-Type: application/json
```

#### Delete Threshold (Soft Delete)
```
DELETE /api/v1/thresholds/{id}
```

### Case Folder Endpoints

#### Get All Folders
```
GET /api/v1/folders
```

#### Get Folder by ID
```
GET /api/v1/folders/{id}
```

#### Create Folder
```
POST /api/v1/folders
Content-Type: application/json

{
  "folder_name": "New Risk Category",
  "updated_by": "admin@company.com"
}
```

#### Update Folder
```
PUT /api/v1/folders/{id}
Content-Type: application/json
```

### Multiplier Configuration Endpoints

#### Get All Multipliers
```
GET /api/v1/multipliers
```

#### Get Multiplier by ID
```
GET /api/v1/multipliers/{id}
```

#### Create Multiplier
```
POST /api/v1/multipliers
Content-Type: application/json

{
  "country_id": 134,
  "age_multipliers": [18, 19, 20, 21, 22],
  "status_multiplier": true,
  "is_active": true,
  "updated_by": "admin@company.com"
}
```

#### Update Multiplier
```
PUT /api/v1/multipliers/{id}
Content-Type: application/json
```

### Business Profile Endpoints

#### Get All Business Profiles
```
GET /api/v1/business-profiles
```

#### Get Business Profile by ID
```
GET /api/v1/business-profiles/{id}
```

#### Create Business Profile
```
POST /api/v1/business-profiles
Content-Type: application/json

{
  "country_id": 134,
  "risk_status_id": 2,
  "average_deposit": 5000.00,
  "deposit_multiplier": 2.00,
  "time_period_days": 28,
  "enabled": true,
  "updated_by": "admin@company.com"
}
```

#### Update Business Profile
```
PUT /api/v1/business-profiles/{id}
Content-Type: application/json
```

### User-Folder Assignment Endpoints

#### Get All User-Folder Assignments
```
GET /api/v1/user-folders
```

#### Create Assignment
```
POST /api/v1/user-folders
Content-Type: application/json

{
  "folder_id": "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d",
  "user_id": "u001-malta-134-001",
  "updated_by": "admin@company.com"
}
```

#### Delete Assignment
```
DELETE /api/v1/user-folders/{id}
```

## Data Models

### ECDDUserStatus
Represents user ECDD status with monitoring thresholds and review statuses.

### ECDDThresholdConfig
Threshold rules for deposit/stake monitoring with configurable periods and risk levels.

### ECDDCaseManagementFolder
Case folders for organizing users by risk category or review status.

### ECDDMultiplierConfig
Age and status-based threshold multiplier rules (e.g., 0.5x for ages 18-25).

### ECDDBusinessProfile
Business profile rules for risk-based deposit monitoring.

### ECDDUserCaseManagementFolder
Junction table for many-to-many user-folder assignments.

## Region to Country ID Mapping

**Important:** Each region maps to **multiple country IDs** for regulatory and jurisdictional purposes.

The system contains **120 users** distributed across **4 regions** with **30 users per region**.

### Region Overview

| Region | Country IDs | Total Users | State Support | Regulatory Body |
|--------|------------|-------------|---------------|-----------------|
| **Malta** | 134, 135, 136 | 30 | No | Malta Gaming Authority (MGA) |
| **USA** | 231, 232, 233 | 30 | Yes | State Gaming Commissions |
| **Australia** | 13, 14, 15 | 30 | No | ACMA |
| **Gibraltar** | 83, 84, 85 | 30 | No | Gibraltar Gambling Commission |

### Integration Example

```javascript
// Region constants for React app
export const REGIONS = {
  MALTA: {
    name: 'Malta',
    countryIds: [134, 135, 136],
    code: 'MT',
    currency: 'EUR'
  },
  USA: {
    name: 'United States',
    countryIds: [231, 232, 233],
    code: 'US',
    currency: 'USD'
  },
  AUSTRALIA: {
    name: 'Australia',
    countryIds: [13, 14, 15],
    code: 'AU',
    currency: 'AUD'
  },
  GIBRALTAR: {
    name: 'Gibraltar',
    countryIds: [83, 84, 85],
    code: 'GI',
    currency: 'GBP'
  }
};

// Filter users by region
const getMaltaUsers = (users) => 
  users.filter(u => REGIONS.MALTA.countryIds.includes(u.country_id));

// Check if user is in specific region
const isUserInUSA = (user) => 
  REGIONS.USA.countryIds.includes(user.country_id);
```

## Testing with cURL

### Get all users
```bash
curl http://localhost:3001/api/v1/users
```

### Get specific user
```bash
curl http://localhost:3001/api/v1/users/u001-malta-134-001
```

### Create new user
```bash
curl -X POST http://localhost:3001/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 100121,
    "user_name": "Test User",
    "country_id": 134,
    "ecdd_status": 1,
    "ecdd_threshold": 10000.00,
    "ecdd_multiplier": 1.00,
    "ecdd_multiplier_rg_flag": false,
    "updated_by": "test@company.com"
  }'
```

### Update user status
```bash
curl -X PUT http://localhost:3001/api/v1/users/u001-malta-134-001 \
  -H "Content-Type: application/json" \
  -d '{
    "ecdd_status": 2,
    "updated_by": "admin@company.com"
  }'
```

## React Integration Example

```javascript
// Fetch all users
const fetchUsers = async () => {
  const response = await fetch('http://localhost:3001/api/v1/users');
  return await response.json();
};

// Create new threshold
const createThreshold = async (thresholdData) => {
  const response = await fetch('http://localhost:3001/api/v1/thresholds', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(thresholdData),
  });
  return response.json();
};

// Update user status
const updateUserStatus = async (userId, updates) => {
  const response = await fetch(`http://localhost:3001/api/v1/users/${userId}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(updates),
  });
  return response.json();
};
```

## Error Handling

All endpoints return appropriate HTTP status codes:

- `200 OK` - Successful GET/PUT
- `201 Created` - Successful POST
- `400 Bad Request` - Invalid request body
- `404 Not Found` - Resource not found
- `405 Method Not Allowed` - Invalid HTTP method

Error responses follow this format:
```json
{
  "error": "Error message description"
}
```

## Development

### Run the server
```bash
go run main.go
```

### Regenerate mock data
```bash
node data/scripts/generate_users.js
```

### Build for production
```bash
go build -o ecdd-mock-api
./ecdd-mock-api
```

## License

MIT License