package errors

import "errors"

var (
	ErrInvalidAmount       = errors.New("amount must be greater than 0")
	ErrMissingPaymentToken = errors.New("paymentToken is required")
	ErrPaymentProcessing   = errors.New("payment processing failed")
)

// AppError carries an HTTP status code alongside a user-facing message.
type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func New(code int, message string, err error) *AppError {
	return &AppError{Code: code, Message: message, Err: err}
}
