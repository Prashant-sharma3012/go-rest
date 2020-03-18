package server

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	w.Write([]byte("Hello from coupon delete"))
}

func (s *Server) GetCouponByCode(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from get coupon by code"))
}
