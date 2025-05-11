package suser

import (
	"ginp-api/internal/app/gapi/model/muser"
	"ginp-api/internal/db/mysql"
)

var User *muser.Model

func Model() *muser.Model {
	if User == nil {
		dbRead := mysql.GetReadDb()
		dbWrite := mysql.GetWriteDb()
		User = muser.NewModel(dbRead, dbWrite)
	}
	return User
}
