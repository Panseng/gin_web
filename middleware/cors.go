package middleware

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"time"
)

func Cors() gin.HandlerFunc{
	return cors.New(cors.Options{
		//AllowAllOrigins:  true,
		AllowedOrigins:     []string{"*"}, // 等同于允许所有域名 #AllowAllOrigins:  true
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"*","Authorization"},
		ExposedHeaders:    []string{"Content-Length", "text/plain", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:   int(12 * time.Hour),
	})
}