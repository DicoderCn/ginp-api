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
func GenEntity() {

	entityName := gen.Input(InputEntityName, nil)
	entityName = gen.NameToCameBig(entityName)
	if entityName == "" {
		println("实体名称不能为空")
		return
	}
	lineName := gen.NameToLine(entityName)
	var replaceData map[string]string = map[string]string{
		ReplaceEntityName: entityName,
		ReplaceLineName:   lineName,
	}

	//开始生成
	gen.ReplaceAndWriteTemplate(
		filepath.Join(pathTemplateEntity),
		getPathEntity(lineName),
		replaceData,
	)
}

func getPathEntity(lineName string) string {
	p := filepath.Join(dirEntidy, lineName+".e.go")
	return p
}
