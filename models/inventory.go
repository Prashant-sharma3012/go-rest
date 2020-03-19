package models

import (
	"encoding/json"
	"io"
)

type Item struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func InventoryFromJson(data io.Reader) *Item {
	var item *Item
	json.NewDecoder(data).Decode(&item)
	return item
}
