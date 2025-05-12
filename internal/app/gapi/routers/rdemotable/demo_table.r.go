package rdemotable

import (
	"ginp-api/internal/app/gapi/controller/cdemotable"
	"ginp-api/internal/app/gapi/dto/comdto"
	"ginp-api/internal/app/gapi/entity"

	"github.com/DicoderCn/ginp"
)

const (
	ApiDemoTableCreate   = "/api/demo_table/create"
	ApiDemoTableFindById = "/api/demo_table/findById"
	ApiDemoTableSearch   = "/api/demo_table/search"
	ApiDemoTableUpdate   = "/api/demo_table/update"
	ApiDemoTableDelete   = "/api/demo_table/delete"
)

// this is router define file
func init() {

	// Create
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiDemoTableCreate,                      //api路径
		Handlers:       ginp.RegisterHandler(cdemotable.Create), //对应控制器
		HttpType:       ginp.HttpPost,                           //http请求类型
		NeedLogin:      false,                                   //是否需要登录
		NeedPermission: false,                                   //是否需要鉴权
		PermissionName: "DemoTable.create",                      //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "create demo_table",
			Description: "",
			RequestDto:  entity.DemoTable{},
		},
	})

	// FindById
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiDemoTableFindById,                      //api路径
		Handlers:       ginp.RegisterHandler(cdemotable.FindByID), //对应控制器
		HttpType:       ginp.HttpPost,                             //http请求类型
		NeedLogin:      false,                                     //是否需要登录
		NeedPermission: false,                                     //是否需要鉴权
		PermissionName: "DemoTable.findById",                      //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "find demo_table by id",
			Description: "",
			RequestDto:  entity.DemoTable{},
		},
	})

	// 修改
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiDemoTableUpdate,                      //api路径
		Handlers:       ginp.RegisterHandler(cdemotable.Update), //对应控制器
		HttpType:       ginp.HttpPost,                           //http请求类型
		NeedLogin:      true,                                    //是否需要登录
		NeedPermission: true,                                    //是否需要鉴权
		PermissionName: "DemoTable.update",                      //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "modify demo_table",
			Description: "",
			RequestDto:  entity.DemoTable{},
		},
	})

	// 删除
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiDemoTableDelete,                      //api路径
		Handlers:       ginp.RegisterHandler(cdemotable.Delete), //对应控制器
		HttpType:       ginp.HttpPost,                           //http请求类型
		NeedLogin:      true,                                    //是否需要登录
		NeedPermission: true,                                    //是否需要鉴权
		PermissionName: "DemoTable.delete",                      //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "delet demo_table",
			Description: "",
			RequestDto:  entity.DemoTable{},
		},
	})

	// search 搜索
	ginp.RouterAppend(ginp.RouterItem{
		Path:           ApiDemoTableSearch,                      //api路径
		Handlers:       ginp.RegisterHandler(cdemotable.Search), //对应控制器
		HttpType:       ginp.HttpPost,                           //http请求类型
		NeedLogin:      true,                                    //是否需要登录
		NeedPermission: true,                                    //是否需要鉴权
		PermissionName: "DemoTable.search",                      //完整的权限名称,会跟权限表匹配
		Swagger: &ginp.SwaggerInfo{
			Title:       "search demo_table",
			Description: "",
			RequestDto:  comdto.ReqSearch{},
		},
	})

}
