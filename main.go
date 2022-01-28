package main

import (
	"gin_web/model"
	"gin_web/router"
)

func main()  {
	// 初始化数据库
	model.InitDB()

	//// 引入路由组件
	router.InitRouter()
}
