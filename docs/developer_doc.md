# Developer Documentation

## Overview
This document provides an overview of the **Test** module in the Go Blog API, detailing the structure, functionality, and how to interact with its components. The module is designed to manage tags with features such as creating, updating, and retrieving tags with pagination.

---

## Module Structure

### Interfaces
Defines contracts for the repository layer.

```go
type TestRepositoryInterface interface {
    CreateTest(data *models.Test) (models.Test, error)
    UpdateTest(id int, data *models.Test) (models.Test, error)
    FindWithPagination(ctx *gin.Context) ([]models.Test, error)
}
```

### Repository
Implements the data access layer, interfacing with the database using GORM.

```go
type TestRepository struct {
    db *gorm.DB
}

func (repo *TestRepository) CreateTest(data *models.Test) (models.Test, error) { /* ... */ }
func (repo *TestRepository) UpdateTest(id int, data *models.Test) (models.Test, error) { /* ... */ }
func (repo *TestRepository) FindWithPagination(ctx *gin.Context) ([]models.Test, error) { /* ... */ }
```

#### Initialization
```go
func NewTestRepository(db *gorm.DB) interfaces.TestRepositoryInterface {
    return &TestRepository{db: db}
}
```

### Services
Encapsulates business logic and delegates data operations to the repository.

```go
type TestService struct {
    repo interfaces.TestRepositoryInterface
}

func (service *TestService) CreateTest(data *models.Test) (models.Test, error) { /* ... */ }
func (service *TestService) UpdateTest(id int, data *models.Test) (models.Test, error) { /* ... */ }
func (service *TestService) FindWithPagination(ctx *gin.Context) ([]models.Test, error) { /* ... */ }
```

#### Initialization
```go
func NewTestService(repo interfaces.TestRepositoryInterface) *TestService {
    return &TestService{repo: repo}
}
```

### Handlers
Handles HTTP requests and responses for the Test module using `gin-gonic`.

```go
type TestHandler struct {
    tagService *services.TestService
}

func (handler *TestHandler) CreateTest(context *gin.Context) { /* ... */ }
func (handler *TestHandler) UpdateTest(context *gin.Context) { /* ... */ }
func (handler *TestHandler) FindWithPagination(context *gin.Context) { /* ... */ }
```

#### Initialization
```go
func NewTestHandler(tagService *services.TestService) *TestHandler {
    return &TestHandler{tagService: tagService}
}
```

### Routes
Registers the routes for the Test module.

```go
func (s *Server) TestRoutes(router *gin.RouterGroup, deps *Dependencies) {
    tagRouter := router.Group("/tags")
    tagRouter.POST("/", authMiddleware.AuthMiddleware(), deps.TestHandler.CreateTest)
}
```

---

## API Endpoints

### Create Test
- **Endpoint:** `POST /api/v1/tags`
- **Description:** Create a new tag.
- **Authentication:** Required.

**Request Body:**
```json
{
    "name": "TestName",
    "userId": 1,
    "parentId": 0,
    "iconId": 123
}
```

**Response:**
```json
{
    "message": "Success",
    "data": {
        "id": 1,
        "name": "TestName",
        "userId": 1,
        "parentId": 0,
        "iconId": 123
    }
}
```

### Update Test
- **Endpoint:** `PUT /api/v1/tags/:id`
- **Description:** Update a tag by its ID.
- **Authentication:** Required.

**Request Body:**
```json
{
    "name": "UpdatedTestName",
    "userId": 1,
    "parentId": 0,
    "iconId": 456
}
```

**Response:**
```json
{
    "message": "Success",
    "data": {
        "id": 1,
        "name": "UpdatedTestName",
        "userId": 1,
        "parentId": 0,
        "iconId": 456
    }
}
```

### Find Tests with Pagination
- **Endpoint:** `GET /api/v1/tags`
- **Description:** Retrieve tags with pagination.
- **Authentication:** Required.

**Query Parameters:**
- `page`: Page number (default: 1)
- `limit`: Number of tags per page (default: 10)

**Response:**
```json
{
    "message": "Success",
    "data": [
        { "id": 1, "name": "Test1" },
        { "id": 2, "name": "Test2" }
    ]
}
```

---

## Dependencies

### Database
- Uses GORM for database interactions. The database connection is initialized via `database.Connect()`.

### Validation
- Validates request data using the `validator.ValidateStruct` utility.

### Utilities
- **Error Handling:** `utils.ErrorResponse`
- **Success Response:** `utils.SuccessResponse`

---

## Setting Up

### Database Connection
Ensure a valid database connection is established via `database.Connect()`.

### Dependencies Initialization
Initialize all module dependencies using `NewAppDependencies()`.

### Route Registration
Register the Test routes via `s.TestRoutes()`.

---

## Testing
- Use tools like Postman or `curl` to test the API endpoints.
- Ensure authentication headers are included where required.
