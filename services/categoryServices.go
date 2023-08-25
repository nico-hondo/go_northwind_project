package services

import (
	"net/http"

	"codeid.northwind/models"
	"codeid.northwind/models/features"
	"codeid.northwind/repositories"
	"codeid.northwind/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type CategoryService struct {
	//Before
	// categoryRepository *repositories.CategoryRepository

	//after
	repositoryMgr *repositories.RepositoryManager
}

func NewCategoryService(repoMgr *repositories.RepositoryManager) *CategoryService {
	return &CategoryService{
		repositoryMgr: repoMgr,
	}
}

func (cs CategoryService) GetListCategory(ctx *gin.Context, metadata *features.Metadata) ([]*models.Category, *models.ResponseError) {
	// return cs.categoryRepository.GetListCategory(ctx)
	return cs.repositoryMgr.CategoryRepository.GetListCategory(ctx, metadata)
}

func (cs CategoryService) GetCategory(ctx *gin.Context, id int16) (*models.Category, *models.ResponseError) {
	return cs.repositoryMgr.GetCategory(ctx, id)
}

func (cs CategoryService) CreateCategory(ctx *gin.Context, categoryParams *dbContext.CreateCategoryParams) (*models.Category, *models.ResponseError) {
	responseErr := validateCategory(categoryParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.repositoryMgr.CreateCategory(ctx, categoryParams)
}

func validateCategory(categoryParams *dbContext.CreateCategoryParams) *models.ResponseError {
	if categoryParams.CategoryID == 0 {
		return &models.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if categoryParams.CategoryName == "" {
		return &models.ResponseError{
			Message: "Invalid category name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
func (cs CategoryService) UpdateCategory(ctx *gin.Context, categoryParams *dbContext.CreateCategoryParams, id int64) *models.ResponseError {
	responseErr := validateCategory(categoryParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.repositoryMgr.UpdateCategory(ctx, categoryParams)
}

func (cs CategoryService) DeleteCategory(ctx *gin.Context, id int64) *models.ResponseError {
	return cs.repositoryMgr.DeleteCategory(ctx, id)
}

func (cs CategoryService) CreateCateProductDto(ctx *gin.Context, categoryWithProductDto *models.CreateCategoryProductDto) (*models.Category, *models.ResponseError) {

	err := repositories.BeginTransaction(cs.repositoryMgr)
	if err != nil {
		return nil, &models.ResponseError{
			Message: "Failed to start transaction",
			Status:  http.StatusBadRequest,
		}
	}
	//first query statement
	cs.CreateCategory(ctx, (*dbContext.CreateCategoryParams)(&categoryWithProductDto.CreateCategoryDto))
	// if responseErr != nil {
	// 	repositories.RollbackTransaction(cs.repositoryMgr)
	// 	return nil, responseErr
	// }

	// //second query statement
	// responseErr = cs.DeleteCategory(ctx, int64(response.CategoryID))
	// if responseErr != nil {
	// 	//when delete not succeed, transaction will rollback
	// 	repositories.RollbackTransaction(cs.repositoryMgr)
	// 	return nil, responseErr
	// }

	// if all statement ok, transaction will commit/save to db
	repositories.CommitTransaction(cs.repositoryMgr)
	return nil, &models.ResponseError{
		Message: "Data has been created",
		Status:  http.StatusOK,
	}
}
