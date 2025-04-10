package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/atjulia/payment-gateway/gateway-api/internal/repository"
	"github.com/atjulia/payment-gateway/gateway-api/internal/service"
	"github.com/atjulia/payment-gateway/gateway-api/internal/web/server"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "1234"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "3307"),
		getEnv("DB_NAME", "gateway"),
	)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	invoiceRepository := repository.NewInvoiceRepository(db)
	invoiceService := service.NewInvoiceService(invoiceRepository, *accountService)

	port := getEnv("HTTP_PORT", "8080")
	srv := server.NewServer(accountService, invoiceService, port)
	srv.ConfigureRoutes()

	if err := srv.Start(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
