package handler

import (
	"net/http"

	"payment-service/dto"
	"payment-service/response"
	"payment-service/service"

	"github.com/gin-gonic/gin"
)

// PaymentHandler handles payment-related HTTP requests.
type PaymentHandler struct {
	svc service.PaymentService
}

// NewPaymentHandler constructs a PaymentHandler.
func NewPaymentHandler(svc service.PaymentService) *PaymentHandler {
	return &PaymentHandler{svc: svc}
}

// ProcessPayment godoc
//
//	@Summary		Process a payment
//	@Description	Simulate secure payment processing. Validates the request and returns a transaction ID on success.
//	@Tags			payments
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.PaymentRequest		true	"Payment request payload"
//	@Success		200		{object}	response.PaymentSuccess	"Payment processed successfully"
//	@Failure		400		{object}	response.ErrorResponse	"Validation or business rule failure"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/process-payment [post]
func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var req dto.PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid payment request", err.Error())
		return
	}

	result, appErr := h.svc.ProcessPayment(&req)
	if appErr != nil {
		response.Error(c, appErr.Code, appErr.Message, appErr.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        result.Status,
		"transactionId": result.TransactionID,
		"message":       result.Message,
	})
}
