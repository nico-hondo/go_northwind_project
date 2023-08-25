package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.northwind/models"
	"codeid.northwind/models/features"
	"codeid.northwind/repositories/dbContext"
	"codeid.northwind/services"
	"github.com/gin-gonic/gin"
)

const ROLE_ADMIN = "admin"
const ROLE_GUEST = "guest"

type CategoryController struct {
	categoryService *services.CategoryService
	userService     *services.UsersService
}

// declare construtor
func NewCategoryController(serviceMgr *services.ServiceManager) *CategoryController {
	return &CategoryController{
		categoryService: &serviceMgr.CategoryService,
		userService:     &serviceMgr.UsersService,
	} //passing service name as a parameter to the constructor.
}

// create method
func (categoryController CategoryController) GetListCategory(ctx *gin.Context) {
	// add authorization
	accessToken := ctx.Request.Header.Get("Authorization")
	auth, responseErr := categoryController.userService.AuthorizeUser(accessToken, []string{ROLE_ADMIN})
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	if !auth {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	//add metadata to hold data from query parameter, use defaultquery
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("pageNo", "0"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("PageSize", "0"))
	searchBy := ctx.DefaultQuery("searchBy", "")

	metadata := features.Metadata{
		PageNo:   pageNo,
		PageSize: pageSize,
		SearchBy: searchBy,
	}

	response, responseErr := categoryController.categoryService.GetListCategory(ctx, &metadata)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
	}
	ctx.JSON(http.StatusOK, response)
	// ctx.JSON(http.StatusOK, "Hello Gin Framework")
}

func (categoryController CategoryController) GetCategory(ctx *gin.Context) {

	categoryId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := categoryController.categoryService.GetCategory(ctx, int16(categoryId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (categoryController CategoryController) CreateCategory(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var category dbContext.CreateCategoryParams
	err = json.Unmarshal(body, &category)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := categoryController.categoryService.CreateCategory(ctx, &category)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (categoryController CategoryController) UpdateCategory(ctx *gin.Context) {

	categoryId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var category dbContext.CreateCategoryParams
	err = json.Unmarshal(body, &category)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := categoryController.categoryService.UpdateCategory(ctx, &category, int64(categoryId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (categoryController CategoryController) DeleteCategory(ctx *gin.Context) {

	categoryId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := categoryController.categoryService.DeleteCategory(ctx, int64(categoryId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (categoryController CategoryController) CreateCategoryWithProduct(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var category models.CreateCategoryProductDto
	err = json.Unmarshal(body, &category)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := categoryController.categoryService.CreateCateProductDto(ctx, &category)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}
