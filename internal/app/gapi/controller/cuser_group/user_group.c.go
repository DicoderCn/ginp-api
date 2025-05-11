package cuser_group

import (
	"ginp-api/internal/app/gapi/dto/comdto"
	"ginp-api/internal/app/gapi/entity"
	"ginp-api/internal/app/gapi/service/suser_group"
	"ginp-api/pkg/where"

	"github.com/DicoderCn/ginp"
)

func FindByID(ctx *ginp.ContextPlus) {
	var params *comdto.ReqFindById
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.FailData("请求参数有误" + err.Error())
		return
	}
	info, err := suser_group.Model().FindOneById(params.ID)
	if err != nil {
		ctx.FailData(err.Error())
		return
	}
	ctx.SuccessData(info)
}

func Create(ctx *ginp.ContextPlus) {
	var params *entity.UserGroup
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.FailData("请求参数有误" + err.Error())
		return
	}
	//也可以自己创建并传入读写db: tables.NewUserGroup(wdb,rdb)
	info, err := suser_group.Model().Create(params)
	if err != nil {
		ctx.FailData("创建失败" + err.Error())
		return
	}
	ctx.SuccessData(info)
}

func Update(ctx *ginp.ContextPlus) {
	var params *entity.UserGroup
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.FailData("请求参数有误" + err.Error())
		return
	}
	wheres := where.Format(where.OptEqual("id", params.ID))
	//也可以自己创建并传入读写db: tables.NewUserGroup(wdb,rdb)
	err := suser_group.Model().Update(wheres, params)
	if err != nil {
		ctx.FailData("修改失败" + err.Error())
		return
	}
	ctx.Success()
}

func Delete(ctx *ginp.ContextPlus) {
	var params *comdto.ReqDelete
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.FailData("请求参数有误" + err.Error())
		return
	}

	//也可以自己创建并传入读写db: tables.NewUserGroup(wdb,rdb)
	err := suser_group.Model().DeleteById(params.ID)
	if err != nil {
		ctx.FailData("删除失败" + err.Error())
		return
	}
	ctx.Success()
	return
}

func Search(ctx *ginp.ContextPlus) {
	var params *comdto.ReqSearch
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.FailData("请求参数有误" + err.Error())
		return
	}
	//也可以自己创建并传入读写db: tables.NewUserGroup(wdb,rdb)
	list, total, err := suser_group.Model().FindList(params.Wheres, params.Extra)
	if err != nil {
		ctx.FailData("查询失败" + err.Error())
		return
	}

	resp := &comdto.RespSearch{
		List:     list,
		Total:    uint(total),
		PageNum:  uint(params.Extra.PageNum),
		PageSize: uint(params.Extra.PageSize),
	}
	ctx.SuccessData(resp)

}
