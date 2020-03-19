package models

import (
	"encoding/json"
	"io"
)

type Employee struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
}

func EmployeeFromJson(data io.Reader) *Employee {
	var emp *Employee
	json.NewDecoder(data).Decode(&emp)
	return emp
}
