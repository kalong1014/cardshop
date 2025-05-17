package controller

import (
    "io"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "time"

    "card-system/internal/common"
    "github.com/gin-gonic/gin"
)

// FileController 文件控制器
type FileController struct {
    logger    common.Logger
    uploadDir string
}

// NewFileController 创建文件控制器
func NewFileController(logger common.Logger, uploadDir string) *FileController {
    // 确保上传目录存在
    if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
        os.MkdirAll(uploadDir, 0755)
    }
    
    return &FileController{
        logger:    logger,
        uploadDir: uploadDir,
    }
}

// UploadFile 处理文件上传
func (fc *FileController) UploadFile(c *gin.Context) {
    // 单文件上传
    file, err := c.FormFile("file")
    if err != nil {
        fc.logger.Error("获取上传文件失败: %v", err)
        c.JSON(http.StatusBadRequest, common.NewAPIError(http.StatusBadRequest, "获取上传文件失败"))
        return
    }
    
    // 生成唯一文件名
    ext := filepath.Ext(file.Filename)
    timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
    filename := timestamp + ext
    
    // 保存文件
    filePath := filepath.Join(fc.uploadDir, filename)
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        fc.logger.Error("保存文件失败: %v", err)
        c.JSON(http.StatusInternalServerError, common.NewAPIError(http.StatusInternalServerError, "保存文件失败"))
        return
    }
    
    // 返回文件URL
    fileURL := "/uploads/" + filename
    c.JSON(http.StatusOK, gin.H{
        "code": 200,
        "message": "上传成功",
        "data": gin.H{
            "url": fileURL,
        },
    })
}

// DownloadFile 处理文件下载
func (fc *FileController) DownloadFile(c *gin.Context) {
    filename := c.Param("filename")
    filePath := filepath.Join(fc.uploadDir, filename)
    
    // 检查文件是否存在
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        fc.logger.Warn("请求的文件不存在: %s", filePath)
        c.JSON(http.StatusNotFound, common.NewAPIError(http.StatusNotFound, "文件不存在"))
        return
    }
    
    // 设置响应头
    c.Header("Content-Disposition", "attachment; filename="+filename)
    c.Header("Content-Type", "application/octet-stream")
    
    // 发送文件
    c.File(filePath)
}

// DeleteFile 处理文件删除
func (fc *FileController) DeleteFile(c *gin.Context) {
    filename := c.Param("filename")
    filePath := filepath.Join(fc.uploadDir, filename)
    
    // 检查文件是否存在
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        fc.logger.Warn("请求删除的文件不存在: %s", filePath)
        c.JSON(http.StatusNotFound, common.NewAPIError(http.StatusNotFound, "文件不存在"))
        return
    }
    
    // 删除文件
    if err := os.Remove(filePath); err != nil {
        fc.logger.Error("删除文件失败: %v", err)
        c.JSON(http.StatusInternalServerError, common.NewAPIError(http.StatusInternalServerError, "删除文件失败"))
        return
    }
    
    c.JSON(http.StatusOK, common.NewAPIError(http.StatusOK, "文件已删除"))
}

// ListFiles 列出所有上传的文件
func (fc *FileController) ListFiles(c *gin.Context) {
    files := make([]map[string]interface{}, 0)
    
    // 读取目录内容
    dirEntries, err := os.ReadDir(fc.uploadDir)
    if err != nil {
        fc.logger.Error("读取上传目录失败: %v", err)
        c.JSON(http.StatusInternalServerError, common.NewAPIError(http.StatusInternalServerError, "读取文件列表失败"))
        return
    }
    
    // 构建文件列表
    for _, entry := range dirEntries {
        if !entry.IsDir() {
            fileInfo, err := entry.Info()
            if err != nil {
                fc.logger.Warn("获取文件信息失败: %v", err)
                continue
            }
            
            files = append(files, map[string]interface{}{
                "name":    entry.Name(),
                "size":    fileInfo.Size(),
                "modTime": fileInfo.ModTime(),
                "url":     "/uploads/" + entry.Name(),
            })
        }
    }
    
    c.JSON(http.StatusOK, gin.H{
        "code": 200,
        "message": "获取文件列表成功",
        "data": files,
    })
}    