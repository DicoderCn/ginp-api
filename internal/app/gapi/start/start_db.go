package start

import (
	"fmt"
	"ginp-api/configs"
	"ginp-api/internal/db/mysql"
	"ginp-api/pkg/cfg"
)

func startDB() {
	mysql.InitDb(
		cfg.GetString(configs.ConfigKeyMysqlIp),
		cfg.GetString(configs.ConfigKeyMysqlPort),
		cfg.GetString(configs.ConfigKeyMysqlUser),
		cfg.GetString(configs.ConfigKeyMysqlDb),
		cfg.GetString(configs.ConfigKeyMysqlPwd),
	)

	//迁移表
	if mysql.GetWriteDb() != nil {
		//自动迁移表结构
		err := mysql.GetWriteDb().AutoMigrate(EntityAutoMigrateList...)
		if err != nil {
			fmt.Println("迁移表结构失败" + err.Error())
			panic(err)
		}
	}
}
