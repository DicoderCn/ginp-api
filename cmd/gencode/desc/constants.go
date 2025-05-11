package desc

import "ginp-api/internal/gen"

const (
	ReplaceEntityName = "$ENTITY_NAME$"
	ReplaceLineName   = "$ENTITY_LINE$"
)

// 基础替换数据 传入大驼峰如 UserGroup
func getBaseReplaceMap(BigCameName string) map[string]string {
	BigCameName = gen.NameToCameBig(BigCameName)
	lineName := gen.NameToLine(BigCameName)
	var replaceData map[string]string = map[string]string{
		ReplaceEntityName: BigCameName,
		ReplaceLineName:   lineName,
	}
	return replaceData
}
