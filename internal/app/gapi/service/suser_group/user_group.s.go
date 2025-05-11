package suser_group

import (
	"ginp-api/internal/app/gapi/model/muser_group"
	"ginp-api/internal/db/mysql"
)

var UserGroup *muser_group.Model

func Model() *muser_group.Model {
	if UserGroup == nil {
		dbRead := mysql.GetReadDb()
		dbWrite := mysql.GetWriteDb()
		UserGroup = muser_group.NewModel(dbRead, dbWrite)
	}
	return UserGroup
}
