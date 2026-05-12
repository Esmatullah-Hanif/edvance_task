# Payment Processing Microservice

A simulated fintech payment processing API built with **Go** and **Gin**, following clean microservices-style architecture.

---

## Prerequisites

| Tool     | Version |
|----------|---------|
| Go       | 1.21+   |
| swag CLI | latest  |

Install the `swag` CLI (needed to regenerate Swagger docs):

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

---

## Project Structure

```
.
├── config/          # Environment config loader
├── docs/            # Auto-generated Swagger docs (committed for convenience)
├── dto/             # Request data transfer objects
├── errors/          # Centralized custom error types and sentinels
├── handler/         # HTTP handlers (parse request → call service → return response)
├── middleware/      # Logger and recovery middleware
├── model/           # Internal domain models
├── response/        # Response helpers and Swagger type definitions
├── routes/          # Route registration
├── service/         # Business logic
├── main.go          # Entry point with Swagger general annotations
└── go.mod
```

---

## Running the Service

```bash
# 1. Download dependencies
go mod download

# 2. Start the server (default port: 8080)
go run main.go
```

To use a custom port or environment:

```bash
PORT=9090 APP_ENV=production go run main.go
```

---

## Regenerating Swagger Docs

Run this command from the project root whenever you change Swagger annotations:

```bash
swag init
```

This overwrites `docs/docs.go`, `docs/swagger.json`, and `docs/swagger.yaml`.

---

## Swagger UI

Once the server is running, open:

```
http://localhost:8080/swagger/index.html
```

---

## API Reference

### POST /api/process-payment

Simulate a payment charge.

**Request body:**

```json
{
  "amount": 100.50,
  "paymentToken": "tok_test_123"
}
```

**Success response (200):**

```json
{
  "status": "success",
  "transactionId": "txn_550e8400-e29b-41d4-a716-446655440000",
  "message": "Payment processed successfully"
}
```

**Validation error response (400):**

```json
{
  "success": false,
  "message": "Invalid payment request",
  "error": "amount must be greater than 0"
}
```

**Example curl:**

```bash
curl -X POST http://localhost:8080/api/process-payment \
  -H "Content-Type: application/json" \
  -d '{"amount": 100.50, "paymentToken": "tok_test_123"}'
```

---

### GET /health

Returns the service health status.

**Response (200):**

```json
{
  "status": "ok",
  "service": "payment-service"
}
```

**Example curl:**

```bash
curl http://localhost:8080/health
```

---

## Validation Rules

| Field          | Rule                             |
|----------------|----------------------------------|
| `amount`       | Required, must be greater than 0 |
| `paymentToken` | Required, must not be empty      |

The full `paymentToken` value is **never stored or logged**. Only a masked version (`tok_****`) is used internally.

---

## Environment Variables

| Variable  | Default       | Description                           |
|-----------|---------------|---------------------------------------|
| `PORT`    | `8080`        | Port the HTTP server listens on       |
| `APP_ENV` | `development` | Set to `production` for release mode  |
