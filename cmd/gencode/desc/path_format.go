package desc

import (
	"ginp-api/internal/gen"
	"path/filepath"
)

// 组装 controller 路径
func PathController(lineName string) string {
	lineName = gen.NameToLine(lineName)
	allSmallName := gen.NameToAllSmall(lineName)
	return filepath.Join(GetDirController(), "c"+allSmallName, lineName+".c.go")
}

// 组装 service 路径
func PathService(lineName string) string {
	lineName = gen.NameToLine(lineName)
	allSmallName := gen.NameToAllSmall(lineName)
	return filepath.Join(GetDirService(), "s"+allSmallName, lineName+".s.go")
}

// 组装 model 路径
func PathModel(lineName string) string {
	lineName = gen.NameToLine(lineName)
	allSmallName := gen.NameToAllSmall(lineName)
	return filepath.Join(GetDirModel(), "m"+allSmallName, lineName+".m.go")
}

// 组装 fields 路径,放在model所在的目录下
func PathFields(lineName string) string {
	lineName = gen.NameToLine(lineName)
	modelDir := filepath.Dir(PathModel(lineName))
	return filepath.Join(modelDir, lineName+".f.go")
}

// 组装路由路径 router
func PathRouter(lineName string) string {
	// lineName = gen.NameToLine(lineName)
	allSmallName := gen.NameToAllSmall(lineName)
	return filepath.Join(GetDirController(), "c"+allSmallName, lineName+".r.go")
}

// 组装路由入口路径 routers_entry
func PathRouterEntry() string {
	// lineName = gen.NameToLine(lineName)
	// allSmallName := gen.NameToAllSmall(lineName)

	return filepath.Join(GetDirRouter(), "routers_import.go")
}

// 组装实体路径 entity
func PathEntity(lineName string) string {
	lineName = gen.NameToLine(lineName)
	// allSmallName := gen.NameToAllSmall(lineName)
	return filepath.Join(GetDirEntidy(), lineName+".e.go")
}

// 组装 api 路径
func PathAddApi(lineName string, apiNameLine string) string {
	lineName = gen.NameToLine(lineName)
	apiNameLine = gen.NameToLine(apiNameLine)
	allSmallName := gen.NameToAllSmall(lineName)
	return filepath.Join(GetDirController(), "c"+allSmallName, apiNameLine+".a.go")
}
