package api

import (
	"daidai-server/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

func UploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	file, err := fileHeader.Open()
	defer file.Close()
	uuid := uuid.NewV4()

	path := "upload/" + uuid.String() + fileHeader.Filename
	url, err := util.UpLoadFile(path, file)
	if err != nil {
		log.Println("上传文件失败: ", err)

		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "上传失败",
			"data": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "上传成功",
		"data": url,
	})

}
