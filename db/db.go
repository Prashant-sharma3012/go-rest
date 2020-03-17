package db

// in memory db should work for now

type DB struct {
	Coupons *CouponsCollection
}

var Dbinstance *DB

func Connect() *DB {
	if Dbinstance == nil {
		Dbinstance = &DB{
			Coupons: GetCouponsInstance(),
		}
	}

	return Dbinstance
}
