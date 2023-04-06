package repository

import (
	"context"
	"database/sql"

	"github.com/sesaquecruz/goexpert-client-server-api-lab/internal/model"
)

type QuoteRepository struct {
	db *sql.DB
}

func NewQuoteRepository(db *sql.DB) *QuoteRepository {
	return &QuoteRepository{
		db: db,
	}
}

func (q *QuoteRepository) SaveQuote(ctx context.Context, quote *model.Quote) error {
	stmt, err := q.db.Prepare("insert into quotes (bid, timestamp) values (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, quote.Bid, quote.Timestamp)
	return err
}
