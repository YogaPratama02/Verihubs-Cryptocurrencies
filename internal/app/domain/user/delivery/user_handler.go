package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"
	"verihubs-cryptocurrencies/internal/app/domain/user/usecase"
	"verihubs-cryptocurrencies/internal/app/dto"
	"verihubs-cryptocurrencies/internal/pkg/middleware"
	"verihubs-cryptocurrencies/internal/pkg/response"
	utill "verihubs-cryptocurrencies/internal/pkg/util"
	"verihubs-cryptocurrencies/internal/pkg/validation"

	"github.com/labstack/echo"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(e *echo.Echo, userUsecase usecase.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: userUsecase,
	}

	route := e.Group("/api/user")
	route.Use(middleware.ValidateToken)
	route.GET(`/coin-market`, handler.ListCoins)
	route.POST(`/coin-tracker`, handler.CreateCoinTracker)
	route.GET(`/coin-tracker`, handler.GetListCoinTracker)
	route.DELETE(`/coin-tracker/:id`, handler.DeleteCoinTracker)
}

func (h *UserHandler) ListCoins(c echo.Context) error {
	const (
		listCoinsAPIURL = "https://api.coincap.io/v2/assets"
	)

	resp, err := http.Get(listCoinsAPIURL)
	if err != nil {
		response.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var data dto.DataCoinMarket
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		response.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	result, err := h.UserUsecase.UGetCoinMarket(c, data)
	if err != nil {
		response.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	response.NewHandlerResponse("Succesfully get list coin market", result).Success(c)
	return nil
}

func (h *UserHandler) CreateCoinTracker(c echo.Context) error {
	var (
		pl          dto.CreateCoinTracker
		jwtResponse utill.ResponseTokenJwt = c.Get("token-data").(utill.ResponseTokenJwt)
	)

	pl.UserId = int64(jwtResponse.Id)

	if err := c.Bind(&pl); err != nil {
		response.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return nil
	}

	if err := validation.DoValidation(&pl); err != nil {
		response.NewHandlerResponse(err, nil).BadRequest(c)
		return nil
	}

	if err := h.UserUsecase.UCreateCoinTracker(c, &pl); err != nil {
		response.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	response.NewHandlerResponse("Successfully create coin tracker", nil).Success(c)
	return nil
}

func (h *UserHandler) GetListCoinTracker(c echo.Context) error {
	var (
		pl          dto.GetCoinTrackerRequest
		jwtResponse utill.ResponseTokenJwt = c.Get("token-data").(utill.ResponseTokenJwt)
	)
	pl.UserId = int64(jwtResponse.Id)

	if err := c.Bind(&pl); err != nil {
		response.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return nil
	}

	if pl.Id != 0 {
		resp, err := h.UserUsecase.UGetDetailListCoinTracker(c, &pl)
		if err != nil {
			response.NewHandlerResponse(err.Error(), nil).Failed(c)
			return nil
		}

		response.NewHandlerResponse("Successfully get list detail coin tracker user", resp).Success(c)
		return nil
	}

	resp, err := h.UserUsecase.UGetCoinTracker(c, &pl)
	if err != nil {
		response.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	response.NewHandlerResponse("Successfully get list coin tracker user", resp).Success(c)
	return nil
}

func (h *UserHandler) DeleteCoinTracker(c echo.Context) error {
	var (
		err         error
		pl          dto.DeleteCoinTrackerRequest
		jwtResponse utill.ResponseTokenJwt = c.Get("token-data").(utill.ResponseTokenJwt)
	)

	pl.UserId = int64(jwtResponse.Id)
	pl.Id, err = strconv.Atoi(c.Param("id"))

	if err != nil {
		response.NewHandlerResponse(err, nil).BadRequest(c)
		return nil
	}

	if err := validation.DoValidation(&pl); err != nil {
		response.NewHandlerResponse(err, nil).BadRequest(c)
		return nil
	}

	if err = h.UserUsecase.UDeleteCoinTracker(c, &pl); err != nil {
		response.NewHandlerResponse(err.Error(), nil).Failed(c)
		return nil
	}

	response.NewHandlerResponse("Successfully delete coin tracker", nil).Success(c)
	return nil
}
