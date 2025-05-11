package templates

import (
	"ginp-api/cmd/gencode/desc"
	"ginp-api/internal/gen"
	"path/filepath"
	"testing"
)

// 快速将现有的文件做成模板文件
func TestReplace(t *testing.T) {

	demoLineName := "user"

	replaceData := map[string]string{
		demoLineName: desc.ReplaceLineName,
		"User":       desc.ReplaceEntityName,
	}

	//1.开始生成：entity文件
	tPathEntity := filepath.Join(desc.GetDirTemplate(), "entity.tmpl")
	oPathEntity := filepath.Join(desc.GetDirEntidy(), demoLineName, demoLineName+".e.go")
	gen.ReplaceAndWriteTemplate(oPathEntity, tPathEntity, replaceData)

	//2.开始生成：router文件
	tPathRouter := filepath.Join(desc.GetDirTemplate(), "router.tmpl")
	oPathRouter := filepath.Join(desc.GetDirRouter(), demoLineName+".r.go")
	gen.ReplaceAndWriteTemplate(oPathRouter, tPathRouter, replaceData)

	//3.开始生成：controller文件
	tPathController := filepath.Join(desc.GetDirTemplate(), "controller.tmpl")
	oPathController := desc.PathController(demoLineName)
	gen.ReplaceAndWriteTemplate(oPathController, tPathController, replaceData)

	//4.开始生成：service文件
	tPathService := filepath.Join(desc.GetDirTemplate(), "service.tmpl")
	oPathService := desc.PathService(demoLineName)
	gen.ReplaceAndWriteTemplate(oPathService, tPathService, replaceData)

	//5.开始生成：model文件
	tPathRepository := filepath.Join(desc.GetDirTemplate(), "model.tmpl")
	oPathRepository := desc.PathModel(demoLineName)
	gen.ReplaceAndWriteTemplate(oPathRepository, tPathRepository, replaceData)
}
