package controller

import "github.com/DicoderCn/ginp"

func Index(c *ginp.ContextPlus) {
	c.Success("ok")
}
