package desc

import (
	"ginp-api/internal/gen"
	"path/filepath"
)

// entityName传入大驼峰如 UserGroup
func GenEntity() {

	entityName := gen.Input(InputEntityName, nil)
	entityName = gen.NameToCameBig(entityName)
	if entityName == "" {
		println("实体名称不能为空")
		return
	}
	lineName := gen.NameToLine(entityName)
	replaceData := getBaseReplaceMap(entityName)

	//1.开始生成：entity文件
	tPathEntity := filepath.Join(GetDirTemplate(), "entity.tmpl")
	oPathEntity := filepath.Join(GetDirEntidy(), lineName+".e.go")
	gen.ReplaceAndWriteTemplate(tPathEntity, oPathEntity, replaceData)

	//2.开始生成：router文件
	tPathRouter := filepath.Join(GetDirTemplate(), "router.tmpl")
	oPathRouter := filepath.Join(GetDirRouter(), lineName+".r.go")
	gen.ReplaceAndWriteTemplate(tPathRouter, oPathRouter, replaceData)

	//3.开始生成：controller文件
	tPathController := filepath.Join(GetDirTemplate(), "controller.tmpl")
	oPathController := filepath.Join(GetDirController(), lineName+".c.go")
	gen.ReplaceAndWriteTemplate(tPathController, oPathController, replaceData)

	//4.开始生成：service文件
	tPathService := filepath.Join(GetDirTemplate(), "service.tmpl")
	oPathService := filepath.Join(GetDirService(), lineName+".s.go")
	gen.ReplaceAndWriteTemplate(tPathService, oPathService, replaceData)

	//5.开始生成：repository文件
	tPathRepository := filepath.Join(GetDirTemplate(), "repository.tmpl")
	oPathRepository := filepath.Join(GetDirRepository(), lineName+".p.go")
	gen.ReplaceAndWriteTemplate(tPathRepository, oPathRepository, replaceData)

}
