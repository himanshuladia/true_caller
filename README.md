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
  ├── internal/                     # Internal packages (business logic, handlers, etc.)
  ├── pkg/                          # Shared packages (if needed)
  ├── go.mod                        # Go module definition
  └── README.md                     # Project documentation
```

## Getting Started

### Prerequisites
- Go 1.18+

### Setup
```sh
# Clone the repository
# cd into the project directory
# Initialize Go modules (already done)

# Run the application
cd cmd/truecaller-lite
go run main.go
```

---

## Next Steps
- Implement API endpoints and in-memory data structures
- Add phone number validation
- Add nightly job simulation for spam status

---

This README will be updated as the project evolves.
