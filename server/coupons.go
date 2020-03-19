package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-rest/models"
	"github.com/gorilla/mux"
)

func (s *Server) ListCoupons(w http.ResponseWriter, r *http.Request) {
	coupons, err := json.Marshal(s.DB.Coupons.List())

	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(coupons)
}

func (s *Server) AddCoupon(w http.ResponseWriter, r *http.Request) {

	coupon := models.CouponFromJson(r.Body)
	coupon.Id = s.DB.Coupons.NextId()

	success := s.DB.Coupons.Create(*coupon)

	w.Write([]byte(success))
}

func (s *Server) DeleteCoupon(w http.ResponseWriter, r *http.Request) {
	couponId, _ := strconv.Atoi(r.FormValue("couponId"))

	err, _ := s.DB.Coupons.Delete(couponId)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Coupon deleted"))
}

func (s *Server) GetCouponByCode(w http.ResponseWriter, r *http.Request) {
	couponCode := mux.Vars(r)["couponCode"]

	err, coupon := s.DB.Coupons.FindByCode(couponCode)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	res, _ := json.Marshal(coupon)

	w.Write(res)
}
