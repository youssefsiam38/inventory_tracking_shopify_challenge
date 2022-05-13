package repository

import (
	"encoding/json"

	"github.com/replit/database-go"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/domain"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/errors"
)

type InventoryReplitRepository struct{}

func NewInventoryReplitRepository() IInventoryRepository {
	return InventoryReplitRepository{}
}

func (repository InventoryReplitRepository) Create(inventory domain.InventoryItem) error {
	bytes, err := json.Marshal(inventory)
	if err != nil {
		return errors.UserError{Err: err, UserMessage: "Error marshalling inventory"}
	}

	err = database.Set(inventory.Slug, string(bytes))
	if err != nil {
		return errors.UserError{Err: err, UserMessage: "Error creating inventory"}
	}
	return nil
}

func (repository InventoryReplitRepository) List() ([]domain.InventoryItem, error) {
	inventoryItems := []domain.InventoryItem{}
	list, err := database.ListKeys("")
	if err != nil {
		return nil, errors.UserError{Err: err, UserMessage: "Error listing inventory"}
	}

	for _, value := range list {
		inventoryItem := domain.InventoryItem{}
		err := json.Unmarshal([]byte(value), &inventoryItem)
		if err != nil {
			return nil, errors.UserError{Err: err, UserMessage: "Error unmarshalling inventory"}
		}
		inventoryItems = append(inventoryItems, inventoryItem)
	}
	return inventoryItems, nil
}

func (repository InventoryReplitRepository) GET(slug string) (domain.InventoryItem, error) {
	inventoryItem := domain.InventoryItem{}
	jsonInventoryItem, err := database.Get(slug)
	if err != nil {
		return inventoryItem, errors.UserError{Err: err, UserMessage: "Error getting inventory"}
	}
	json.Unmarshal([]byte(jsonInventoryItem), &inventoryItem)
	return inventoryItem, nil
}

func (repository InventoryReplitRepository) Update(inventory domain.InventoryItem) error {
	bytes, err := json.Marshal(inventory)
	if err != nil {
		return errors.UserError{Err: err, UserMessage: "Error marshalling inventory"}
	}
	err = database.Set(inventory.Slug, string(bytes))
	if err != nil {
		return errors.UserError{Err: err, UserMessage: "Error updating inventory"}
	}
	return nil
}

func (repository InventoryReplitRepository) Delete(item domain.InventoryItem) error {
	inventoryItem := domain.InventoryItem{}
	jsonInventoryItem, err := database.Get(item.Slug)
	if err != nil {
		return errors.UserError{Err: err, UserMessage: "Error getting inventory"}
	}

	err = json.Unmarshal([]byte(jsonInventoryItem), &inventoryItem)
	if err != nil {
		return errors.UserError{Err: err, UserMessage: "Error unmarshalling inventory"}
	}
	inventoryItem.Deleted = true
	inventoryItem.UpdateUpdatedAt()
	inventoryItem.DeleteComment = item.DeleteComment

	bytes, err := json.Marshal(inventoryItem)
	if err != nil {
		return errors.UserError{Err: err, UserMessage: "Error marshalling inventory"}
	}
	err = database.Set(inventoryItem.Slug, string(bytes))
	if err != nil {
		return errors.UserError{Err: err, UserMessage: "Error updating inventory"}
	}
	return nil
}
