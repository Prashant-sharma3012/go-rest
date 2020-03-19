package db

import (
	"errors"

	"github.com/go-rest/models"
)

type SalesCollection struct {
	Sales []models.Sale
}

var salesCol *SalesCollection

// to make sure only one instance is returned
func GetSalesInstance() *SalesCollection {
	if collection == nil {
		salesCol = &SalesCollection{}
	}

	return salesCol
}

func (s *SalesCollection) NextId() int {
	return len(s.Sales) + 1
}

func (s *SalesCollection) List() []models.Sale {
	return s.Sales
}

func (s *SalesCollection) Create(sale models.Sale) string {
	s.Sales = append(s.Sales, sale)
	return "Item added successfully"
}

func (s *SalesCollection) FindById(saleId int) (error, models.Sale) {

	for _, value := range s.Sales {
		if value.Id == saleId {
			return nil, value
		}
	}

	return errors.New("Not Found"), models.Sale{}
}
