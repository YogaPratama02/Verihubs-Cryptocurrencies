package dto

import "time"

type DataCoinMarket struct {
	Data []AssetCoinResponse `json:"data"`
}
type AssetCoinResponse struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Explorer          string `json:"explorer"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUsd      string `json:"marketCapUsd"`
	PriceUsd          string `json:"priceUsd"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Supply            string `json:"supply"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
}

type GetCoinMarketResponse struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Price  string `json:"price"`
	Rank   string `json:"rank"`
	Symbol string `json:"symbol"`
}

type CreateCoinTracker struct {
	Name   string `json:"name" validate:"required"`
	UserId int64  `json:"user_id" validate:"required"`
	Rank   string `json:"rank" validate:"required"`
	Symbol string `json:"symbol" validate:"required"`
}

type GetCoinTrackerRequest struct {
	Id     int64 `json:"id"`
	UserId int64 `json:"user_id"`
	Limit  int64 `json:"name"`
	Page   int64 `json:"page"`
}

type GetCoinTrackerResponse struct {
	Id        int       `json:"id"`
	UserId    int64     `json:"user_id"`
	Name      string    `json:"name"`
	Rank      string    `json:"rank"`
	Symbol    string    `json:"symbol"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type DeleteCoinTrackerRequest struct {
	Id     int   `json:"id" validate:"required" param:"id"`
	UserId int64 `json:"user_id" validate:"required"`
}
