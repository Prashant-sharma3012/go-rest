package db

import (
	"errors"
	"time"
)

type Coupon struct {
	Id     int
	Code   string
	Expiry time.Time
}

type CouponsCollection struct {
	Coupons []Coupon
}

var collection *CouponsCollection

// to make sure only one instance is returned
func GetCouponsInstance() *CouponsCollection {
	if collection == nil {
		collection = &CouponsCollection{}
		return collection
	}

	return collection
}

func (c *CouponsCollection) NextId() int {
	return len(c.Coupons)
}

func (c *CouponsCollection) Create(couponCode string) string {
	currLen := len(c.Coupons)

	coupon := Coupon{
		Id:     currLen,
		Code:   couponCode,
		Expiry: time.Now(),
	}

	c.Coupons = append(c.Coupons, coupon)
	return "coupon added successfully"
}

func (c *CouponsCollection) Update(coupon Coupon) (error, string) {
	found := false

	for indx, value := range c.Coupons {
		if value.Id == coupon.Id {
			c.Coupons[indx] = coupon
			found = true
			break
		}
	}

	if found {
		return nil, "Updated successfully"
	}

	return errors.New("Not Found"), ""
}

func (c *CouponsCollection) FindByCode(couponCode string) (error, Coupon) {

	for _, value := range c.Coupons {
		if value.Code == couponCode {
			return nil, value
		}
	}

	return errors.New("Not Found"), Coupon{}
}

func (c *CouponsCollection) Delete(couponCode string) (error, string) {
	removeIndx := -1

	for indx, value := range c.Coupons {
		if value.Code == couponCode {
			removeIndx = indx
			break
		}
	}

	if removeIndx != -1 {
		c.Coupons = append(c.Coupons[:removeIndx], c.Coupons[removeIndx+1:]...)
		return nil, "Deleted"
	}

	return errors.New("Not Found"), ""
}
