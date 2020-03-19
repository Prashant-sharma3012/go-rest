package db

// in memory db should work for now

type DB struct {
	Coupons   *CouponsCollection
	Employees *EmployeeCollection
	Inventory *InventoryCollection
	Sales     *SalesCollection
}

var Dbinstance *DB

func Connect() *DB {
	if Dbinstance == nil {
		Dbinstance = &DB{
			Coupons:   GetCouponsInstance(),
			Employees: GetEmployeeInstance(),
			Inventory: GetInventoryInstance(),
			Sales:     GetSalesInstance(),
		}
	}

	return Dbinstance
}
