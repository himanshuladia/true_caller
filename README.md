# TrueCaller-Lite

A minimal in-memory TrueCaller-like service in Go.

## Requirements

### 1. Contact Upload (POST API)
- Accepts multiple contacts per request.
- Each contact: phone number (must be 10 digits, start with 91), name.
- Contacts are associated with the uploading phone number (the uploader's phone number is the user identifier).
- The POST request must specify the uploader's phone number.
- Duplicate contacts (same uploader, same contact phone number) are overridden by the latest write.

### 2. Lookup (GET API)
- No authentication required.
- Returns: name (most recent), spam status.
- Spam status is updated nightly by an internal job (not exposed via API).

### 3. Spam Reporting
- No public API for reporting spam; handled internally.

### 4. User Management & Security
- No user management or authentication required.
- The uploader's phone number is used as the unique identifier for associating uploaded contacts.

### 5. Data Privacy
- No privacy requirements for this prototype.

### 6. Technical/Operational
- In-memory storage only (no database).
- Target: 100K RPS.
- No containerization or deployment setup required.

---

## Project Structure

```
Cred/
  ├── cmd/truecaller-lite/main.go   # Application entry point
  ├── pkg/dao/                      # Data access layer (DAO, mocks, errors)
  ├── pkg/models/                   # Domain models
  ├── pkg/service/                  # Service layer and business logic
  ├── go.mod                        # Go module definition
  └── README.md                     # Project documentation
```

## Getting Started

### Prerequisites
- Go 1.18+

### Setup & Running
```sh
# Clone the repository
git clone <repo-url>
cd Cred

# Install dependencies (if any)
go mod tidy

# Run the application
cd cmd/truecaller-lite
go run main.go
```

### Running Tests & Checking Coverage
```sh
# Run all tests
go test ./...

# Run tests with coverage and generate a report
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

---

## Development Guide

### Code Quality
- Use `go vet` and `golint` (if available) to check for code issues.
- Write table-driven tests for all business logic and validation.
- Use context.Context in all service and DAO methods for cancellation/timeouts.
- Use shared error variables and errors.Is for error handling.
- Keep business logic in the service layer, not in handlers or DAOs.

### Contributing
- Fork the repository and create a feature branch.
- Write clear, descriptive commit messages.
- Add/maintain GoDoc comments for all exported types and methods.
- Ensure all tests pass before submitting a PR.
- Run `go mod tidy` to keep dependencies clean.

### Directory Conventions
- `pkg/dao/` - Data access layer, including mocks and error definitions
- `pkg/models/` - Domain models and validation
- `pkg/service/` - Service interfaces and business logic
- `cmd/truecaller-lite/` - Application entry point

---

## Code Quality & Coverage
- Use `go vet ./...` to check for code issues.
- Use `go test ./... -cover` to check test coverage. Aim for high functional coverage, not just line coverage.
- Keep code and comments clean; remove redundant or outdated comments.

---

## Contact
For questions or contributions, please open an issue or submit a pull request.
