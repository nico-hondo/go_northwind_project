package services

import (
	"encoding/base64"
	"net/http"

	"codeid.northwind/models"
	"codeid.northwind/repositories"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	repoMgr *repositories.RepositoryManager
}

func NewUsersService(repoMgr *repositories.RepositoryManager) *UsersService {
	return &UsersService{
		repoMgr: repoMgr,
	}
}

func (us UsersService) Signup(ctx *gin.Context, users *models.User) *models.ResponseError {
	if users.Username == "" || users.Password == "" {
		return &models.ResponseError{
			Message: "username or password must not empty",
			Status:  http.StatusBadRequest,
		}
	}
	return us.repoMgr.UsersRepository.Signup(users)
}

func (us UsersService) Login(username string, password string) (string, *models.ResponseError) {
	if username == "" || password == "" {
		return "", &models.ResponseError{
			Message: "Invalid username or password",
			Status:  http.StatusBadRequest,
		}
	}

	id, responseErr := us.repoMgr.UsersRepository.LoginUser(username, password)
	if responseErr != nil {
		return "", responseErr
	}

	if id == "" {
		return "", &models.ResponseError{
			Message: "Login failed",
			Status:  http.StatusUnauthorized,
		}
	}

	accessToken, responseErr := generateAccessToken(username)
	if responseErr != nil {
		return "", responseErr
	}

	us.repoMgr.UsersRepository.SetAccessToken(accessToken, id)

	return accessToken, nil
}

func (us UsersService) Logout(accessToken string) *models.ResponseError {
	if accessToken == "" {
		return &models.ResponseError{
			Message: "Invalid access token",
			Status:  http.StatusBadRequest,
		}
	}

	return us.repoMgr.UsersRepository.RemoveAccessToken(accessToken)
}

func (us UsersService) AuthorizeUser(accessToken string, roles []string) (bool, *models.ResponseError) {
	if accessToken == "" {
		return false, &models.ResponseError{
			Message: "Invalid access token",
			Status:  http.StatusBadRequest,
		}
	}

	role, responseErr := us.repoMgr.UsersRepository.GetUserRole(accessToken)
	if responseErr != nil {
		return false, responseErr
	}

	if role == "" {
		return false, &models.ResponseError{
			Message: "Failed to authorize user",
			Status:  http.StatusUnauthorized,
		}
	}

	for _, expectedRole := range roles {
		if expectedRole == role {
			return true, nil
		}
	}

	return false, nil
}

func generateAccessToken(username string) (string, *models.ResponseError) {
	hash, err := bcrypt.GenerateFromPassword([]byte(username), bcrypt.DefaultCost)
	if err != nil {
		return "", &models.ResponseError{
			Message: "Failed to generate token",
			Status:  http.StatusInternalServerError,
		}
	}

	return base64.StdEncoding.EncodeToString(hash), nil
}
