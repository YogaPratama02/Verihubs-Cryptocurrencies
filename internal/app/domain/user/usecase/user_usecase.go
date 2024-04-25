package usecase

import (
	"log"
	"strconv"
	"time"
	"verihubs-cryptocurrencies/internal/app/domain/user/repository"
	"verihubs-cryptocurrencies/internal/app/dto"
	"verihubs-cryptocurrencies/internal/app/model"
	utill "verihubs-cryptocurrencies/internal/pkg/util"

	"github.com/labstack/echo"
	"github.com/yudapc/go-rupiah"
)

type UserUsecase interface {
	UGetCoinMarket(c echo.Context, pl dto.DataCoinMarket) ([]*dto.GetCoinMarketResponse, error)
	UCreateCoinTracker(c echo.Context, pl *dto.CreateCoinTracker) error
	UGetCoinTracker(c echo.Context, pl *dto.GetCoinTrackerRequest) ([]*dto.GetCoinTrackerResponse, error)
	UGetDetailListCoinTracker(c echo.Context, pl *dto.GetCoinTrackerRequest) (*dto.GetCoinTrackerResponse, error)
	UDeleteCoinTracker(c echo.Context, pl *dto.DeleteCoinTrackerRequest) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(repository repository.UserRepository) UserUsecase {
	return &userUsecase{repository}
}

func (s *userUsecase) UGetCoinMarket(c echo.Context, pl dto.DataCoinMarket) ([]*dto.GetCoinMarketResponse, error) {
	var result = []*dto.GetCoinMarketResponse{}
	exchangeRateRupiah, err := utill.FetchExchangeRates()
	if err != nil {
		log.Printf("Can't get the exchange rate with err: %s\n", err)
		return nil, err
	}

	for _, val := range pl.Data {
		priceUsdFloat, err := strconv.ParseFloat(val.PriceUsd, 64)
		if err != nil {
			return nil, err
		}
		priceRupiahFloat := exchangeRateRupiah * priceUsdFloat
		formatRupiah := rupiah.FormatRupiah(priceRupiahFloat)
		getCoinMarket := &dto.GetCoinMarketResponse{
			Id:     val.Id,
			Name:   val.Name,
			Price:  formatRupiah,
			Rank:   val.Rank,
			Symbol: val.Symbol,
		}

		result = append(result, getCoinMarket)
	}

	return result, nil
}

func (s *userUsecase) UCreateCoinTracker(c echo.Context, pl *dto.CreateCoinTracker) error {
	err := s.userRepository.RCreateCoinTracker(c, &model.CoinTracker{
		Name:      pl.Name,
		Rank:      pl.Rank,
		UserId:    pl.UserId,
		Symbol:    pl.Symbol,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *userUsecase) UGetCoinTracker(c echo.Context, pl *dto.GetCoinTrackerRequest) ([]*dto.GetCoinTrackerResponse, error) {
	var result []*dto.GetCoinTrackerResponse
	listCoinTracker, err := s.userRepository.RGetListCoinTracker(c, pl)
	if err != nil {
		return nil, err
	}

	for _, val := range listCoinTracker {
		data := &dto.GetCoinTrackerResponse{
			Id:        val.Id,
			UserId:    val.UserId,
			Name:      val.Name,
			Symbol:    val.Symbol,
			Rank:      val.Rank,
			CreatedAt: val.CreatedAt,
		}

		result = append(result, data)
	}

	return result, nil
}

func (s *userUsecase) UGetDetailListCoinTracker(c echo.Context, pl *dto.GetCoinTrackerRequest) (*dto.GetCoinTrackerResponse, error) {
	listCoinTracker, err := s.userRepository.RGetDetailListCoinTracker(c, pl)
	if err != nil {
		return nil, err
	}

	data := &dto.GetCoinTrackerResponse{
		Id:        listCoinTracker.Id,
		UserId:    listCoinTracker.UserId,
		Name:      listCoinTracker.Name,
		Symbol:    listCoinTracker.Symbol,
		Rank:      listCoinTracker.Rank,
		CreatedAt: listCoinTracker.CreatedAt,
	}

	return data, nil
}

func (s *userUsecase) UDeleteCoinTracker(c echo.Context, pl *dto.DeleteCoinTrackerRequest) error {
	if err := s.userRepository.RCheckDataCoinTracker(c, pl); err != nil {
		return err
	}

	if err := s.userRepository.RDeleteCoinTracker(c, pl); err != nil {
		return err
	}

	return nil
}
