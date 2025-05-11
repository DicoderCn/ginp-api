package router

import (
	"ginp-api/internal/app/gapi/controller/cindex"

	"github.com/DicoderCn/ginp"
)

func InitRouter() {
	ginp.RouterAppend(ginp.RouterItem{
		Path:      "/",                                    //api路径
		Handlers:  ginp.RegisterHandler(cindex.IndexView), //对应控制器
		HttpType:  ginp.HttpGet,                           //http请求类型
		NeedLogin: false,                                  //是否需要登录
	})
}
