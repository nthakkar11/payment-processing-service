package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/nthakkar11/payment-processing-service/internal/models"
)

type CustomerService struct {
	DB *sql.DB
}

func NewCustomerService(db *sql.DB) *CustomerService {
	return &CustomerService{DB: db}
}

func (s *CustomerService) Create(ctx context.Context, email string) (*models.Customer, error) {
	customer := &models.Customer{
		ID:        uuid.NewString(),
		Email:     email,
		CreatedAt: time.Now(),
	}

	query := `
		INSERT INTO customers (id, email, created_at)
		VALUES ($1, $2, $3)
	`

	_, err := s.DB.ExecContext(ctx, query, customer.ID, customer.Email, customer.CreatedAt)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
