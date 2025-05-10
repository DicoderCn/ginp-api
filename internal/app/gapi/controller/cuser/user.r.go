package cuser

import "github.com/DicoderCn/ginp"

// this is router define file
func init() {
	ginp.RouterAppend(ginp.RouterItem{
		Path:     "/api/",
		HttpType: ginp.HttpGet,
		Handlers: ginp.RegisterHandler(Create),
	})
}
