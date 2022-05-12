package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type inventoryAPI struct {
	InventoryService InventoryService
}

func (api inventoryAPI) Create(c *gin.Context) {
	var inventory Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := api.InventoryService.Create(inventory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
