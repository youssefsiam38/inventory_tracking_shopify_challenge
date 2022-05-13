package main

import (
	"flag"
	"fmt"

	"github.com/ieee0824/getenv"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/repository"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/server"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/service"
)

func main() {
	dbType := getenv.String("DB", "inmemory")
	db := flag.String("db", dbType, "db type (replit/inmemory)")
	flag.Parse()

	fmt.Println("db type:", *db)
	var inventoryRepository repository.IInventoryRepository
	if *db == "replit" {
		inventoryRepository = repository.NewInventoryReplitRepository()
	} else if *db == "inmemory" {
		inventoryRepository = repository.NewInventoryInMemoryRepository()
	} else {
		panic("Invalid db type")
	}
	service := service.NewInventoryService(inventoryRepository)
	inventoryAPI := server.NewInventoryAPI(service)

	apis := server.APIs{
		InventoryAPI: inventoryAPI,
	}

	server.StartServer(apis)
}
