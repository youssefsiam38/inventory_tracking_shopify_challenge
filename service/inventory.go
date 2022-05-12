package service

import (
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/domain"
)

type InventoryService interface {
	Create(inventory domain.Inventory) error
	List() ([]domain.Inventory, error)
	GET(slug string) (domain.Inventory, error)
	Update(inventory domain.Inventory) error
	DELETE(id string) error
}

type inventoryService struct {
	repository repository.InventoryRepository
}