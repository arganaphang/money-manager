package repository

import (
	"context"

	"github.com/arganaphang/money-manager/internal/dto"
	"github.com/arganaphang/money-manager/internal/model"
	"github.com/google/uuid"
)

type ITransactionRepository interface {
	CreateTransaction(ctx context.Context, data dto.CreateTransactionRequest) error
	GetTransactions(ctx context.Context) ([]model.Transaction, error)
	GetTransactionByID(ctx context.Context, id uuid.UUID) (*model.Transaction, error)
	UpdateTransactionByID(ctx context.Context, id uuid.UUID, data dto.UpdateTransactionRequest) error
	DeleteTransactionByID(ctx context.Context, id uuid.UUID) error
}

type TransactionRepository struct {
	// DB CONN
}

func NewTransactionRepository( /* DB CONN */ ) ITransactionRepository {
	return &TransactionRepository{
		// DB CONN
	}
}

func (r TransactionRepository) CreateTransaction(ctx context.Context, data dto.CreateTransactionRequest) error {
	return nil
}

func (r TransactionRepository) GetTransactions(ctx context.Context) ([]model.Transaction, error) {
	return []model.Transaction{}, nil
}

func (r TransactionRepository) GetTransactionByID(ctx context.Context, id uuid.UUID) (*model.Transaction, error) {
	return &model.Transaction{}, nil
}

func (r TransactionRepository) UpdateTransactionByID(ctx context.Context, id uuid.UUID, data dto.UpdateTransactionRequest) error {
	return nil
}

func (r TransactionRepository) DeleteTransactionByID(ctx context.Context, id uuid.UUID) error {
	return nil
}
