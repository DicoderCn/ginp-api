package router

import (
	"ginp-api/internal/app/gapi/controller"
	"ginp-api/internal/app/gapi/dto/comdto"
	"ginp-api/internal/app/gapi/entity"

	"github.com/DicoderCn/ginp"
)

const (
	ApiFindById = "/api/user/findById"
	ApiSearch   = "/api/user/search"
	ApiCreate   = "/api/user/create"
	ApiUpdate   = "/api/user/update"
	ApiDelete   = "/api/user/delete"
)

// this is router define file
func init() {
	// FindById
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiFindById,                               //api路径
		Handlers:       ginp.RegisterHandler(controller.FindByID), //对应控制器
		HttpType:       ginp.HttpPost,                             //http请求类型
		NeedLogin:      false,                                     //是否需要登录
		NeedPermission: false,                                     //是否需要鉴权
		PermissionName: "User.findById",                           //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "find user by id",
			Description: "",
			RequestDto:  entity.User{},
		},
	})

	// Create
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiCreate,                               //api路径
		Handlers:       ginp.RegisterHandler(controller.Create), //对应控制器
		HttpType:       ginp.HttpPost,                           //http请求类型
		NeedLogin:      false,                                   //是否需要登录
		NeedPermission: false,                                   //是否需要鉴权
		PermissionName: "User.create",                           //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "create user",
			Description: "",
			RequestDto:  entity.User{},
		},
	})

	// 修改
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiUpdate,                               //api路径
		Handlers:       ginp.RegisterHandler(controller.Update), //对应控制器
		HttpType:       ginp.HttpPost,                           //http请求类型
		NeedLogin:      true,                                    //是否需要登录
		NeedPermission: true,                                    //是否需要鉴权
		PermissionName: "User.update",                           //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "modify user",
			Description: "",
			RequestDto:  entity.User{},
		},
	})

	// 删除
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiDelete,                               //api路径
		Handlers:       ginp.RegisterHandler(controller.Delete), //对应控制器
		HttpType:       ginp.HttpPost,                           //http请求类型
		NeedLogin:      true,                                    //是否需要登录
		NeedPermission: true,                                    //是否需要鉴权
		PermissionName: "User.delete",                           //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "delet user",
			Description: "",
			RequestDto:  entity.User{},
		},
	})

	// search 搜索
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiSearch,                               //api路径
		Handlers:       ginp.RegisterHandler(controller.Search), //对应控制器
		HttpType:       ginp.HttpPost,                           //http请求类型
		NeedLogin:      true,                                    //是否需要登录
		NeedPermission: true,                                    //是否需要鉴权
		PermissionName: "User.search",                           //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "search user",
			Description: "",
			RequestDto:  comdto.ReqSearch{},
		},
	})

}
