package desc

import (
	"ginp-api/internal/gen"
	"ginp-api/pkg/filehelper"
	"strings"
)

func AddImportRouterPackage(lineName string) {
	allSmallName := gen.NameToAllSmall(lineName)
	packgeName := "r" + allSmallName
	content, err := filehelper.ReadContent(PathRouterEntry())
	if err != nil {
		println(err.Error())
		return
	}

	if strings.Contains(content, packgeName) {
		println("outers_entry.go 已经存在", packgeName, "不再添加")
		return
	}

	placeHolder := `//{{placeholder}}//`
	importStr := `_ "ginp-api/internal/app/gapi/routers/` + packgeName + `"`

	newContent := strings.Replace(content, placeHolder, importStr+"\n\t"+placeHolder, -1)
	err = filehelper.WriteContent(PathRouterEntry(), newContent)
	if err != nil {
		println("outers_entry 写入失败" + err.Error())
		return
	} else {
		println("outers_entry import写入成功")
	}
}
