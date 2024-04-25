package usecase

import (
	"log"
	"time"
	"verihubs-cryptocurrencies/internal/app/domain/auth/repository"
	"verihubs-cryptocurrencies/internal/app/dto"
	"verihubs-cryptocurrencies/internal/app/model"
	utill "verihubs-cryptocurrencies/internal/pkg/util"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	URegister(c echo.Context, pl *dto.RegisterRequest) (*dto.RegisterResponse, error)
	ULogin(c echo.Context, pl *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authUsecase struct {
	authRepository repository.AuthRepository
}

func NewAuthUsecase(repository repository.AuthRepository) AuthUsecase {
	return &authUsecase{repository}
}

func (s *authUsecase) URegister(c echo.Context, pl *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	if pl.Password != pl.PasswordConfirmation {
		return nil, errors.New("Password doesn't match")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(pl.Password), 4)
	if err != nil {
		return nil, err
	}
	pl.Password = string(password)

	registerCreate := model.User{
		UserName:  pl.UserName,
		Email:     pl.Email,
		Password:  pl.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	responseId, err := s.authRepository.RRegister(c, &registerCreate)
	if err != nil {
		return nil, err
	}

	payloadResponse := dto.RegisterResponse{
		Id:       responseId,
		UserName: pl.UserName,
		Email:    pl.Email,
	}

	return &payloadResponse, nil
}

func (s *authUsecase) ULogin(c echo.Context, pl *dto.LoginRequest) (*dto.LoginResponse, error) {
	data := &model.User{
		Email: pl.Email,
	}

	if err := s.authRepository.RLogin(c, data); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(pl.Password)); err != nil {
		log.Printf("Email or Password Incorrect with err: %s\n", err)
		return nil, err
	}

	tokenResponse, err := utill.GenerateJWT(data)
	if err != nil {
		log.Printf("Can't generate JTW with err: %s\n", err)
		return nil, err
	}

	resp := &dto.LoginResponse{
		Token: tokenResponse.Token,
	}

	return resp, nil
}
