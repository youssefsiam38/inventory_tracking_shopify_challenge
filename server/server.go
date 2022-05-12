package server

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ieee0824/getenv"
	"net/http"
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
	itemsSubRouter.POST("/", apis.InventoryAPI.Create)
	itemsSubRouter.GET("/", apis.InventoryAPI.List)
	itemsSubRouter.GET("/:slug", apis.InventoryAPI.GET)
	itemsSubRouter.PUT("/", apis.InventoryAPI.Update)
	itemsSubRouter.DELETE("/", apis.InventoryAPI.DELETE)

	router.GET("/healthz", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("OK"))
	})

	router.Run(fmt.Sprintf(":%s", *port))
}
