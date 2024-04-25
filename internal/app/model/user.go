package model

import "time"

type CoinTracker struct {
	Id        int       `json:"id"`
	UserId    int64     `json:"user_id"`
	Name      string    `json:"name"`
	Rank      string    `json:"rank"`
	Symbol    string    `json:"symbol"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
