package desc

import "ginp-api/internal/gen"

const (
	ReplaceEntityName = "$ENTITY_NAME$"
	ReplaceLineName   = "$ENTITY_LINE$"
	//全小写命名
	ReplacePackageName      = "$PACKAGE_NAME$"
	ReplaceApiNameBig       = "$API_NAME_BIG$"
	ReplaceApiNameLine      = "$API_NAME_LINE$"
	PlaceholderRouterImport = "//{{placeholder_router_import}}//"
	RouterReplaceStr        = `_ "ginp-api/internal/app/gapi/controller/c`
)

// 基础替换数据 传入大驼峰如 $ENTITY_NAME$Group
func getBaseReplaceMap(BigCameName string) map[string]string {
	BigCameName = gen.NameToCameBig(BigCameName)
	lineName := gen.NameToLine(BigCameName)
	var replaceData map[string]string = map[string]string{
		ReplaceEntityName:  BigCameName,
		ReplaceLineName:    lineName,
		ReplacePackageName: gen.NameToAllSmall(lineName),
	}
	return replaceData
}
