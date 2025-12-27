package service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/yeencloud/lib-base/transaction"
	databaseDomain "github.com/yeencloud/lib-database/domain"
)

func handleWithTransaction[T any](ctx context.Context, trxItf transaction.TransactionInterface, handler func(ctx context.Context) (*T, error)) (*T, error) {
	if trxItf == nil {
		trxItf = transaction.NoTransaction{}
	}

	trx := trxItf.Begin()
	log.Info("Begin transaction")

	ctx = context.WithValue(ctx, databaseDomain.DatabaseCtxKey, trx)

	t, err := handler(ctx)
	if err != nil {
		log.WithError(err).Error("Rollback transaction")
		trx.Rollback()
		return nil, err
	}

	log.Info("Commit transaction")
	trx.Commit()
	return t, nil
}
