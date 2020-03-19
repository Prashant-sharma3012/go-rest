package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-rest/models"
	"github.com/gorilla/mux"
)

func (s *Server) ListItems(w http.ResponseWriter, r *http.Request) {
	items, err := json.Marshal(s.DB.Inventory.List())

	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(items)
}

func (s *Server) AddItem(w http.ResponseWriter, r *http.Request) {

	item := models.InventoryFromJson(r.Body)
	item.Id = s.DB.Inventory.NextId()

	success := s.DB.Inventory.Create(*item)

	w.Write([]byte(success))
}

func (s *Server) DeleteItem(w http.ResponseWriter, r *http.Request) {
	itemId, _ := strconv.Atoi(r.FormValue("itemId"))

	err, _ := s.DB.Inventory.Delete(itemId)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Item deleted"))
}

func (s *Server) UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("In Progress"))
}

func (s *Server) GetItemById(w http.ResponseWriter, r *http.Request) {
	itemId, _ := strconv.Atoi(mux.Vars(r)["itemId"])

	err, item := s.DB.Inventory.FindById(itemId)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	res, _ := json.Marshal(item)

	w.Write(res)
}
