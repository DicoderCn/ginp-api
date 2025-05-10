package desc

import (
	"ginp-api/internal/gen"
	"path/filepath"
)

const (
	pathTemplateEntity = "./templates/entity.tmpl"
	dirEntidy          = "../../internal/app/gapi/entity"
)

// entityName传入大驼峰如 UserGroup
func GenEntity(entityName string) {
	if entityName == "" {
		println("实体名称不能为空")
		return
	}
	var replaceData map[string]string = map[string]string{
		ReplaceEntityName: entityName,
		ReplaceLineName:   gen.NameToLine(entityName),
	}
	//开始生成
	gen.ReplaceAndWriteTemplate(
		pathTemplateEntity,
		getPathEntity(entityName),
		replaceData,
	)
}

func getPathEntity(entityName string) string {
	return filepath.Join(dirEntidy, entityName+".go")
}
