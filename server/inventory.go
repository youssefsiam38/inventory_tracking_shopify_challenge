package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/domain"
	"github.com/youssefsiam38/inventory_tracking_shopify_challenge/errors"
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

func (api InventoryAPI) CreateView(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", nil)
}

func (api InventoryAPI) Create(c *gin.Context) {
	item := domain.InventoryItem{}
	item.Name = c.PostForm("name")
	item.Description = c.PostForm("description")
	item.Slug = c.PostForm("slug")

	if err := api.InventoryService.Create(item); err != nil {
		fmt.Println("Error: ", err)
		items, _ := api.InventoryService.List(false)
		fmt.Println("Error: ", err)
		if err, ok := errors.IsUserError(err); ok {
			c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err.UserError()})
			return
		}
		c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err})
		return
	}

	items, err := api.InventoryService.List(false)
	if err != nil {
		if err, ok := errors.IsUserError(err); ok {
			c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err.UserError()})
			return
		}
		c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err})
		return
	}
	c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err, "Message": "Created Successfully"})
}

func (api InventoryAPI) List(c *gin.Context) {
	strDeleted := c.Query("deleted")
	fmt.Println("strDeleted: ", strDeleted)
	deleted, err := strconv.ParseBool(strDeleted)
	fmt.Println("deleted: ", deleted)
	if err != nil {
		deleted = false
	}
	items, err := api.InventoryService.List(deleted)
	if err != nil {
		c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err})
		return
	}

	c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err, "Message": c.Query("msg")})
}

func (api InventoryAPI) Get(c *gin.Context) {
	slug := c.Param("slug")
	item, err := api.InventoryService.GET(slug)
	if err != nil {
		items, err := api.InventoryService.List(false)
		c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err})
		return
	}
	c.HTML(http.StatusOK, "show.html", item)
}

func (api InventoryAPI) UpdateView(c *gin.Context) {
	slug := c.Param("slug")
	item, err := api.InventoryService.GET(slug)
	if err != nil {
		items, err := api.InventoryService.List(false)
		if err, ok := errors.IsUserError(err); ok {
			c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err.UserError()})
			return
		}
		c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err})
		return
	}
	c.HTML(http.StatusOK, "edit.html", item)
}

func (api InventoryAPI) Update(c *gin.Context) {
	item := domain.InventoryItem{}
	item.Name = c.PostForm("name")
	item.Description = c.PostForm("description")
	item.Slug = c.PostForm("slug")
	strDeleted := c.PostForm("deleted")
	var err error
	item.Deleted, err = strconv.ParseBool(strDeleted)
	if err != nil {
		item.Deleted = false
	}
	if err := api.InventoryService.Update(item); err != nil {
		items, err := api.InventoryService.List(false)
		if err, ok := errors.IsUserError(err); ok {
			c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err.UserError()})
			return
		}
		c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err})
		return
	}

	items, err := api.InventoryService.List(false)
	if err != nil {
		if err, ok := errors.IsUserError(err); ok {
			c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err.UserError()})
			return
		}
		c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err})
		return
	}
	c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err, "Message": "Updated Successfully"})
}

func (api InventoryAPI) Delete(c *gin.Context) {

	slug := c.PostForm("slug")
	comment := c.PostForm("comment")
	if err := api.InventoryService.Delete(slug, comment); err != nil {
		items, err := api.InventoryService.List(false)
		if err, ok := errors.IsUserError(err); ok {
			c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err.UserError()})
			return
		}
		c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err})
		return
	}

	items, err := api.InventoryService.List(false)
	if err != nil {
		if err, ok := errors.IsUserError(err); ok {
			c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err.UserError()})
			return
		}
		c.HTML(http.StatusOK, "list.html", gin.H{"Items": items, "Error": err})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/inventory?msg=Deleted Successfully")
}
