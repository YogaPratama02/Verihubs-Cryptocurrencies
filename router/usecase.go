package router

import (
	_authUsecase "verihubs-cryptocurrencies/internal/app/domain/auth/usecase"
	_userUsecase "verihubs-cryptocurrencies/internal/app/domain/user/usecase"
)

type usecases struct {
	AuthUsecase _authUsecase.AuthUsecase
	UserUsecase _userUsecase.UserUsecase
}

func newUsecases(repo repositories) usecases {
	authUsecase := _authUsecase.NewAuthUsecase(repo.AuthRepository)
	userUsecase := _userUsecase.NewUserUsecase(repo.UserRepository)

	return usecases{
		AuthUsecase: authUsecase,
		UserUsecase: userUsecase,
	}
}
