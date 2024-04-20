package service

import (
	"context"

	"github.com/arganaphang/money-manager/internal/dto"
	"github.com/arganaphang/money-manager/internal/model"
	"github.com/arganaphang/money-manager/internal/repository"
	"github.com/google/uuid"
)

type ITransactionService interface {
	CreateTransaction(ctx context.Context, data dto.CreateTransactionRequest) error
	GetTransactions(ctx context.Context) ([]model.Transaction, error)
	GetTransactionByID(ctx context.Context, id uuid.UUID) (*model.Transaction, error)
	UpdateTransactionByID(ctx context.Context, id uuid.UUID, data dto.UpdateTransactionRequest) error
	DeleteTransactionByID(ctx context.Context, id uuid.UUID) error
}

type TransactionService struct {
	Repositories repository.Repositories
}

func NewTransactionService(repositories repository.Repositories) ITransactionService {
	return &TransactionService{Repositories: repositories}
}

func (s TransactionService) CreateTransaction(ctx context.Context, data dto.CreateTransactionRequest) error {
	return s.Repositories.TransactionRepository.CreateTransaction(ctx, data)
}

func (s TransactionService) GetTransactions(ctx context.Context) ([]model.Transaction, error) {
	return s.Repositories.TransactionRepository.GetTransactions(ctx)
}

func (s TransactionService) GetTransactionByID(ctx context.Context, id uuid.UUID) (*model.Transaction, error) {
	return s.Repositories.TransactionRepository.GetTransactionByID(ctx, id)
}

func (s TransactionService) UpdateTransactionByID(ctx context.Context, id uuid.UUID, data dto.UpdateTransactionRequest) error {
	return s.Repositories.TransactionRepository.UpdateTransactionByID(ctx, id, data)
}

func (s TransactionService) DeleteTransactionByID(ctx context.Context, id uuid.UUID) error {
	return s.Repositories.TransactionRepository.DeleteTransactionByID(ctx, id)
}
