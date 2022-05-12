package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/domain"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/service"
)

type InventoryAPI struct {
	InventoryService service.IInventoryService
}

func NewInventoryAPI(service service.IInventoryService) InventoryAPI {
	return InventoryAPI{
		InventoryService: service,
	}
}

func (api InventoryAPI) Create(c *gin.Context) {
	var inventory domain.InventoryItem
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

func (api InventoryAPI) List(c *gin.Context) {
	inventories, err := api.InventoryService.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"inventories": inventories})
}

func (api InventoryAPI) GET(c *gin.Context) {
	slug := c.Param("slug")
	inventory, err := api.InventoryService.GET(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"inventory": inventory})
}

func (api InventoryAPI) Update(c *gin.Context) {
	var inventory domain.InventoryItem
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := api.InventoryService.Update(inventory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (api InventoryAPI) DELETE(c *gin.Context) {
	id := c.Param("id")
	if err := api.InventoryService.DELETE(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
