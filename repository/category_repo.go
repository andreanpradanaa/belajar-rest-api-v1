package repository

import (
	"belajar-rest-api/model/domain"
	"context"
	"database/sql"
)

type CategoryRepoInterfaces interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FIndAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
