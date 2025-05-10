package main

import (
	"ginp-api/cmd/gencode/desc"
	"ginp-api/internal/gen"
)

func main() {
	result := gen.Input(desc.InputGenType, nil)
	if result != "1" || result != "2" || result != "3" {
		println("输入错误")
		return
	}

	switch result {
	case "1":
		gen.Input(desc.InputEntityName, func(input string) {
			if input == "" {
				println("实体名称不能为空")
				return
			}
			desc.GenEntity(gen.NameToCameBig(input))
		})
	}

}
