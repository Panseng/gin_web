package middleware

import (
	"errors"
	"gin_web/utils"
	"gin_web/utils/status_msg"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(utils.JwtKey),
	}
}
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// CreateToken 生成token
func (j *JWT)CreateToken(claims MyClaims)(string, error)  {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

func (j *JWT)ParserToken(tokenString string)(*MyClaims, error)  {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})

	if err != nil {
		return nil, err
	}
	if token != nil{
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid{
			return claims, nil
		}
	}
	return  nil, errors.New("这不是一个token,请重新登录")
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = status_msg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": status_msg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			code = status_msg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": status_msg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			code = status_msg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": status_msg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		j := NewJWT()
		// 解析token
		claims, err := j.ParserToken(checkToken[1])
		if err != nil {
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"status":  status_msg.ERROR,
				"message": err.Error(),
				"data":    nil,
			})
			c.Abort()
			return
		}

		c.Set("username", claims)
		c.Next()
	}
}