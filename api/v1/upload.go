package v1

import (
	"gin_web/model"
	statusMsg "gin_web/utils/status_msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := model.UploadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": statusMsg.GetStatusMsg(code),
		"url":     url,
	})
}
