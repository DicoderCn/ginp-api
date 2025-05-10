package main

import (
	"ginp-api/cmd/gencode/desc"
	"ginp-api/internal/gen"
)

func main() {
	t := gen.Input(desc.GenType, func(result string) {
		if result != "1" || result != "2" || result != "3" {
			println("输入错误")
		}
	})
	println("result" + t)
}
