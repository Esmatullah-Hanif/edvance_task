package routes

import (
	"payment-service/handler"
	"payment-service/middleware"
	"payment-service/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Setup registers all routes on the given engine.
func Setup(r *gin.Engine) {
	paymentSvc := service.NewPaymentService()
	paymentHandler := handler.NewPaymentHandler(paymentSvc)

	r.Use(middleware.Recovery())
	r.Use(middleware.Logger())

	r.GET("/health", handler.Health)

	api := r.Group("/api")
	{
		api.POST("/process-payment", paymentHandler.ProcessPayment)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
