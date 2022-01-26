package main

import (
	"gin_web/model"
	"gin_web/router"
)

func main()  {
	model.InitDB()
	// 引入路由组件
	router.InitRouter()
}
