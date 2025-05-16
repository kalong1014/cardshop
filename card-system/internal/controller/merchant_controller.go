package controller

import (
	"net/http"

	"card-system/internal/config"
	"card-system/internal/model"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	cfg *config.Config
}

func (c *MerchantController) RegisterRoutes(r *gin.RouterGroup) {
	// 注册商户相关路由
}

func NewMerchantController(cfg *config.Config) *MerchantController {
	return &MerchantController{cfg: cfg}
}

type CreateMerchantRequest struct {
	Name   string `json:"name" binding:"required"`
	Logo   string `json:"logo"`
	Domain string `json:"domain" binding:"required"`
}

func (ctrl *MerchantController) CreateMerchant(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	var req CreateMerchantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查域名是否已存在
	var existingMerchant model.Merchant
	if err := model.DB.Where("domain = ?", req.Domain).First(&existingMerchant).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Domain already exists"})
		return
	}

	// 创建商户
	merchant := model.Merchant{
		UserID: userID.(uint),
		Name:   req.Name,
		Logo:   req.Logo,
		Status: "pending",
		Level:  "basic",
		Domain: req.Domain,
	}

	if err := model.DB.Create(&merchant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create merchant"})
		return
	}

	c.JSON(http.StatusCreated, merchant)
}

func (ctrl *MerchantController) GetAllMerchants(c *gin.Context) {
	var merchants []model.Merchant
	if err := model.DB.Preload("User").Find(&merchants).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch merchants"})
		return
	}

	c.JSON(http.StatusOK, merchants)
}

type UpdateMerchantStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

func (ctrl *MerchantController) UpdateMerchantStatus(c *gin.Context) {
	merchantID := c.Param("id")

	var merchant model.Merchant
	if err := model.DB.First(&merchant, merchantID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Merchant not found"})
		return
	}

	var req UpdateMerchantStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	merchant.Status = req.Status
	if err := model.DB.Save(&merchant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update merchant status"})
		return
	}

	c.JSON(http.StatusOK, merchant)
}
