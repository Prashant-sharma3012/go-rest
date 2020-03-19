package api

func (a *Api) InitEmployeeRoutes() {
	a.Employees = a.Root.PathPrefix("/employees").Subrouter()

	a.Coupons.HandleFunc("/list", a.Srv.ListEmployees).Methods("Get")
	a.Coupons.HandleFunc("/add", a.Srv.AddEmployee).Methods("Post")
	a.Coupons.HandleFunc("/update", a.Srv.UpdateEmployee).Methods("Put")
	a.Coupons.HandleFunc("/delete", a.Srv.DeleteEmployee).Methods("Delete").Queries("employeeId", "{[0-9]*?}")
	a.Coupons.HandleFunc("/{employeeId}", a.Srv.GetEmployeeById).Methods("Get")
}
