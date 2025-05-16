package controller

import (
	"card-system/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PageController struct {
	pageService *service.PageService
}

func NewPageController(pageService *service.PageService) *PageController {
	return &PageController{pageService: pageService}
}

// CreatePage 创建页面
func (c *PageController) CreatePage(ctx *gin.Context) {
	// 获取当前商户ID
	merchantID, exists := ctx.Get("merchant_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	// 解析请求
	var request struct {
		Name string `json:"name" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建页面
	page, err := c.pageService.CreatePage(merchantID.(uint), request.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "创建页面失败"})
		return
	}

	ctx.JSON(http.StatusCreated, page)
}

// GetMerchantPages 获取商户所有页面
func (c *PageController) GetMerchantPages(ctx *gin.Context) {
	// 获取当前商户ID
	merchantID, exists := ctx.Get("merchant_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	// 获取页面列表
	pages, err := c.pageService.GetMerchantPages(merchantID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取页面列表失败"})
		return
	}

	ctx.JSON(http.StatusOK, pages)
}

// GetPage 获取单个页面
func (c *PageController) GetPage(ctx *gin.Context) {
	// 获取当前商户ID
	merchantID, exists := ctx.Get("merchant_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	// 获取页面ID
	pageIDStr := ctx.Param("id")
	pageID, err := strconv.ParseUint(pageIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的页面ID"})
		return
	}

	// 获取页面
	page, err := c.pageService.GetPageByID(uint(pageID), merchantID.(uint))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "页面不存在"})
		return
	}

	ctx.JSON(http.StatusOK, page)
}

// SavePage 保存页面
func (c *PageController) SavePage(ctx *gin.Context) {
	// 获取当前商户ID
	merchantID, exists := ctx.Get("merchant_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}

	// 获取页面ID
	pageIDStr := ctx.Param("id")
	pageID, err := strconv.ParseUint(pageIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的页面ID"})
		return
	}

	// 解析请求
	var request struct {
		Name     string      `json:"name" binding:"required"`
		Elements interface{} `json:"elements" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 保存页面
	err = c.pageService.SavePage(uint(pageID), merchantID.(uint), request.Name, request.Elements)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "保存页面失败"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "页面保存成功"})
}

// 注册路由
func (c *PageController) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/pages", c.CreatePage)
	r.GET("/pages", c.GetMerchantPages)
	r.GET("/pages/:id", c.GetPage)
	r.POST("/pages/:id/save", c.SavePage)
}
