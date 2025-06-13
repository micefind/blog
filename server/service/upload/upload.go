package upload

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	model "server/model/upload"
	"strings"
	"time"
)

func UploadImage(file *multipart.FileHeader) (fileInfo model.File, err error) {
	// 获取文件信息
	fileInfo.FileName = file.Filename
	fileInfo.FileSize = fmt.Sprintf("%.2f MB", float64(file.Size)/(1024*1024))
	fileInfo.FileType = file.Header.Get("Content-Type")

	// 检查文件格式
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !strings.Contains(".jpg.jpeg.png.gif", ext) {
		return fileInfo, errors.New("仅支持 jpg, jpeg, png, gif 格式的图片")
	}
	// 检查文件大小
	if file.Size > 1024*1024*20 {
		return fileInfo, errors.New("文件大小不能超过20MB")
	}

	// 使用时间戳生成唯一文件名
	uniqueName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	
	// 创建保存目录
	saveDir := "static/image"
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		if err := os.MkdirAll(saveDir, 0755); err != nil {
			return fileInfo, fmt.Errorf("创建目录失败: %v", err)
		}
	}
	
	// 完整文件路径
	filePath := filepath.Join(saveDir, uniqueName)
	fileInfo.FilePath = filePath
	
	// 保存文件
	src, err := file.Open()
	if err != nil {
		return fileInfo, fmt.Errorf("打开上传文件失败: %v", err)
	}
	defer src.Close()
	
	dst, err := os.Create(filePath)
	if err != nil {
		return fileInfo, fmt.Errorf("创建保存文件失败: %v", err)
	}
	defer dst.Close()
	
	if _, err := io.Copy(dst, src); err != nil {
		return fileInfo, fmt.Errorf("保存文件失败: %v", err)
	}
	
	return fileInfo, nil
}

func UploadFile(file *multipart.FileHeader) (fileInfo model.File, err error) { 
	// 获取文件信息
	fileInfo.FileName = file.Filename
	fileInfo.FileSize = fmt.Sprintf("%.2f MB", float64(file.Size)/(1024*1024))
	fileInfo.FileType = file.Header.Get("Content-Type")

	// 检查文件大小
	if file.Size > 1024*1024*100 {
		return fileInfo, errors.New("文件大小不能超过100MB")
	}

	// 使用时间戳生成唯一文件名
	ext := strings.ToLower(filepath.Ext(file.Filename))
	uniqueName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	
	// 创建保存目录
	saveDir := "static/file"
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		if err := os.MkdirAll(saveDir, 0755); err != nil {
			return fileInfo, fmt.Errorf("创建目录失败: %v", err)
		}
	}
	
	// 完整文件路径
	filePath := filepath.Join(saveDir, uniqueName)
	fileInfo.FilePath = filePath
	
	// 保存文件
	src, err := file.Open()
	if err != nil {
		return fileInfo, fmt.Errorf("打开上传文件失败: %v", err)
	}
	defer src.Close()
	
	dst, err := os.Create(filePath)
	if err != nil {
		return fileInfo, fmt.Errorf("创建保存文件失败: %v", err)
	}
	defer dst.Close()
	
	if _, err := io.Copy(dst, src); err != nil {
		return fileInfo, fmt.Errorf("保存文件失败: %v", err)
	}
	
	return fileInfo, nil
}