package service

import (
	"ginp-api/internal/app/gapi/repository"
	"ginp-api/internal/db/mysql"
)

var User *repository.DbUser

func DbUser() *repository.DbUser {
	if User == nil {
		dbRead := mysql.GetReadDb()
		dbWrite := mysql.GetWriteDb()
		User = repository.NewDbUser(dbRead, dbWrite)
	}
	return User
}
