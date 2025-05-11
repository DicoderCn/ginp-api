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
func GetDirRepository() string {
	return filepath.Join(GetDirAPP(), "repository")
}

// 组装 controller 路径
func PathController(lineName string) string {
	lineName = gen.NameToLine(lineName)
	return filepath.Join(GetDirController(), "c"+lineName, lineName+".c.go")
}

// 组装 service 路径
func PathService(lineName string) string {
	lineName = gen.NameToLine(lineName)
	return filepath.Join(GetDirService(), "s"+lineName, lineName+".s.go")
}

// 组装 model 路径
func PathModel(lineName string) string {
	lineName = gen.NameToLine(lineName)
	return filepath.Join(GetDirRepository(), "m"+lineName, lineName+".m.go")
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
