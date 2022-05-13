package repository

import "github.com/youssefsiam38/inventory_tracking_shopify_challenge/domain"

type IInventoryRepository interface {
	Create(inventory domain.InventoryItem) error
	List() ([]domain.InventoryItem, error)
	GET(slug string) (domain.InventoryItem, error)
	Update(inventory domain.InventoryItem) error
	Delete(inventory domain.InventoryItem) error
}