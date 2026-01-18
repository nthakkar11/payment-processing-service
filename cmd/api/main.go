package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nthakkar11/payment-processing-service/internal/db"
	"github.com/nthakkar11/payment-processing-service/internal/handlers"
	"github.com/nthakkar11/payment-processing-service/internal/services"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/payments?sslmode=disable"
	}

	database := db.New(dsn)

	customerService := services.NewCustomerService(database)
	customerHandler := handlers.NewCustomerHandler(customerService)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/v1/customers", customerHandler.Create)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
