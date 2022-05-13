package server

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/inventory")
	})
	itemsSubRouter := router.Group("/inventory")

	// create
	itemsSubRouter.GET("/create", apis.InventoryAPI.CreateView)
	itemsSubRouter.POST("/", apis.InventoryAPI.Create)

	// read
	itemsSubRouter.GET("/", apis.InventoryAPI.List)
	itemsSubRouter.GET("/:slug", apis.InventoryAPI.Get)

	// update
	itemsSubRouter.GET("/edit/:slug", apis.InventoryAPI.UpdateView)
	itemsSubRouter.POST("/:slug", apis.InventoryAPI.Update)

	// delete
	itemsSubRouter.POST("/delete", apis.InventoryAPI.Delete)

	router.GET("/healthz", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("OK"))
	})

	router.Run(fmt.Sprintf(":%s", *port))
}
