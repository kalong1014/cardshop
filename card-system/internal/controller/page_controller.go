package controller

import (
	"card-system/internal/model"
	"card-system/internal/service"
	"net/http"
	"strconv" // 添加此行

	"github.com/gin-gonic/gin"
)

type PageController struct {
	pageService *service.PageService
}

func NewPageController(pageService *service.PageService) *PageController {
	return &PageController{pageService: pageService}
}

func (c *PageController) CreatePage(ctx *gin.Context) {
	merchantID, err := strconv.ParseUint(ctx.Param("merchantID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid merchant ID"})
		return
	}

	var pageData struct {
		Name string `json:"name" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&pageData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page, err := c.pageService.CreatePage(uint(merchantID), pageData.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create page"})
		return
	}

	ctx.JSON(http.StatusCreated, page)
}

func (c *PageController) GetMerchantPages(ctx *gin.Context) {
	merchantID, err := strconv.ParseUint(ctx.Param("merchantID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid merchant ID"})
		return
	}

	pages, err := c.pageService.GetMerchantPages(uint(merchantID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pages"})
		return
	}

	ctx.JSON(http.StatusOK, pages)
}

func (c *PageController) GetPage(ctx *gin.Context) {
	pageID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page ID"})
		return
	}

	page, err := c.pageService.GetPageByID(uint(pageID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}

	ctx.JSON(http.StatusOK, page)
}

func (c *PageController) UpdatePage(ctx *gin.Context) {
	pageID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page ID"})
		return
	}

	var page model.Page
	if err := ctx.ShouldBindJSON(&page); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page.ID = uint(pageID)

	if err := c.pageService.UpdatePage(&page); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update page"})
		return
	}

	ctx.JSON(http.StatusOK, page)
}

func (c *PageController) DeletePage(ctx *gin.Context) {
	pageID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page ID"})
		return
	}

	if err := c.pageService.DeletePage(uint(pageID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete page"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Page deleted successfully"})
}
