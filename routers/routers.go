package routers

import (
	"gin_templete/conf"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	//logger中间件
	router.Use(gin.Logger())
	//recover中间件
	router.Use(gin.Recovery())

	//配置全局的模板文件路径
	router.LoadHTMLGlob(conf.ProjectCfg.TemplateGlob)
	return router
}
