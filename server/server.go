package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/ieee0824/getenv"
)

type APIs struct {
	InventoryAPI InventoryAPI
}

func StartServer(apis APIs) {
	serverPort := getenv.String("PORT", "8080")
	port := flag.String("port", serverPort, "Port to listen on")
	flag.Parse()
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/inventory")
	})
	itemsSubRouter := router.Group("/inventory")
	itemsSubRouter.POST("/", inventoryAPI.InventoryService.Create)
	itemsSubRouter.GET("/", inventoryAPI.InventoryService.List)
	itemsSubRouter.GET("/{slug}", inventoryAPI.InventoryService.GET)
	itemsSubRouter.PUT("/", inventoryAPI.InventoryService.Update)
	itemsSubRouter.DELETE("/", inventoryAPI.InventoryService.DELETE)

	router.GET("/healthz", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("OK"))
	})

	router.Run(fmt.Sprintf(":%s", *port))
}
