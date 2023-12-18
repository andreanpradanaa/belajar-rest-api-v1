package helper

import (
	"belajar-rest-api/model/domain"
	"belajar-rest-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {

	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}

}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categorieyResponses []web.CategoryResponse
	for _, category := range categories {
		categories = append(categories, domain.Category(ToCategoryResponse(category)))
	}

	return categorieyResponses
}
