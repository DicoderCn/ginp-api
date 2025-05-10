package entity

import (
	"ginp-api/internal/app/gapi/typ"
	"ginp-api/internal/gen"
	"time"
)

var _ typ.IEntity = (*User)(nil) // U实体必须实现接口GenConfig
const tableNameUser = "users"

type User struct {
	ID int `json:"id"`
	//... other fields
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at" `
}

func (User) GenConfig() *gen.EntityConfig {
	return &gen.EntityConfig{
		TableName: tableNameUser,
	}
}

func (User) TableName() string {
	return tableNameUser
}
