package v1

import (
	"gin_web/model"
	"gin_web/utils/status_msg"
	"gin_web/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	var validCode int
	_ = c.ShouldBindJSON(&data)

	msg, validCode = validator.Validate(&data)
	if validCode != statusMsg.SUCCSE {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  validCode,
				"message": msg,
			},
		)
		c.Abort()
		return
	}

	code := model.CheckUserName(data.Username)
	if code == statusMsg.SUCCSE {
		model.CreateUser(&data)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": statusMsg.GetStatusMsg(code),
		},
	)
}

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetUser(id)
	//var maps = make(map[string]interface{})
	//maps["username"] = data.Username
	//maps["role"] = data.Role

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   1,
			"message": statusMsg.GetStatusMsg(code),
		},
	)

}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.CheckUpUser(id, data.Username)
	if code == statusMsg.SUCCSE {
		model.EditUser(id, &data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": statusMsg.GetStatusMsg(code),
		},
	)
}

// ChangeUserPassword 修改密码
func ChangeUserPassword(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.ChangePassword(id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": statusMsg.GetStatusMsg(code),
		},
	)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": statusMsg.GetStatusMsg(code),
	},
	)
}

// GetUserList 查询用户列表
func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.GetUserList(username, pageSize, pageNum)

	code := statusMsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": statusMsg.GetStatusMsg(code),
		},
	)
}

// GetManagerUserCounts 查询管理员总数
func GetManagerUserCounts(c *gin.Context) {
	count := model.GetManagerUserCount()
	c.JSON(http.StatusOK, gin.H{
		"status": statusMsg.SUCCSE,
		"data":  map[string]int64{"count": count},
		"message": statusMsg.GetStatusMsg(statusMsg.SUCCSE),
	})
}
