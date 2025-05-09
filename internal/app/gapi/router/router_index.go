package router

import (
	"ginp-api/internal/app/gapi/controller"

	"github.com/DicoderCn/ginp"
)

func init() {
	ginp.RouterAppend(ginp.RouterItem{
		Path:     "/",
		HttpType: ginp.HttpGet,
		Handlers: ginp.RegisterHandler(controller.Index),
	})
}
