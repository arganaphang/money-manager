package repository

import (
	"context"
	"time"

	"github.com/arganaphang/money-manager/internal/dto"
	"github.com/arganaphang/money-manager/internal/model"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ITransactionRepository interface {
	CreateTransaction(ctx context.Context, data dto.CreateTransactionRequest) error
	GetTransactions(ctx context.Context, req dto.GetTransactionsRequest) ([]model.Transaction, error)
	GetTransactionByID(ctx context.Context, id uuid.UUID) (*model.Transaction, error)
	UpdateTransactionByID(ctx context.Context, id uuid.UUID, data dto.UpdateTransactionByIDRequest) error
	DeleteTransactionByID(ctx context.Context, id uuid.UUID) error
}

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) ITransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, data dto.CreateTransactionRequest) error {
	query, _, err := goqu.
		Insert("transactions").
		Rows(goqu.Record{
			"title":  data.Title,
			"note":   data.Note,
			"amount": data.Amount,
			"type":   data.Type,
		}).
		ToSQL()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r TransactionRepository) GetTransactions(ctx context.Context, req dto.GetTransactionsRequest) ([]model.Transaction, error) {
	limitOffset := req.TransformToLimitOffset()
	sql := goqu.From("transactions")
	query, _, err := sql.
		Where(goqu.Ex{
			"deleted_at": nil,
		}).
		Order(goqu.I("created_at").Desc()).
		Limit(limitOffset.Limit).
		Offset(limitOffset.Offset).
		ToSQL()
	if err != nil {
		return nil, err
	}

	var transactions []model.Transaction
	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var transaction model.Transaction
		err = rows.StructScan(&transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (r TransactionRepository) GetTransactionByID(ctx context.Context, id uuid.UUID) (*model.Transaction, error) {
	query, _, err := goqu.
		From("transactions").
		Where(goqu.Ex{
			"id":         id,
			"deleted_at": nil,
		}).
		Limit(1).
		ToSQL()
	if err != nil {
		return nil, err
	}

	var transaction model.Transaction
	err = r.db.QueryRowx(query).StructScan(&transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r TransactionRepository) UpdateTransactionByID(ctx context.Context, id uuid.UUID, data dto.UpdateTransactionByIDRequest) error {
	// TODO: Update Record
	now := time.Now()
	query, _, err := goqu.
		Update("transactions").
		Set(goqu.Record{
			"updated_at": now,
		}).
		Where(goqu.Ex{"id": id}).
		ToSQL()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r TransactionRepository) DeleteTransactionByID(ctx context.Context, id uuid.UUID) error {
	now := time.Now()
	query, _, err := goqu.
		Update("notifications").
		Set(goqu.Record{"deleted_at": now}).
		Where(goqu.Ex{"id": id}).
		ToSQL()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
