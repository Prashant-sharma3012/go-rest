package api

func (a *Api) InitSalesRoutes() {
	a.Sales = a.Root.PathPrefix("/sales").Subrouter()

	a.Coupons.HandleFunc("/made", a.Srv.ListSales).Methods("Get")
	a.Coupons.HandleFunc("/make", a.Srv.MakeSale).Methods("Post")
	a.Coupons.HandleFunc("/{saleId}", a.Srv.GetSaleById).Methods("Get")
}
