package repository

import (
	"belajar-rest-api/helper"
	"belajar-rest-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepoImpl struct {
}

func NewCategoryRepo() CategoryRepoInterfaces {
	return &CategoryRepoImpl{}
}

func (repository CategoryRepoImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "INSERT INTO category(name) VALUES (?)"

	result, err := tx.ExecContext(ctx, sql, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository CategoryRepoImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "UPDATE category set name = ? WHERE id = ?"

	_, err := tx.ExecContext(ctx, sql, &category.Name, &category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository CategoryRepoImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	sql := "DELETE FROM category WHERE id = ?"

	_, err := tx.ExecContext(ctx, sql, category.Id)
	helper.PanicIfError(err)
}

func (repository CategoryRepoImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	sql := "SELECT id, name FROM category WHERE id = ?"

	row, err := tx.QueryContext(ctx, sql, categoryId)
	helper.PanicIfError(err)
	defer row.Close()

	category := domain.Category{}
	if row.Next() {
		err := row.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil

	} else {
		return category, errors.New("category tidak ditemukan")
	}
}

func (repository CategoryRepoImpl) FIndAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sql := "SELECT id, name FROM category"

	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}

		err := rows.Scan(&category.Id, category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}
