package configs

import "ginp-api/pkg/cfg"

const (
	ServerPort = "server.port"
)

// 初始化配置
func init() {
	cfg.SetDefault(ServerPort, ":8082")
}
