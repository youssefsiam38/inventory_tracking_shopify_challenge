package main

import (
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/repository"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/server"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/service"
)

func main() {
	repo := repository.NewInventoryRepository()
	service := service.NewInventoryService(repo)
	inventoryAPI := server.NewInventoryAPI(service)

	apis := server.APIs{
		InventoryAPI: inventoryAPI,
	}

	server.StartServer(apis)
}