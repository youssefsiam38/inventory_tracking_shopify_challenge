package domain

import (
	"time"
	"github.com/google/uuid"
)

type InventoryItem struct {
	ID          string    `json:"id"`
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Deleted     bool      `json:"deleted"`
}

func (inventory InventoryItem) IsDeleted() bool {
	return inventory.Deleted
}

func (inventory *InventoryItem) UpdateUpdatedAt() {
	inventory.UpdatedAt = time.Now()
}

func (inventory *InventoryItem) SetCreatedAt() {
	inventory.CreatedAt = time.Now()
}

func (inventory *InventoryItem) GenerateID() {
	inventory.ID = uuid.New().String()
}