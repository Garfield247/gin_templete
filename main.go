package main

import (
	"gin_templete/conf"
	"gin_templete/models"
	"gin_templete/routers"
)

//初始化方法
func init() {
	//加载配置文件
	conf.Setup()
	//初始化数据库连接
	models.SetUp()
}

func main() {
	router := routers.InitRouter()
	panic(router.Run())
}
