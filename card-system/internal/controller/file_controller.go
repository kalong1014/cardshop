package controller

import (
	"net/http"
	"strconv"

	"card-system/internal/middleware"
	"card-system/internal/service"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	fileService service.FileService
}

func NewFileController(fileService service.FileService) *FileController {
	return &FileController{fileService: fileService}
}

func (fc *FileController) SetupRoutes(router *gin.RouterGroup) {
	fileGroup := router.Group("/files")
	{
		fileGroup.GET("/:id", fc.GetFile)
		fileGroup.POST("/upload", middleware.RateLimit(), fc.UploadFile)
		fileGroup.DELETE("/:id", fc.DeleteFile)
	}
}

func (fc *FileController) GetFile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文件ID"})
		return
	}

	file, err := fc.fileService.GetFileByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文件不存在"})
		return
	}

	c.JSON(http.StatusOK, file)
}

func (fc *FileController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少文件上传"})
		return
	}

	userID := c.GetInt("user_id")
	newFile, err := fc.fileService.SaveFile(file, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}

	c.JSON(http.StatusCreated, newFile)
}

func (fc *FileController) DeleteFile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文件ID"})
		return
	}

	userID := c.GetInt("user_id")
	err = fc.fileService.DeleteFile(id, userID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权限删除文件"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文件删除成功"})
}
