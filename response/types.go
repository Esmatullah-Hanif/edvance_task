package response

// PaymentSuccess is the success response shape for POST /api/process-payment.
//
//	@Description	Successful payment response
type PaymentSuccess struct {
	Status        string `json:"status"        example:"success"`
	TransactionID string `json:"transactionId" example:"txn_550e8400-e29b-41d4-a716-446655440000"`
	Message       string `json:"message"       example:"Payment processed successfully"`
}

// ErrorResponse is the standard error envelope.
//
//	@Description	Error response envelope
type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Invalid payment request"`
	Error   string `json:"error"   example:"amount must be greater than 0"`
}

// HealthResponse is returned by GET /health.
//
//	@Description	Health check response
type HealthResponse struct {
	Status  string `json:"status"  example:"ok"`
	Service string `json:"service" example:"payment-service"`
}
