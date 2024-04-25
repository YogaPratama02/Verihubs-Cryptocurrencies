package router

import (
	"database/sql"
	_authRepository "verihubs-cryptocurrencies/internal/app/domain/auth/repository"
	_userRepository "verihubs-cryptocurrencies/internal/app/domain/user/repository"
)

type repositories struct {
	AuthRepository _authRepository.AuthRepository
	UserRepository _userRepository.UserRepository
}

func newRepositories(db *sql.DB) repositories {
	authRepository := _authRepository.NewAuthRepository(db)
	userRepository := _userRepository.NewUserRepository(db)

	return repositories{
		AuthRepository: authRepository,
		UserRepository: userRepository,
	}
}
