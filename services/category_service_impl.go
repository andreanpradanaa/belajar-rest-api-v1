package services

import (
	"belajar-rest-api/exception"
	"belajar-rest-api/helper"
	"belajar-rest-api/model/domain"
	"belajar-rest-api/model/web"
	"belajar-rest-api/repository"
	"context"
	"database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepo repository.CategoryRepoInterfaces
	DB           *sql.DB
	Validate     *validator.Validate
}

func NewCategoryService(CategoryRepo repository.CategoryRepoInterfaces, DB *sql.DB, Validate *validator.Validate) CategoryServiceInterface {
	return &CategoryServiceImpl{
		CategoryRepo: CategoryRepo,
		DB:           DB,
		Validate:     Validate,
	}
}

func (service CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepo.Save(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepo.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = service.CategoryRepo.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepo.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepo.Delete(ctx, tx, category)
}

func (service CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepo.FindById(ctx, tx, categoryId)
	if err != nil {
		fmt.Println("hello")
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := service.CategoryRepo.FIndAll(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToCategoryResponses(category)
}
