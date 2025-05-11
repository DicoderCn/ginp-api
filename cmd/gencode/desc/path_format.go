package desc

import (
	"ginp-api/internal/gen"
	"path/filepath"
)

// 组装 controller 路径
func PathController(lineName string) string {
	//baseName := gen.NameToLine(lineName)
	baseName := gen.NameToAllSmall(lineName)
	return filepath.Join(GetDirController(), "c"+baseName, baseName+".c.go")
}

// 组装 service 路径
func PathService(lineName string) string {

	//baseName := gen.NameToLine(lineName)
	baseName := gen.NameToAllSmall(lineName)
	return filepath.Join(GetDirService(), "s"+baseName, baseName+".s.go")
}

// 组装 model 路径
func PathModel(lineName string) string {
	//baseName := gen.NameToLine(lineName)
	baseName := gen.NameToAllSmall(lineName)
	return filepath.Join(GetDirModel(), "m"+baseName, baseName+".m.go")
}

// 组装路由路径 router
func PathRouter(lineName string) string {
	baseName := gen.NameToAllSmall(lineName)
	return filepath.Join(GetDirRouter(), baseName+".r.go")
}

// 组装实体路径 entity
func PathEntity(lineName string) string {
	baseName := gen.NameToAllSmall(lineName)
	return filepath.Join(GetDirEntidy(), baseName, baseName+".e.go")
}
