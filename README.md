# Lab Inventory Backend

> [!WARNING]
> This project has been replaced by [a newer version](https://github.com/jjacobsohn-eppendorf/lit-lab-inventory-tracker), but is being kept for reference.

Backend project for lab inventory management system. 

**Really just naive first few steps in Go, so don't expect too much - I am learning as I _go_ (pun intended).**

## Installation

### Locally
> You need to have go installed and set up
- `go mod download` - Install necessary go modules
- `go run ./internal/main.go` - Start server on port 8080

### Docker
- `docker build -t lab-inventory-backend:latest .`
- `docker run -p 8080:8080 --name lab-inventory-backend-container lab-inventory-backend:latest`

## Environment variables

| Key  | Required | Default | Description                      |
|------|----------|---------|----------------------------------|
| PORT | no       | 8081    | Port on which the server listens |

## API endpoints (GET)

- `/devices/` - list devices
- `/devices/:id` - list one device by id
