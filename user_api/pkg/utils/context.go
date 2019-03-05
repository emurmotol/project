package utils

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"
)

type contextKey string

const (
	DBContextKey contextKey = "db"
)

func GetDB(ctx context.Context) *gorm.DB {
	return ctx.Value(DBContextKey).(*gorm.DB)
}

func MustGetDB(ctx context.Context) (*gorm.DB, error) {
	db, ok := ctx.Value(DBContextKey).(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to get db from context")
	}
	return db, nil
}
