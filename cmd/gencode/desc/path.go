package desc

import (
	"ginp-api/internal/gen"
	"os"
	"path/filepath"
	"strings"
)

func GetPwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// println("pwd:", pwd)
	return pwd
}

func GetDirProject() string {
	pwd := GetPwd()
	if strings.Contains(pwd, "/cmd") {
		// /Users/xiaod/projects_new/my/ginp-api/cmd/gencode
		arr := strings.Split(pwd, "/cmd")
		return arr[0]
	}
	return pwd
}

func GetDirAPP() string {
	projetDir := GetDirProject()
	return filepath.Join(projetDir, "internal", "app", "gapi")
}
func GetDirGencode() string {
	projetDir := GetDirProject()
	return filepath.Join(projetDir, "cmd", "gencode")
}

func GetDirTemplate() string {
	return filepath.Join(GetDirGencode(), "templates")
}
func GetDirRouter() string {
	return filepath.Join(GetDirAPP(), "router")
}

func GetDirEntidy() string {
	return filepath.Join(GetDirAPP(), "entity")
}

func GetDirController() string {
	return filepath.Join(GetDirAPP(), "controller")
}
func GetDirService() string {
	return filepath.Join(GetDirAPP(), "service")
}
func GetDirModel() string {
	return filepath.Join(GetDirAPP(), "model")
}

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

// 模板路径 entity
func TemplatePathEntity() string {
	return filepath.Join(GetDirTemplate(), "entity.tmpl")
}

// 模板路径 router
func TemplatePathRouter() string {
	return filepath.Join(GetDirTemplate(), "router.tmpl")
}

// 模板路径 controller
func TemplatePathController() string {
	return filepath.Join(GetDirTemplate(), "controller.tmpl")
}

// 模板路径 service
func TemplatePathService() string {
	return filepath.Join(GetDirTemplate(), "service.tmpl")
}

// 模板路径 model
func TemplatePathModel() string {
	return filepath.Join(GetDirTemplate(), "model.tmpl")
}
