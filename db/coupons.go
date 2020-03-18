package db

import (
	"errors"

	"github.com/go-rest/models"
)

type CouponsCollection struct {
	Coupons []models.Coupon
}

var collection *CouponsCollection

// to make sure only one instance is returned
func GetCouponsInstance() *CouponsCollection {
	if collection == nil {
		collection = &CouponsCollection{}
	}

	return collection
}

func (c *CouponsCollection) NextId() int {
	return len(c.Coupons) + 1
}

func (c *CouponsCollection) List() []models.Coupon {
	return c.Coupons
}

func (c *CouponsCollection) Create(coupon models.Coupon) string {
	c.Coupons = append(c.Coupons, coupon)
	return "coupon added successfully"
}

func (c *CouponsCollection) Update(coupon models.Coupon) (error, string) {
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

func (c *CouponsCollection) FindByCode(couponCode string) (error, models.Coupon) {

	for _, value := range c.Coupons {
		if value.Code == couponCode {
			return nil, value
		}
	}

	return errors.New("Not Found"), models.Coupon{}
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
