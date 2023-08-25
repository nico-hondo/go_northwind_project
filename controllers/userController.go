package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"codeid.northwind/models"
	"codeid.northwind/services"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	usersService *services.UsersService
}

func NewUsersController(usersService *services.UsersService) *UsersController {
	return &UsersController{
		usersService: usersService,
	}
}

func (uc UsersController) Signup(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := uc.usersService.Signup(ctx, &user)
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

}

func (uc UsersController) Signin(ctx *gin.Context) {
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		log.Println("Error while reading credentials")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	accessToken, responseErr := uc.usersService.Login(username, password)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, &models.User{
		Username: username,
		Token:    accessToken,
	})
}

func (uc UsersController) Logout(ctx *gin.Context) {
	accessToken := ctx.Request.Header.Get("Authorization")

	responseErr := uc.usersService.Logout(accessToken)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
