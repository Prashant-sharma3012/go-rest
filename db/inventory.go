package db

import (
	"errors"

	"github.com/go-rest/models"
)

type InventoryCollection struct {
	Inventory []models.Item
}

var invCol *InventoryCollection

// to make sure only one instance is returned
func GetInventoryInstance() *InventoryCollection {
	if collection == nil {
		invCol = &InventoryCollection{}
	}

	return invCol
}

func (i *InventoryCollection) itemQtyAvailable(item models.Item) bool {

	for _, value := range i.Inventory {
		if value.Id == item.Id && value.Quantity > item.Quantity {
			return true
		}
	}

	return false
}

func (i *InventoryCollection) ItemsQtyAvailable(items []models.Item) (bool, models.Item) {

	for _, value := range items {
		if !i.itemQtyAvailable(value) {
			return false, value
		}
	}

	return true, models.Item{}
}

func (i *InventoryCollection) UpdateInventory(items []models.Item) error {

	for _, value := range items {
		i.Update(value)
	}

	return nil
}

func (i *InventoryCollection) NextId() int {
	return len(i.Inventory) + 1
}

func (i *InventoryCollection) List() []models.Item {
	return i.Inventory
}

func (i *InventoryCollection) Create(item models.Item) string {
	i.Inventory = append(i.Inventory, item)
	return "Item added successfully"
}

func (i *InventoryCollection) Update(item models.Item) (error, string) {
	found := false

	for indx, value := range i.Inventory {
		if value.Id == item.Id {
			i.Inventory[indx] = item
			found = true
			break
		}
	}

	if found {
		return nil, "Updated successfully"
	}

	return errors.New("Not Found"), ""
}

func (i *InventoryCollection) FindById(itemId int) (error, models.Item) {

	for _, value := range i.Inventory {
		if value.Id == itemId {
			return nil, value
		}
	}

	return errors.New("Not Found"), models.Item{}
}

func (i *InventoryCollection) Delete(itemId int) (error, string) {
	removeIndx := -1

	for indx, value := range i.Inventory {
		if value.Id == itemId {
			removeIndx = indx
			break
		}
	}

	if removeIndx != -1 {
		i.Inventory = append(i.Inventory[:removeIndx], i.Inventory[removeIndx+1:]...)
		return nil, "Deleted"
	}

	return errors.New("Not Found"), ""
}
