package desc

import (
	"ginp-api/internal/gen"
	"ginp-api/pkg/filehelper"
	"strings"
)

func AddImportRouterPackage(lineName string) {
	allSmallName := gen.NameToAllSmall(lineName)
	// packgeName := "r" + allSmallName
	content, err := filehelper.ReadContent(PathRouterEntry())
	if err != nil {
		println(err.Error())
		return
	}

	if strings.Contains(content, allSmallName) {
		println("outers_entry.go 已经存在", allSmallName, "不再添加")
		return
	}

	placeHolder := PlaceholderRouterImport
	importStr := RouterReplaceStr + allSmallName + `"`

	newContent := strings.Replace(content, placeHolder, importStr+"\n\t"+placeHolder, -1)
	err = filehelper.WriteContent(PathRouterEntry(), newContent)
	if err != nil {
		println("outers_entry 写入失败" + err.Error())
		return
	} else {
		println("outers_entry import写入成功")
	}
}
