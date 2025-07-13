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

---

## Project Structure

```
Cred/
  ├── pkg/dao/                      # Data access layer (DAO, mocks, errors)
  ├── pkg/models/                   # Domain models
  ├── pkg/service/                  # Service layer and business logic
  ├── go.mod                        # Go module definition
  └── README.md                     # Project documentation
```

---

## Service Layer Overview

### UserService
- **Responsibilities:**
  - Handles uploading contacts (validates and stores phone books for a user)
  - Looks up user details (name, spam status) by phone number
- **Key Methods:**
  - `UploadContacts(ctx, ownerPhoneNumber, contacts)`
  - `LookupUser(ctx, phoneNumber)`
- **Business Logic:**
  - Validates phone numbers and contact data
  - Associates contacts with the uploader's phone number
  - Returns the most recent name and spam status for a number

### SpamService
- **Responsibilities:**
  - Periodically updates the spam status for all users (simulates a nightly job)
- **Key Method:**
  - `UpdateSpamStatus(ctx)`
- **Business Logic:**
  - Fetches all users and updates their spam status based on internal rules or a simulated data science model
  - Designed for batch/background operation, not user-triggered

---

## Getting Started

### Prerequisites
- Go 1.18+

### Setup
```sh
# Clone the repository
git clone <repo-url>
cd Cred

# Install dependencies
go mod tidy
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

---

## Code Quality & Coverage
- Use `go vet ./...` to check for code issues.
- Use `go test ./... -cover` to check test coverage. Aim for high functional coverage, not just line coverage.
- Keep code and comments clean; remove redundant or outdated comments.

---

## Contact
For questions or contributions, please open an issue or submit a pull request.
