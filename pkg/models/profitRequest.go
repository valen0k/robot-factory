package models

type ProfitRequest struct {
	Year       int `json:"year"`
	Month      int `json:"month"`
	Day        int `json:"day"`
	AmountDays int `json:"amount_days"`
}
