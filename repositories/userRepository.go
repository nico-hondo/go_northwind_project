package repositories

import (
	"database/sql"
	"net/http"
	"strings"

	"codeid.northwind/models"
)

type UsersRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUsersRepository(dbHandler *sql.DB) *UsersRepository {
	return &UsersRepository{
		dbHandler: dbHandler,
	}
}

func (ur UsersRepository) Signup(users *models.User) *models.ResponseError {
	query := `
	INSERT INTO users(user_name,user_password,user_role)
	VALUES
	($1, crypt($2,gen_salt('bf')),$3)
	`

	rows, err := ur.dbHandler.Query(query, users.Username, users.Password, users.Roles)

	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()

	return &models.ResponseError{
		Message: "User already signup",
		Status:  http.StatusCreated,
	}
}

func (ur UsersRepository) LoginUser(username string, password string) (string, *models.ResponseError) {
	query := `
		SELECT user_id
		FROM users
		WHERE user_name = $1 and user_password = crypt($2, user_password)`

	rows, err := ur.dbHandler.Query(query, username, password)
	if err != nil {
		return "", &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()

	var id string
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return "", &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}

	if rows.Err() != nil {
		return "", &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return id, nil
}

func (ur UsersRepository) GetUserRole(accessToken string) (string, *models.ResponseError) {
	query := `
		SELECT user_role
		FROM users
		WHERE user_token = $1`

	rows, err := ur.dbHandler.Query(query, strings.Replace(accessToken, "Bearer ", "", 1))
	if err != nil {
		return "", &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()

	var role string
	for rows.Next() {
		err := rows.Scan(&role)
		if err != nil {
			return "", &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}

	if rows.Err() != nil {
		return "", &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return role, nil
}

func (ur UsersRepository) SetAccessToken(accessToken string, id string) *models.ResponseError {
	query := `UPDATE users SET user_token = $1 WHERE user_id = $2`

	_, err := ur.dbHandler.Exec(query, accessToken, id)
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (ur UsersRepository) RemoveAccessToken(accessToken string) *models.ResponseError {
	query := `UPDATE users SET user_token = '' WHERE user_token = $1`

	_, err := ur.dbHandler.Exec(query, strings.Replace(accessToken, "Bearer ", "", 1))
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}
