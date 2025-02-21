# ISA Investment System - Developer Specification

## Overview

This document provides a detailed specification for implementing an ISA investment system using Golang (Echo framework) and PostgreSQL. The system allows retail customers to invest in a selection of preloaded funds, tracking their investments while ensuring transactional consistency.

## Tech Stack

- **Backend Framework**: Golang (Echo)
- **Database**: PostgreSQL
- **ORM**: `go-pg/pg v10`
- **Migrations**: `golang-migrate`
- **Logging**: `sirupsen/logrus` (JSON format, includes timestamps)
- **Testing**: `testing` package, `dockertest` for integration tests
- **API Documentation**: `swaggo/echo-swagger`
- **Deployment**: Dockerized setup with `docker-compose`

## Functional Requirements

### Core Features

- Customers can invest in one or multiple funds in a single request.
- Investments are stored in PostgreSQL and updated if a customer reinvests in the same fund (summing the amounts instead of replacing).
- Customers can retrieve all their past investments.
- Funds are preloaded from a `funds.json` file at application startup.
- API requests and responses follow a structured format, including error codes.
- Graceful shutdown ensures database connections are cleaned up.

### API Endpoints

#### 1. **Funds Management**

- `GET /api/v1/funds` - Returns all available funds loaded from `funds.json`.

#### 2. **ISA Investments**

- `POST /api/v1/isa-investments/:customer_id`
  - Allows investing in multiple funds at once.
  - Updates existing investments by summing amounts if the same fund is reinvested in.
  - Requires at least one fund selection.
- `GET /api/v1/isa-investments/:customer_id`
  - Returns all investments for a given customer, including timestamps and total invested.

#### 3. **Health Check**

- `GET /api/v1/health` - Checks database connectivity and app readiness.

## Database Schema

### Customers Table

```sql
CREATE TABLE customers (
    id UUID PRIMARY KEY
);
```

### Funds Table

```sql
CREATE TABLE funds (
    id UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL
);
```

### ISA Investments Table

```sql
CREATE TABLE isa_investments (
    id UUID PRIMARY KEY,
    customer_id UUID REFERENCES customers(id) ON DELETE CASCADE,
    fund_id UUID REFERENCES funds(id) ON DELETE CASCADE,
    amount DECIMAL NOT NULL CHECK (amount > 0),
    created_at TIMESTAMP DEFAULT NOW()
);
```

## Data Handling

### **Transactions & Concurrency**

- Investments are handled within a **database transaction** to ensure atomicity.
- Customer investments are **locked per customer (`FOR UPDATE`)** to prevent race conditions.
- Investments are updated instead of duplicated when reinvesting in the same fund.

### **Logging Strategy**

- All API requests are logged.
- Errors are logged at `ERROR` level.
- Sensitive request data (e.g., request body) is omitted from logs.

## Error Handling & Validation

### **General Error Format**

```json
{
    "error": "Invalid request",
    "code": 1001,
    "details": "Amount must be greater than zero"
}
```

### **Validation Rules**

- At least **one fund must be provided** in investment requests.
- **Duplicate funds** in the same request trigger a `400 Bad Request`.
- **Amounts must be positive**.
- **Invalid fund IDs** result in a `400 Bad Request`.

## Testing Plan

### **Unit Tests**

- Table-driven tests for validation and business logic.
- Ensures investment calculations and constraints work as expected.

### **Integration Tests**

- Uses `dockertest` to spin up PostgreSQL for tests.
- Truncates test data **before and after** test execution.

## CI/CD Pipeline (GitHub Actions)

### **Automated Checks**

- **Linting** using `golangci-lint`.
- **Unit tests** using `go test ./...`.
- **Does not** run `dockertest` in CI (not needed at this stage).

## Deployment & Future Considerations

### **Current Deployment Strategy**

- Uses a **Dockerfile** for containerization.
- `docker-compose.yml` for local development with PostgreSQL.

### **Future Enhancements (Not Implemented Now)**

- **Authentication & Authorization** (JWT, API keys).
- **Feature Flags** for enabling/disabling features dynamically.
- **Rate Limiting** for abuse prevention.
- **Background Jobs** for batch processing.
- **External API Integrations** (market data, webhooks).
- **Performance Monitoring** (Prometheus, OpenTelemetry).

## Summary

This specification provides a **minimal but scalable** foundation for an ISA investment system. The architecture prioritizes **simplicity, data consistency, and developer productivity**, with future extensibility built-in.
