package cuser

import (
	"github.com/DicoderCn/ginp"
)

func LoginByEmail(c *ginp.ContextPlus) {
	var requestParams *RequestLoginByEmail
	if err := c.ShouldBindJSON(&requestParams); err != nil {
		c.FailData("request param error:" + err.Error())
		return
	}

	//TODO:
	//c.SuccessData(&RespondLoginByEmail{})
}

const ApiLoginByEmail = "/api/user/login_by_email" //API Path

type RequestLoginByEmail struct {
}

type RespondLoginByEmail struct {
}

func init() {
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiLoginByEmail,              //api路径
		Handlers:       ginp.RegisterHandler(Create), //对应控制器
		HttpType:       ginp.HttpPost,                //http请求类型
		NeedLogin:      false,                        //是否需要登录
		NeedPermission: false,                        //是否需要鉴权
		PermissionName: "User.login_by_email",        //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "login_by_email",
			Description: "",
			RequestDto:  RequestLoginByEmail{},
		},
	})
}
