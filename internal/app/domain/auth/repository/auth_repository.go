package repository

import (
	"database/sql"
	"log"
	"verihubs-cryptocurrencies/internal/app/model"

	"github.com/labstack/echo"
)

type AuthRepository interface {
	RRegister(c echo.Context, request *model.User) (int64, error)
	RLogin(c echo.Context, request *model.User) error
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) RRegister(c echo.Context, request *model.User) (int64, error) {
	db := r.db
	var lastInsertId int64
	sqlStatement := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	res, err := db.ExecContext(c.Request().Context(), sqlStatement, &request.UserName, &request.Email, &request.Password, &request.CreatedAt, &request.UpdatedAt)
	if err != nil {
		log.Printf("Error register to database with err: %s", err)
		return lastInsertId, err
	}

	if lastInsertId, err = res.LastInsertId(); err != nil {
		return lastInsertId, err
	}

	return lastInsertId, nil
}

func (r *authRepository) RLogin(c echo.Context, request *model.User) error {
	db := r.db

	sqlInfo := `SELECT id, name, email, password FROM users WHERE email = $1`
	err := db.QueryRow(sqlInfo, request.Email).Scan(&request.Id, &request.UserName, &request.Email, &request.Password)
	if err != nil {
		log.Printf("Phone number not found with err: %s", err)
		return err
	}

	return nil
}
