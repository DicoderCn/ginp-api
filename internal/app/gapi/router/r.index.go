package router

import (
	"ginp-api/internal/app/gapi/controller"

	"github.com/DicoderCn/ginp"
)

func init() {
	// this is view
	ginp.RouterAppend(ginp.RouterItem{
		Path:      "/",                                        //api路径
		Handlers:  ginp.RegisterHandler(controller.IndexView), //对应控制器
		HttpType:  ginp.HttpGet,                               //http请求类型
		NeedLogin: false,                                      //是否需要登录
	})
}
