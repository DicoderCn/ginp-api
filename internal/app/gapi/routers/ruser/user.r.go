package ruser

import (
	"ginp-api/internal/app/gapi/controller/cuser"
	"ginp-api/internal/app/gapi/dto/comdto"
	"ginp-api/internal/app/gapi/entity"

	"github.com/DicoderCn/ginp"
)

const (
	ApiUserCreate   = "/api/user/create"
	ApiUserFindById = "/api/user/findById"
	ApiUserSearch   = "/api/user/search"
	ApiUserUpdate   = "/api/user/update"
	ApiUserDelete   = "/api/user/delete"
)

// this is router define file
func InitRouter() {

	// Create
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiUserCreate,                      //api路径
		Handlers:       ginp.RegisterHandler(cuser.Create), //对应控制器
		HttpType:       ginp.HttpPost,                      //http请求类型
		NeedLogin:      false,                              //是否需要登录
		NeedPermission: false,                              //是否需要鉴权
		PermissionName: "User.create",                      //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "create user",
			Description: "",
			RequestDto:  entity.User{},
		},
	})

	// FindById
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiUserFindById,                      //api路径
		Handlers:       ginp.RegisterHandler(cuser.FindByID), //对应控制器
		HttpType:       ginp.HttpPost,                        //http请求类型
		NeedLogin:      false,                                //是否需要登录
		NeedPermission: false,                                //是否需要鉴权
		PermissionName: "User.findById",                      //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "find user by id",
			Description: "",
			RequestDto:  entity.User{},
		},
	})

	// 修改
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiUserUpdate,                      //api路径
		Handlers:       ginp.RegisterHandler(cuser.Update), //对应控制器
		HttpType:       ginp.HttpPost,                      //http请求类型
		NeedLogin:      true,                               //是否需要登录
		NeedPermission: true,                               //是否需要鉴权
		PermissionName: "User.update",                      //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "modify user",
			Description: "",
			RequestDto:  entity.User{},
		},
	})

	// 删除
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiUserDelete,                      //api路径
		Handlers:       ginp.RegisterHandler(cuser.Delete), //对应控制器
		HttpType:       ginp.HttpPost,                      //http请求类型
		NeedLogin:      true,                               //是否需要登录
		NeedPermission: true,                               //是否需要鉴权
		PermissionName: "User.delete",                      //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "delet user",
			Description: "",
			RequestDto:  entity.User{},
		},
	})

	// search 搜索
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiUserSearch,                      //api路径
		Handlers:       ginp.RegisterHandler(cuser.Search), //对应控制器
		HttpType:       ginp.HttpPost,                      //http请求类型
		NeedLogin:      true,                               //是否需要登录
		NeedPermission: true,                               //是否需要鉴权
		PermissionName: "User.search",                      //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "search user",
			Description: "",
			RequestDto:  comdto.ReqSearch{},
		},
	})

}
