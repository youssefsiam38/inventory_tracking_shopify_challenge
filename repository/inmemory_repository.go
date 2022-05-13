package repository

import (
	"fmt"

	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/domain"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/errors"
)

var inMemoryInventory map[string]domain.InventoryItem

func init() {
	inMemoryInventory = make(map[string]domain.InventoryItem)
}

type InventoryInMemoryRepository struct{}

func NewInventoryInMemoryRepository() IInventoryRepository {
	return InventoryInMemoryRepository{}
}

func (repository InventoryInMemoryRepository) Create(inventory domain.InventoryItem) error {
	inMemoryInventory[inventory.Slug] = inventory
	return nil
}

func (repository InventoryInMemoryRepository) List() ([]domain.InventoryItem, error) {
	inventoryItems := []domain.InventoryItem{}

	for key := range inMemoryInventory {
		inventoryItems = append(inventoryItems, inMemoryInventory[key])
	}
	return inventoryItems, nil
}

func (repository InventoryInMemoryRepository) GET(slug string) (domain.InventoryItem, error) {
	item, ok := inMemoryInventory[slug]
	if !ok {
		return item, errors.UserError{Err: nil, UserMessage: "Error getting inventory"}
	}
	return item, nil
}

func (repository InventoryInMemoryRepository) Update(item domain.InventoryItem) error {
	_, ok := inMemoryInventory[item.Slug]
	if !ok {
		return errors.UserError{Err: nil, UserMessage: "Error getting inventory"}
	}
	inMemoryInventory[item.Slug] = item
	return nil
}

func (repository InventoryInMemoryRepository) Delete(item domain.InventoryItem) error {
	item.Deleted = true
	item.UpdateUpdatedAt()
	inMemoryInventory[item.Slug] = item
	return nil
}
