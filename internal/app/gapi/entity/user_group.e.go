package entity

import (
	"ginp-api/internal/app/gapi/typ"
	"ginp-api/internal/gen"
	"time"
)

var _ typ.IEntity = (*UserGroup)(nil) // U实体必须实现接口GenConfig
const tableNameUserGroup = "user_group"

type UserGroup struct {
	ID int `json:"id"`
	//... other fields
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at" `
}

func (UserGroup) GenConfig() *gen.EntityConfig {
	return &gen.EntityConfig{
		TableName: tableNameUserGroup,
	}
}

func (UserGroup) TableName() string {
	return tableNameUserGroup
}
