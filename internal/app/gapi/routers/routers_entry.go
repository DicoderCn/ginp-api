package routers

//显式导入：来确保它的 init 函数被调用
import (
	_ "ginp-api/internal/app/gapi/routers/rdemotable"
	_ "ginp-api/internal/app/gapi/routers/rindex"
	_ "ginp-api/internal/app/gapi/routers/ruser"
	// {{placeholder}}//
	// 上面的占位符请不要动动，否则会导致生成工具无法自动替换
)

func InitRouters() {
	//这里什么都不用执行，请勿删除，只需要在router包调用该函数即可
}
