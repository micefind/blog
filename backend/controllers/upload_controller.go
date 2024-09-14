package controllers

import (
	"backend/utils"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

// UploadImage 处理图片上传并压缩
func UploadImage(c *gin.Context) {
	// 1. 获取上传的文件头信息
	header, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("获取文件失败: %v", err)})
		return
	}

	// 2. 限制文件大小不超过20MB
	const maxUploadSize = 20 * 1024 * 1024 // 20MB
	if header.Size > maxUploadSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件大小不能超过20MB"})
		return
	}

	// 3. 检查文件扩展名，仅允许 jpg、jpeg、png、gif
	fileExt := strings.ToLower(filepath.Ext(header.Filename))
	if fileExt != ".jpg" && fileExt != ".jpeg" && fileExt != ".png" && fileExt != ".gif" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "仅支持 jpg, jpeg, png, gif 格式的图片"})
		return
	}

	// 4. 生成唯一文件名，避免文件名冲突
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt)

	// 5. 指定文件保存路径
	savePath := filepath.Join("./static/images", fileName)

	// 6. 创建保存路径（如果不存在）
	if err := os.MkdirAll("./static/images", os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("创建文件夹失败: %v", err)})
		return
	}

	// 7. 打开上传的文件进行解码
	file, err := header.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("无法打开文件: %v", err)})
		return
	}
	defer file.Close()

	var img image.Image
	switch fileExt {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JPEG 图片解码失败"})
			return
		}
	case ".png":
		img, err = png.Decode(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "PNG 图片解码失败"})
			return
		}
	case ".gif":
		img, err = gif.Decode(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "GIF 图片解码失败"})
			return
		}
	}

	// 8. 压缩图片，保持图片宽度为800像素
	newImage := resize.Resize(800, 0, img, resize.Lanczos3)

	// 9. 保存压缩后的图片到服务器
	outFile, err := os.Create(savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("保存文件失败: %v", err)})
		return
	}
	defer outFile.Close()

	switch fileExt {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(outFile, newImage, &jpeg.Options{Quality: 80}) // 压缩JPEG
	case ".png":
		err = png.Encode(outFile, newImage) // PNG 无损压缩
	case ".gif":
		err = gif.Encode(outFile, newImage, nil) // GIF 无损压缩
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("图片压缩失败: %v", err)})
		return
	}

	// 10. 构建文件的完整URL，包括协议、域名和端口号
	protocol := "http"
	if c.Request.TLS != nil {
		protocol = "https"
	}
	host := c.Request.Host // 获取域名和端口号
	fileURL := fmt.Sprintf("%s://%s/static/images/%s", protocol, host, fileName)

	// 11. 返回文件的完整URL，供客户端访问
	utils.JSONResponse(c, http.StatusOK, "文件上传并压缩成功", gin.H{"url": fileURL})
}
