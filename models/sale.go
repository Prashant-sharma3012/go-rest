package models

import (
	"encoding/json"
	"io"
	"time"
)

type Sale struct {
	Id     int       `json:"id"`
	Items  []Item    `json:"items"`
	Amount int       `json:"amount"`
	Date   time.Time `json:"date"`
}

func SaleFromJson(data io.Reader) *Sale {
	var sale *Sale
	json.NewDecoder(data).Decode(&sale)
	return sale
}
