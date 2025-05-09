package start

import (
	"ginp-api/configs"
	"ginp-api/internal/app/gapi/router"
	"ginp-api/pkg/cfg"

	"github.com/gin-gonic/gin"
)

// 程序入口
func Run() {

	//配置路径已定死在 cmd/gapi/configs.yaml ,要修改只能去改cfg包
	//由于cionfigs包调用了init()初始化 因此使用cfg.initCfg函数可能会无效
	port := cfg.GetString(configs.ServerPort)
	r := gin.Default()
	router.Register(r)
	println("start server on port: " + port)
	r.Run(":" + port)
}
