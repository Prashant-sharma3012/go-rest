package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-rest/models"
)

func (s *Server) ListCoupons(w http.ResponseWriter, r *http.Request) {
	coupons, err := json.Marshal(s.DB.Coupons.List())

	if err != nil {
		fmt.Println(err)
	}

	w.Write(coupons)
}

func (s *Server) AddCoupon(w http.ResponseWriter, r *http.Request) {

	coupon := models.CouponFromJson(r.Body)
	coupon.Id = s.DB.Coupons.NextId()

	success := s.DB.Coupons.Create(*coupon)

	w.Write([]byte(success))
}

func (s *Server) UpdateCoupon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from update coupon"))
}

func (s *Server) DeleteCoupon(w http.ResponseWriter, r *http.Request) {
	couponId, _ := strconv.Atoi(r.FormValue("couponId"))

	err, _ := s.DB.Coupons.Delete(couponId)
	if err != nil {
		w.Write([]byte("An Error Occured"))
	}

	w.Write([]byte("Coupon deleted"))
}

func (s *Server) GetCouponByCode(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from get coupon by code"))
}
