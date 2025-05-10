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
			desc.GenEntity(gen.NameToCameBig(input))
		})
	}

}
