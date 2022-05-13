package service

import (
	"fmt"

	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/domain"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/errors"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/repository"
)

type IInventoryService interface {
	Create(inventory domain.InventoryItem) error
	List(deleted bool) ([]domain.InventoryItem, error)
	GET(slug string) (domain.InventoryItem, error)
	Update(inventory domain.InventoryItem) error
	Delete(slug string) error
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

	_, err := service.repository.GET(inventory.Slug)
	if err != nil {
		inventory.GenerateID()
		inventory.SetCreatedAt()
		inventory.UpdateUpdatedAt()
		return service.repository.Create(inventory)
	}
	return errors.UserError{Err: nil, UserMessage: "There is already an item with this slug"}
}

func (service InventoryService) List(deleted bool) ([]domain.InventoryItem, error) {
	items, err := service.repository.List()
	if err != nil {
		return items, err
	}
	fmt.Println(items)
	if deleted {
		deletedItems := []domain.InventoryItem{}
		for _, v := range items {
			if v.Deleted {
				deletedItems = append(deletedItems, v)
			}
		}
		items = deletedItems
	} else {
		activeItems := []domain.InventoryItem{}
		for _, v := range items {
			if !v.Deleted {
				activeItems = append(activeItems, v)
			}
		}
		items = activeItems
	}
	return items, err
}

func (service InventoryService) GET(slug string) (domain.InventoryItem, error) {
	return service.repository.GET(slug)
}

func (service InventoryService) Update(item domain.InventoryItem) error {
	oldItem, err := service.repository.GET(item.Slug)
	if err == nil {
		item.UpdateUpdatedAt()
		item.CreatedAt = oldItem.CreatedAt
		return service.repository.Update(item)
	}
	return errors.UserError{Err: nil, UserMessage: "There is already an item with this slug"}
}

func (service InventoryService) Delete(slug string) error {
	item, err := service.repository.GET(slug)
	if err != nil {
		return errors.UserError{Err: nil, UserMessage: "There is no item with this slug"}
	}
	item.UpdateUpdatedAt()
	return service.repository.Delete(item)
}