package templates

import (
	"ginp-api/cmd/gencode/desc"
	"ginp-api/internal/gen"
	"testing"
)

// 快速将现有的文件做成模板文件
func TestReplace(t *testing.T) {

	demoLineName := "user"
	demoBigName := "User"
	replaceData := map[string]string{
		demoLineName: desc.ReplaceLineName,
		demoBigName:  desc.ReplaceEntityName,
	}

	//1.开始生成：entity文件
	tPathEntity := desc.TemplatePathEntity()
	oPathEntity := desc.PathEntity(demoLineName)
	gen.ReplaceAndWriteTemplate(oPathEntity, tPathEntity, replaceData)

	//2.开始生成：router文件
	tPathRouter := desc.TemplatePathRouter()
	oPathRouter := desc.PathRouter(demoLineName)
	gen.ReplaceAndWriteTemplate(oPathRouter, tPathRouter, replaceData)

	//3.开始生成：controller文件
	tPathController := desc.TemplatePathController()
	oPathController := desc.PathController(demoLineName)
	gen.ReplaceAndWriteTemplate(oPathController, tPathController, replaceData)

	//4.开始生成：service文件
	tPathService := desc.TemplatePathService()
	oPathService := desc.PathService(demoLineName)
	gen.ReplaceAndWriteTemplate(oPathService, tPathService, replaceData)

	//5.开始生成：model文件
	tPathRepository := desc.TemplatePathModel()
	oPathRepository := desc.PathModel(demoLineName)
	gen.ReplaceAndWriteTemplate(oPathRepository, tPathRepository, replaceData)
}
