package model

// PaymentResult is the internal outcome of a processed payment.
type PaymentResult struct {
	TransactionID string
	Status        string
	Message       string
}
