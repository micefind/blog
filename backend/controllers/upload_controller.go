package controllers

import (
	"backend/utils"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// formatFileSize 格式化文件大小
func formatFileSize(size int64) string {
	if size > 1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(size)/1024/1024)
	}
	return fmt.Sprintf("%.2f KB", float64(size)/1024)
}

// getImageDimensions 获取图片分辨率
func getImageDimensions(img image.Image) (int, int) {
	return img.Bounds().Dx(), img.Bounds().Dy()
}

// getProtocol 获取请求协议 (http/https)
func getProtocol(c *gin.Context) string {
	if c.Request.TLS != nil {
		return "https"
	}
	return "http"
}

// saveAndRespond 保存文件并响应
func saveAndRespond(c *gin.Context, header *multipart.FileHeader, fileName, savePath string, err error) {
	if saveErr := c.SaveUploadedFile(header, savePath); saveErr != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("保存文件失败: %v", saveErr), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, fmt.Sprintf("图片处理失败，但已保存原图: %v", err), gin.H{
		"url":       fmt.Sprintf("%s://%s/static/images/%s", getProtocol(c), c.Request.Host, fileName),
		"file_name": fileName,
	})
}

// UploadImage 处理图片上传并压缩
func UploadImage(c *gin.Context) {
	header, err := c.FormFile("file")
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("获取文件失败: %v", err), nil)
		return
	}

	const maxUploadSize = 20 * 1024 * 1024 // 20MB
	if header.Size > maxUploadSize {
		utils.JSONResponse(c, http.StatusBadRequest, "文件大小不能超过20MB", nil)
		return
	}

	fileExt := strings.ToLower(filepath.Ext(header.Filename))
	if !strings.Contains(".jpg.jpeg.png.gif", fileExt) {
		utils.JSONResponse(c, http.StatusBadRequest, "仅支持 jpg, jpeg, png, gif 格式的图片", nil)
		return
	}

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt)
	savePath := filepath.Join("./static/images", fileName)
	if err := os.MkdirAll(filepath.Dir(savePath), os.ModePerm); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("创建文件夹失败: %v", err), nil)
		return
	}

	file, err := header.Open()
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无法打开文件: %v", err), nil)
		return
	}
	defer file.Close()

	var img image.Image
	shouldCompress := true

	switch fileExt {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	case ".gif":
		img, err = gif.Decode(file)
		shouldCompress = false
	}

	if err != nil {
		saveAndRespond(c, header, fileName, savePath, err)
		return
	}

	origWidth, origHeight := getImageDimensions(img)
	origSize := header.Size

	outFile, err := os.Create(savePath)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("保存文件失败: %v", err), nil)
		return
	}
	defer outFile.Close()

	if shouldCompress {
		switch fileExt {
		case ".jpg", ".jpeg":
			err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 80})
		case ".png":
			err = png.Encode(outFile, img)
		}

		if err != nil {
			saveAndRespond(c, header, fileName, savePath, err)
			return
		}
	} else {
		saveAndRespond(c, header, fileName, savePath, nil)
		return
	}

	compressedFileInfo, err := outFile.Stat()
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("获取压缩文件信息失败: %v", err), nil)
		return
	}

	compressedSize := compressedFileInfo.Size()
	compressedWidth, compressedHeight := getImageDimensions(img)

	fileURL := fmt.Sprintf("%s://%s/static/images/%s", getProtocol(c), c.Request.Host, fileName)

	utils.JSONResponse(c, http.StatusOK, "文件上传并压缩成功", gin.H{
		"url":                   fileURL,
		"file_name":             fileName,
		"original_resolution":   fmt.Sprintf("%dx%d", origWidth, origHeight),
		"original_size":         formatFileSize(origSize),
		"compressed_resolution": fmt.Sprintf("%dx%d", compressedWidth, compressedHeight),
		"compressed_size":       formatFileSize(compressedSize),
	})
}
