package server

import "net/http"

func (s *Server) ListCoupons(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from coupon list"))
}

func (s *Server) AddCoupon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from add coupon"))
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
