# ISA Investment System

## Overview

The ISA Investment System is a backend service that enables retail customers to invest in Individual Savings Accounts (ISAs) by selecting from a predefined list of funds. Built with Go using the Echo framework and PostgreSQL as the database, this system ensures efficient handling of investments with a focus on scalability and maintainability.

## Features

- **Fund Management**: Load and serve a list of available funds from a JSON file.
- **Investment Processing**: Allow customers to invest in one or multiple funds in a single transaction.
- **Investment Retrieval**: Fetch all investments made by a specific customer.
- **Health Monitoring**: Provide a health check endpoint to monitor system status.

## Tech Stack

- **Language**: Go
- **Framework**: Echo
- **Database**: PostgreSQL
- **ORM**: go-pg/pg v10
- **Migrations**: golang-migrate
- **Logging**: logrus
- **API Documentation**: swaggo/echo-swagger
- **Containerization**: Docker & Docker Compose

## Prerequisites

- Go 1.23 or higher
- Docker & Docker Compose

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/warrenb95/isa-investment-system.git
cd isa-investment-system
```

### 2. Start server with default postgres DB

```bash
docker-compose up --build
```

## API Endpoints

### 1. Get All Funds

**Endpoint**: `GET /api/v1/retail/funds`

**Description**: Retrieve a list of all available funds.

**Response**:

```json
[
  {
    "id": "uuid",
    "name": "Fund Name",
    "description": "Fund Description"
  },
  ...
]
```

### 2. Create Investments

**Endpoint**: `POST /api/v1/retail/isa-investments/:customer_id`

**Description**: Create a new investment for a customer.

**Request Parameters**:

- `customer_id` (UUID): The ID of the customer.

**Request Body**:

```json
{
  "investments": [
    {
      "fund_id": "uuid",
      "amount": 1000.00
    },
    ...
  ]
}
```

**Response**:

```json
{
  "message": "Investments successfully created."
}
```

### 3. Get Customer Investments

**Endpoint**: `GET /api/v1/retail/isa-investments/:customer_id`

**Description**: Retrieve all investments for a specific customer.

**Response**:

```json
[
  {
    "fund_id": "uuid",
    "amount": 1000.00,
    "created_at": "2025-02-21T12:00:00Z"
  },
  ...
]
```

### 4. Health Check

**Endpoint**: `GET /api/v1/health`

**Description**: Check the health status of the application.

**Response**:

```json
{
  "status": "healthy"
}
```

## Testing

### Running Tests

To execute unit and integration tests:

```bash
go test ./...
```

*Note: Ensure that the PostgreSQL database is running and accessible before running integration tests.*
