package api

func (a *Api) InitInventoryRoutes() {
	a.Items = a.Root.PathPrefix("/items").Subrouter()

	a.Coupons.HandleFunc("/list", a.Srv.ListItems).Methods("Get")
	a.Coupons.HandleFunc("/add", a.Srv.AddItem).Methods("Post")
	a.Coupons.HandleFunc("/update", a.Srv.UpdateItem).Methods("Put")
	a.Coupons.HandleFunc("/delete", a.Srv.DeleteItem).Methods("Delete").Queries("itemId", "{[0-9]*?}")
	a.Coupons.HandleFunc("/{itemId}", a.Srv.GetItemById).Methods("Get")
}
