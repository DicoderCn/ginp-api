package router

import (
	"ginp-api/internal/app/gapi/controller/cuser_group"
	"ginp-api/internal/app/gapi/dto/comdto"
	"ginp-api/internal/app/gapi/entity"

	"github.com/DicoderCn/ginp"
)

const (
	ApiUserGroupCreate   = "/api/user_group/create"
	ApiUserGroupFindById = "/api/user_group/findById"
	ApiUserGroupSearch   = "/api/user_group/search"
	ApiUserGroupUpdate   = "/api/user_group/update"
	ApiUserGroupDelete   = "/api/user_group/delete"
)

// this is router define file
func init() {

	// Create
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiUserGroupCreate,                           //api路径
		Handlers:       ginp.RegisterHandler(cuser_group.Create), //对应控制器
		HttpType:       ginp.HttpPost,                           //http请求类型
		NeedLogin:      false,                                   //是否需要登录
		NeedPermission: false,                                   //是否需要鉴权
		PermissionName: "UserGroup.create",                           //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "create user_group",
			Description: "",
			RequestDto:  entity.UserGroup{},
		},
	})

	// FindById
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiUserGroupFindById,                           //api路径
		Handlers:       ginp.RegisterHandler(cuser_group.FindByID), //对应控制器
		HttpType:       ginp.HttpPost,                             //http请求类型
		NeedLogin:      false,                                     //是否需要登录
		NeedPermission: false,                                     //是否需要鉴权
		PermissionName: "UserGroup.findById",                           //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "find user_group by id",
			Description: "",
			RequestDto:  entity.UserGroup{},
		},
	})

	// 修改
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiUserGroupUpdate,                           //api路径
		Handlers:       ginp.RegisterHandler(cuser_group.Update), //对应控制器
		HttpType:       ginp.HttpPost,                           //http请求类型
		NeedLogin:      true,                                    //是否需要登录
		NeedPermission: true,                                    //是否需要鉴权
		PermissionName: "UserGroup.update",                           //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "modify user_group",
			Description: "",
			RequestDto:  entity.UserGroup{},
		},
	})

	// 删除
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiUserGroupDelete,                           //api路径
		Handlers:       ginp.RegisterHandler(cuser_group.Delete), //对应控制器
		HttpType:       ginp.HttpPost,                           //http请求类型
		NeedLogin:      true,                                    //是否需要登录
		NeedPermission: true,                                    //是否需要鉴权
		PermissionName: "UserGroup.delete",                           //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "delet user_group",
			Description: "",
			RequestDto:  entity.UserGroup{},
		},
	})

	// search 搜索
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiUserGroupSearch,                           //api路径
		Handlers:       ginp.RegisterHandler(cuser_group.Search), //对应控制器
		HttpType:       ginp.HttpPost,                           //http请求类型
		NeedLogin:      true,                                    //是否需要登录
		NeedPermission: true,                                    //是否需要鉴权
		PermissionName: "UserGroup.search",                           //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "search user_group",
			Description: "",
			RequestDto:  comdto.ReqSearch{},
		},
	})

}
