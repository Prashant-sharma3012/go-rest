package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-rest/models"
	"github.com/gorilla/mux"
)

func (s *Server) listSales(w http.ResponseWriter, r *http.Request) {
	sales := s.DB.Sales.List()
	res, _ := json.Marshal(sales)

	w.Write(res)
}

func (s *Server) MakeSale(w http.ResponseWriter, r *http.Request) {
	sale := models.SaleFromJson(r.Body)
	// check if items are present in inventory and reduce teh quantity

	success := s.DB.Sales.Create(*sale)
	w.Write([]byte(success))
}

func (s *Server) GetSaleById(w http.ResponseWriter, r *http.Request) {
	saleId, _ := strconv.Atoi(mux.Vars(r)["saleId"])
	err, sale := s.DB.Sales.FindById(saleId)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	res, _ := json.Marshal(sale)
	w.Write(res)
}
