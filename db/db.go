package db

// in memory db should work for now

type DB struct {
	Coupons *CouponsCollection
}

func Connect() *DB {
	return &DB{}
}
