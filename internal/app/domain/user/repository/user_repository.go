package repository

import (
	"database/sql"
	"log"
	"time"
	"verihubs-cryptocurrencies/internal/app/dto"
	"verihubs-cryptocurrencies/internal/app/model"

	"github.com/labstack/echo"
)

type UserRepository interface {
	RCreateCoinTracker(c echo.Context, request *model.CoinTracker) error
	RGetListCoinTracker(c echo.Context, pl *dto.GetCoinTrackerRequest) ([]*model.CoinTracker, error)
	RGetDetailListCoinTracker(c echo.Context, pl *dto.GetCoinTrackerRequest) (*model.CoinTracker, error)
	RCheckDataCoinTracker(c echo.Context, pl *dto.DeleteCoinTrackerRequest) error
	RDeleteCoinTracker(c echo.Context, pl *dto.DeleteCoinTrackerRequest) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) RCreateCoinTracker(c echo.Context, request *model.CoinTracker) error {
	db := r.db
	var lastInsertId int64
	sqlStatement := `INSERT INTO coin_trackers (user_id, name, rank, symbol, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	res, err := db.ExecContext(c.Request().Context(), sqlStatement, &request.UserId, &request.Name, &request.Rank, &request.Symbol, &request.CreatedAt, &request.UpdatedAt)
	if err != nil {
		log.Printf("Error create coin tracker to database with err: %s", err)
		return err
	}

	if lastInsertId, err = res.LastInsertId(); err != nil {
		return err
	}

	log.Printf("Successfully create coin tracker to database with id: %d", lastInsertId)
	return nil
}

func (r *userRepository) RGetListCoinTracker(c echo.Context, pl *dto.GetCoinTrackerRequest) ([]*model.CoinTracker, error) {
	var result []*model.CoinTracker
	db := r.db
	sqlStatement := `SELECT id, user_id, name, rank, symbol, datetime(created_at) FROM coin_trackers WHERE user_id = ? LIMIT ? OFFSET ?`

	rows, err := db.Query(sqlStatement, pl.UserId, pl.Limit, (pl.Page-1)*pl.Limit)
	if err != nil {
		log.Printf("Error get list coin tracker to database with err: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		data := model.CoinTracker{}
		var createdAtStr string
		err = rows.Scan(&data.Id, &data.UserId, &data.Name, &data.Rank, &data.Symbol, &createdAtStr)
		if err != nil {
			log.Printf("Error get list coin tracker to database with err: %s", err)
			return nil, err
		}

		createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			log.Printf("Error can't convert created_at with err: %s", err)
			return nil, err
		}
		data.CreatedAt = createdAt

		result = append(result, &data)
	}

	return result, nil
}

func (r *userRepository) RGetDetailListCoinTracker(c echo.Context, pl *dto.GetCoinTrackerRequest) (*model.CoinTracker, error) {
	var result model.CoinTracker
	var createdAtStr string
	db := r.db
	sqlStatement := `SELECT id, user_id, name, rank, symbol, datetime(created_at) FROM coin_trackers WHERE id = ? AND user_id = ?`

	row := db.QueryRow(sqlStatement, pl.Id, pl.UserId)

	if err := row.Scan(&result.Id, &result.UserId, &result.Name, &result.Rank, &result.Symbol, &createdAtStr); err != nil {
		log.Printf("Error get detail coin tracker with err: %s", err)
		return nil, err
	}

	createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		log.Printf("Error can't convert created_at with err: %s", err)
		return nil, err
	}

	result.CreatedAt = createdAt

	return &result, nil
}

func (r *userRepository) RCheckDataCoinTracker(c echo.Context, pl *dto.DeleteCoinTrackerRequest) error {
	var result model.CoinTracker
	db := r.db
	sqlStatement := `SELECT id FROM coin_trackers WHERE id = ? AND user_id = ?`

	row := db.QueryRow(sqlStatement, pl.Id, pl.UserId)

	if err := row.Scan(&result.Id); err != nil {
		log.Printf("Error check data coin tracker to database with err: %s", err)
		return err
	}

	return nil
}

func (r *userRepository) RDeleteCoinTracker(c echo.Context, pl *dto.DeleteCoinTrackerRequest) error {
	var lastInsertId int64

	db := r.db
	sqlStatement := `DELETE FROM coin_trackers WHERE id = ? AND user_id = ?`

	res, err := db.Exec(sqlStatement, pl.Id, pl.UserId)
	if err != nil {
		log.Printf("Error delete coin tracker to database with err: %s", err)
		return err
	}

	if lastInsertId, err = res.LastInsertId(); err != nil {
		return err
	}

	log.Printf("Successfully delete coin tracker to database with id: %d", lastInsertId)

	return nil
}
