package router

import (
	"database/sql"
	_authDelivery "verihubs-cryptocurrencies/internal/app/domain/auth/delivery"
	_userDelivery "verihubs-cryptocurrencies/internal/app/domain/user/delivery"
	"verihubs-cryptocurrencies/internal/pkg/middleware"

	"github.com/labstack/echo"
)

type Router struct {
	Echo         *echo.Echo
	Repositories repositories
	Usecases     usecases
}

func NewRoutes(db *sql.DB) Router {
	e := echo.New()
	e.Debug = false

	// CORS
	middleware.MiddlewareCors(e)

	// Logger
	middleware.LoggerConfig(e)

	repos := newRepositories(db)

	return Router{
		Echo:         e,
		Repositories: repos,
		Usecases:     newUsecases(repos),
	}
}

func (r *Router) LoadHandlers() {
	_authDelivery.NewAuthHandler(r.Echo, r.Usecases.AuthUsecase)
	_userDelivery.NewUserHandler(r.Echo, r.Usecases.UserUsecase)
}
