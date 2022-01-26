package v1

import (
	"gin_web/middleware"
	"gin_web/model"
	"gin_web/utils/status_msg"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 后台登录
func Login(c *gin.Context)  {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var token string
	var code int
	formData, code = model.CheckLogin(formData.Username, formData.Password)

	if code == statusMsg.SUCCSE{
		setToken(c, formData)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    formData.Username,
			"id":      formData.ID,
			"message": statusMsg.GetErrMsg(code),
			"token":   token,
		})
	}
}

// LoginFront 前台登录
func LoginFront(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var code int

	formData, code = model.CheckLoginFront(formData.Username, formData.Password)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    formData.Username,
		"id":      formData.ID,
		"message": statusMsg.GetErrMsg(code),
	})
}

// token生成函数
func setToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			//NotBefore: time.Now()- 100,
			NotBefore: jwt.At(time.Unix(time.Now().Unix()-100, 0)), // 时间戳转换为时间
			//ExpiresAt: time.Now().Unix() + 604800,
			ExpiresAt: jwt.At(time.Unix(time.Now().Unix() + 604800, 0)),
			Issuer:    "GinWeb",
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  statusMsg.ERROR,
			"message": statusMsg.GetErrMsg(statusMsg.ERROR),
			"token":   token,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    user.Username,
		"id":      user.ID,
		"message": statusMsg.GetErrMsg(200),
		"token":   token,
	})
	return
}
