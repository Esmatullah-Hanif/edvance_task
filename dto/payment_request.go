package dto

// PaymentRequest is the incoming body for POST /api/process-payment.
//
//	@Description	Payment request payload
type PaymentRequest struct {
	// Amount to charge; must be greater than 0
	Amount float64 `json:"amount" binding:"required,gt=0"`
	// Opaque payment token from the client; never stored or logged in full
	PaymentToken string `json:"paymentToken" binding:"required"`
}
