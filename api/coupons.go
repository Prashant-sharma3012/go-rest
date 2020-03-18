package api

func (a *Api) InitCouponRoutes() {
	a.Coupons = a.Root.PathPrefix("/coupons").Subrouter()

	a.Coupons.HandleFunc("/list", a.Srv.ListCoupons).Methods("Get")
	a.Coupons.HandleFunc("/add", a.Srv.AddCoupon).Methods("Post")
	a.Coupons.HandleFunc("/update", a.Srv.UpdateCoupon).Methods("Put")
	a.Coupons.HandleFunc("/delete", a.Srv.DeleteCoupon).Methods("Delete").Queries("couponId", "{[0-9]*?}")
	a.Coupons.HandleFunc("/{couponCode}", a.Srv.GetCouponByCode).Methods("Get")
}
