package service

import (
	"fmt"
	"net/http"

	"payment-service/dto"
	apperrors "payment-service/errors"
	"payment-service/model"

	"github.com/google/uuid"
)

// PaymentService defines the contract for payment processing.
type PaymentService interface {
	ProcessPayment(req *dto.PaymentRequest) (*model.PaymentResult, *apperrors.AppError)
}

type paymentService struct{}

// NewPaymentService returns a new PaymentService implementation.
func NewPaymentService() PaymentService {
	return &paymentService{}
}

func (s *paymentService) ProcessPayment(req *dto.PaymentRequest) (*model.PaymentResult, *apperrors.AppError) {
	if req.Amount <= 0 {
		return nil, apperrors.New(http.StatusBadRequest, "Invalid payment request", apperrors.ErrInvalidAmount)
	}
	if req.PaymentToken == "" {
		return nil, apperrors.New(http.StatusBadRequest, "Invalid payment request", apperrors.ErrMissingPaymentToken)
	}

	// Mask token in any internal representations: keep only first 4 chars.
	_ = maskToken(req.PaymentToken)

	txnID := fmt.Sprintf("txn_%s", uuid.New().String())

	return &model.PaymentResult{
		TransactionID: txnID,
		Status:        "success",
		Message:       "Payment processed successfully",
	}, nil
}

// maskToken returns a safe, loggable version of the token.
func maskToken(token string) string {
	if len(token) <= 4 {
		return "****"
	}
	return token[:4] + "****"
}
