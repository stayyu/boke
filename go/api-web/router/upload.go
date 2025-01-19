package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func handleUpload(c *gin.Context) {
	fmt.Print(c.Request)
	file, header, err := c.Request.FormFile("files")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": 1,
			"message": "文件读取失败！"})
		fmt.Println("文件读取失败！")
		return
	}
	defer file.Close()
	filePath := fmt.Sprintf("%s/%s", "./go/static/img", header.Filename)
	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": 1,
			"message": "上传读取失败！"})
		fmt.Println("上传读取失败！")
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": 1,
			"message": "上传复制失败！"})
		fmt.Println("上传复制失败！")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"error": 0,
		"data": gin.H{"URL": filePath}})
	fmt.Println("ok")

}
func handleFile(c *gin.Context) {
	fileName := c.Param("filename")
	filePath := fmt.Sprintf("%s/%s", "./go/static/img", fileName)

	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": 1,
			"message": "文件读取失败！"})
		return
	}

	defer file.Close()

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Type", "application/octet-stream") // Set Content-Type to audio/mpeg
	io.Copy(c.Writer, file)
}
