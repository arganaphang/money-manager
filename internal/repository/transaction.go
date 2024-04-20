package repository

import (
	"context"
	"errors"
	"time"

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
	DB []model.Transaction
}

func NewTransactionRepository( /* DB CONN */ ) ITransactionRepository {
	return &TransactionRepository{
		// DB CONN
		DB: []model.Transaction{},
	}
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, data dto.CreateTransactionRequest) error {
	now := time.Now()
	r.DB = append(r.DB, model.Transaction{
		ID:        uuid.New(),
		Title:     data.Title,
		Note:      data.Note,
		Amount:    data.Amount,
		Type:      data.Type,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	})
	return nil
}

func (r TransactionRepository) GetTransactions(ctx context.Context) ([]model.Transaction, error) {
	return r.DB, nil
}

func (r TransactionRepository) GetTransactionByID(ctx context.Context, id uuid.UUID) (*model.Transaction, error) {
	var result *model.Transaction // nil
	for _, item := range r.DB {
		item := item
		if item.ID == id {
			result = &item
			break
		}
	}
	if result == nil {
		return nil, errors.New("not found")
	}
	return &model.Transaction{}, nil
}

func (r TransactionRepository) UpdateTransactionByID(ctx context.Context, id uuid.UUID, data dto.UpdateTransactionRequest) error {
	return nil
}

func (r TransactionRepository) DeleteTransactionByID(ctx context.Context, id uuid.UUID) error {
	return nil
}
