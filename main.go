// Package main is the entry point for the payment processing microservice.
//
//	@title			Payment Processing API
//	@version		1.0
//	@description	A simulated fintech payment processing microservice built with Go and Gin.
//	@termsOfService	http://example.com/terms/
//
//	@contact.name	API Support
//	@contact.email	support@example.com
//
//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT
//
//	@host		localhost:8080
//	@BasePath	/
//
//	@schemes	http https
package main

import (
	"log"

	_ "payment-service/docs"
	"payment-service/config"
	"payment-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	routes.Setup(r)

	log.Printf("Starting payment-service on port %s (env: %s)", cfg.Port, cfg.Env)
	log.Printf("Swagger UI available at http://localhost:%s/swagger/index.html", cfg.Port)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
