package models

import (
	"encoding/json"
	"io"
	"time"
)

type Coupon struct {
	Id     int       `json:"id"`
	Code   string    `json:"code"`
	Expiry time.Time `json:"expiry"`
}

func CouponFromJson(data io.Reader) *Coupon {
	var coupon *Coupon
	json.NewDecoder(data).Decode(&coupon)
	return coupon
}
