package repository

import (
	"fmt"
	"ginp-api/internal/app/gapi/entity"
	"ginp-api/pkg/dbops"
	"ginp-api/pkg/where"

	"gorm.io/gorm"
)

type DbUser struct {
	dbWrite *gorm.DB
	dbRead  *gorm.DB
}

func NewDbUser(dbWrite_, dbRead_ *gorm.DB) *DbUser {
	return &DbUser{
		dbWrite: dbWrite_,
		dbRead:  dbRead_,
	}
}

// Create 创建数据
func (s *DbUser) Create(dtoCreate *entity.User) (*entity.User, error) {
	err := dbops.Create(dtoCreate, s.dbWrite)
	if err != nil {
		return nil, err
	}
	if dtoCreate == nil || dtoCreate.ID <= 0 {
		return nil, fmt.Errorf("创建失败，info数据为空")
	}

	return dtoCreate, err
}

// FindOne 查询一条数据
func (s *DbUser) FindOne(wheres []*where.Condition) (*entity.User, error) {
	entityInfo := new(entity.User)
	err := dbops.FindOne(&dbops.FindOneConfig{
		Wheres:    wheres,
		Db:        s.dbRead,
		NewEntity: entityInfo,
	})
	if err != nil {
		return nil, err
	}
	if entityInfo == nil || entityInfo.ID <= 0 {
		return nil, fmt.Errorf("查询到的数据为空")
	}
	return entityInfo, nil
}

func (s *DbUser) FindOneById(id uint) (*entity.User, error) {
	return s.FindOne(where.New("id", "=", id).Conditions())
}

// FindList 查询列表数据
func (s *DbUser) FindList(wheres []*where.Condition, extra *where.Extra) ([]*entity.User, uint, error) {

	var entityList []*entity.User
	//传入的entityList必须要加 &取地址符号，切片本身并不是指针，向函数传递一个切片时，实际上是复制了该切片的结构体
	err := dbops.FindList(&dbops.FindListConfig{
		Conditions:     wheres,
		Db:             s.dbRead,
		Extra:          extra,
		NewEntityList:  &entityList,
		GetSoftDelData: false,
		// Fields:        []string{"ID"},
		// RelationList: []*dbops.RelationItem{},
	})
	if err != nil {
		return nil, 0, err //返回空切片，0，错误
	}

	//开始统计总数
	total, err := dbops.GetTotal(wheres, new(entity.User), s.dbRead)
	if err != nil {
		return []*entity.User{}, 0, err
	}

	return entityList, uint(total), nil
}

// Update 更新数据
func (s *DbUser) Update(wheres []*where.Condition, dtoUpdate *entity.User, columnsCfg ...string) error {
	// dbops.UpdateWithDb(wheres, new(entity.User), dtoUpdate, s.dbWrite, columnsCfg...)
	err := dbops.Update(&dbops.UpdateConfNew{
		Wheres:           wheres,
		NewEntity:        new(entity.User),
		Db:               s.dbWrite,
		UpdateColumnsCfg: columnsCfg,
		DataUpdate:       dtoUpdate,
	})
	return err
}

// Delete 删除数据
func (s *DbUser) Delete(wheres []*where.Condition) error {
	err := dbops.Delete(&dbops.DeleteConfig{
		Wheres:     wheres,
		Db:         s.dbWrite,
		SoftDelete: false,
		NewEntity:  new(entity.User),
	})
	return err
}

func (s *DbUser) DeleteById(id uint) error {
	err := dbops.Delete(&dbops.DeleteConfig{
		Wheres:     where.New("id", "=", id).Conditions(),
		Db:         s.dbWrite,
		SoftDelete: false,
		NewEntity:  new(entity.User),
	})
	return err
}

// GetTotal 获取总数
func (s *DbUser) GetTotal(wheres []*where.Condition) (int64, error) {
	total, err := dbops.GetTotal(wheres, new(entity.User), s.dbRead)
	return total, err
}
