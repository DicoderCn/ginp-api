package controller

import "github.com/DicoderCn/ginp"

func init() {
	// 注册路由
	// ginp.RegisterRouter("/", Index)
}

func Index(c *ginp.ContextPlus) {
	c.Success("ok")
}
