package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-rest/models"
	"github.com/gorilla/mux"
)

func (s *Server) ListEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := json.Marshal(s.DB.Employees.List())

	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(employees)
}

func (s *Server) AddEmployee(w http.ResponseWriter, r *http.Request) {

	employee := models.EmployeeFromJson(r.Body)
	employee.Id = s.DB.Coupons.NextId()

	success := s.DB.Employees.Create(*employee)

	w.Write([]byte(success))
}

func (s *Server) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("in progress"))
}

func (s *Server) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	employeeId, _ := strconv.Atoi(r.FormValue("employeeId"))

	err, _ := s.DB.Employees.Delete(employeeId)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Employee deleted"))
}

func (s *Server) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	employeeId, _ := strconv.Atoi(mux.Vars(r)["employeeId"])

	err, employee := s.DB.Employees.FindById(employeeId)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	res, _ := json.Marshal(employee)

	w.Write(res)
}
