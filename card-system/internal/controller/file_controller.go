package controller

import (
	"card-system/internal/model"
	"card-system/internal/utils" // 添加此行
	"card-system/pkg/storage"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// controller/file_controller.go
func UploadLogo(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件上传失败"})
		return
	}

	merchantID := getCurrentMerchantID(c)
	filename := fmt.Sprintf("merchant_%d_logo_%s", merchantID, uuid.New().String())
	dst := fmt.Sprintf("static/logos/%s.jpg", filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		handleError(c, err)
		return
	}

	// 更新商户LOGO路径
	if err := utils.DB.Model(&model.Merchant{}).Where("id = ?", merchantID).Update("logo", dst).Error; err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": "/static/logos/" + filename + ".jpg"})
}

// 控制器集成（支持多存储切换）
func UploadLogo(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件读取失败"})
		return
	}

	// 动态获取存储驱动（从系统配置）
	driver, _ := service.Config.Get("storage.driver")
	var storage storage.Storage
	switch driver {
	case "local":
		storage = &storage.LocalStorage{BasePath: "static/logos/"}
	case "minio":
		storage = &storage.MinIOStorage{
			Client: minioClient, // 初始化MinIO客户端
			Bucket: "merchant-logos",
		}
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "不支持的存储驱动"})
		return
	}

	filename := fmt.Sprintf("merchant_%d_logo_%s.jpg", getCurrentMerchantID(c), uuid.New().String())
	if err := storage.Upload(context.Background(), filename, file.Open()); err != nil {
		return handleError(c, err)
	}

	// 保存文件路径到数据库
	if err := utils.DB.Model(&model.Merchant{}).Where("id = ?", getCurrentMerchantID(c)).Update("logo", filename).Error; err != nil {
		return handleError(c, err)
	}

	c.JSON(http.StatusOK, gin.H{"url": storage.GetURL(context.Background(), filename)})
}
