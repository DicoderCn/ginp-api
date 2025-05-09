package start

import (
	"ginp-api/configs"
	"ginp-api/pkg/cfg"

	"github.com/gin-gonic/gin"
)

// 程序入口
func Run() {
	r := gin.Default()
	cfg.InitCfg(ConfigFile)
	r.Run(cfg.GetString(configs.ServerPort))
}
