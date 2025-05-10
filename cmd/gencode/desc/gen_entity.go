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
	var replaceData map[string]string = map[string]string{
		ReplaceEntityName: entityName,
		ReplaceLineName:   lineName,
	}

	//开始生成
	templatePath := filepath.Join(GetDirTemplate(), "entity.tmpl")
	outputFilePath := filepath.Join(GetDirEntidy(), lineName+".e.go")
	// fmt.Printf("templatePath: %v\n", templatePath)
	// fmt.Printf("outputFilePath: %v\n", outputFilePath)

	gen.ReplaceAndWriteTemplate(
		templatePath,
		outputFilePath,
		replaceData,
	)
}
