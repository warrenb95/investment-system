# ISA Investment System - Implementation Steps

## Step 1: Initialize Project

- [x] Initialize Golang module (`go mod init isa-investment`)
- [x] Install Echo (`go get github.com/labstack/echo/v4`)
- [x] Set up basic HTTP server
- [x] Add graceful shutdown handling
- [x] Install `logrus` and configure JSON logging

## Step 2: Database Setup

- [x] Install `go-pg/pg v10` for PostgreSQL ORM
- [x] Install `golang-migrate`
- [x] Define database schema for `customers`, `funds`, `isa_investments`
- [x] Implement database connection function
- [x] Write integration tests for schema validation

## Step 3: Load and Serve Funds

- [x] Load `funds.json` into memory at startup
- [x] Implement `GET /api/v1/funds` to return funds
- [x] Write tests for fund retrieval

## Step 4: Create Investments

- [ ] Define investment request struct with validation
- [ ] Implement `POST /api/v1/isa-investments/:customer_id`
- [ ] Use transactions to ensure atomicity
- [ ] Implement per-customer locking (`FOR UPDATE`)
- [ ] Write unit and integration tests

## Step 5: Retrieve Investments

- [ ] Implement `GET /api/v1/isa-investments/:customer_id`
- [ ] Include timestamps and total invested amount
- [ ] Write tests for investment retrieval

## Step 6: Error Handling and Validation

- [ ] Implement structured error responses
- [ ] Validate request payloads (min 1 fund, positive amounts, no duplicates)
- [ ] Write tests for validation failures

## Step 7: Health Check Endpoint

- [ ] Implement `GET /api/v1/health`
- [ ] Write tests for health check

## Step 8: CI/CD and Testing

- [ ] Set up `golangci-lint` for linting
- [ ] Implement GitHub Actions pipeline for testing
- [ ] Add unit tests for validation logic
- [ ] Add integration tests for API endpoints

## Step 9: Deployment and Dockerization

- [ ] Create a Dockerfile
- [ ] Set up `docker-compose.yml` with PostgreSQL
- [ ] Ensure database migrations run at startup
- [ ] Test full deployment
