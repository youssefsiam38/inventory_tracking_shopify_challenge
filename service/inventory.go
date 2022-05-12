package service

import (
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/domain"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/repository"
)

type IInventoryService interface {
	Create(inventory domain.InventoryItem) error
	List() ([]domain.InventoryItem, error)
	GET(slug string) (domain.InventoryItem, error)
	Update(inventory domain.InventoryItem) error
	DELETE(id string) error
}

type InventoryService struct {
	repository repository.IInventoryRepository
}

func NewInventoryService(repo repository.IInventoryRepository) IInventoryService {
	return InventoryService{
		repository: repo,
	}
}

func (service InventoryService) Create(inventory domain.InventoryItem) error {
	return service.repository.Create(inventory)
}

func (service InventoryService) List() ([]domain.InventoryItem, error) {
	return service.repository.List()
}

func (service InventoryService) GET(slug string) (domain.InventoryItem, error) {
	return service.repository.GET(slug)
}

func (service InventoryService) Update(inventory domain.InventoryItem) error {
	return service.repository.Update(inventory)
}

func (service InventoryService) DELETE(slug string) error {
	return service.repository.DELETE(slug)
}