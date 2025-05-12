package desc

import (
	"fmt"
	"ginp-api/internal/gen"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// GenAddApi 生成API相关代码
// 主要功能：扫描controller目录下的所有子目录，并打印目录名和路径
func GenAddApi() {
	// 获取controller目录路径
	controllerDir := GetDirController()
	// 创建map用于存储目录名和路径的映射
	folderMap := make(map[string]string)

	// 读取controller目录下的所有文件和子目录
	dirs, err := os.ReadDir(controllerDir)
	if err != nil {
		panic(err)
	}

	// 遍历目录项，只处理子目录
	for _, dir := range dirs {
		if dir.IsDir() {
			// 将目录名和完整路径存入map
			folderMap[dir.Name()] = filepath.Join(controllerDir, dir.Name())
		}
	}

	// 打印所有目录名和路径
	index := 1
	for k := range folderMap {
		fmt.Printf("%d.%s\n", index, k)
		index++
	}

	inputCode := gen.Input(Select3, nil)
	// 判断输入是否为数字
	if index, err := strconv.Atoi(inputCode); err == nil {
		// 如果是数字，转换为对应的文件夹名
		if index > 0 && index <= len(folderMap) {
			keys := make([]string, 0, len(folderMap))
			for k := range folderMap {
				keys = append(keys, k)
			}
			inputCode = keys[index-1]
		} else {
			fmt.Println("输入的序号不在范围内，请重新输入")
			return
		}
	}
	//判断inputCode是否在folderMap的key中
	if _, ok := folderMap[inputCode]; !ok {
		fmt.Println("输入的代码不在列表中，请重新输入")
		return
	}

	apiName := gen.Input(Select3ApiName, nil)
	if apiName == "" || strings.Contains(apiName, " ") {
		fmt.Println(`api名称不能为空,且不能包含空格（API name cannot be empty and cannot contain spaces)`)
		return
	}
	apiNameBig := gen.NameToCameBig(apiName)
	apiNameLine := gen.NameToLine(apiNameBig)

	apiDir := filepath.Join(controllerDir, folderMap[inputCode])
	//扫描apiDir下的所有文件匹配到c.go结尾的文件
	apiFiles, err := os.ReadDir(apiDir)
	if err != nil {
		panic(err)
	}
	entityLineName := gen.NameToLine(inputCode)
	for _, apiFile := range apiFiles {
		if apiFile.IsDir() {
			continue
		}
		if strings.HasSuffix(apiFile.Name(), "c.go") {
			//替换apiFile.Name()中的c.go为a.go
			entityLineName = strings.Replace(apiFile.Name(), "c.go", "", -1)
		}
	}
	if entityLineName == "" {
		fmt.Println(apiDir + "下没有找到c.go结尾的文件！")
		return
	}

	apiPath := filepath.Join(apiDir, apiNameLine+".a.go")
	replaceData := getBaseReplaceMap(entityLineName)
	replaceData[ReplaceApiNameBig] = apiNameBig                                //api名称 大驼峰
	replaceData[ReplaceApiNameLine] = apiNameLine                              //api名称 下划线
	replaceData[ReplacePackageName] = "c" + gen.NameToAllSmall(entityLineName) //包名 全小写
	replaceData[ReplaceEntityName] = gen.NameToCameBig(entityLineName)         //实体名称 大驼峰
	gen.ReplaceAndWriteTemplate(TemplatePathAddApi(), apiPath, replaceData)

}
