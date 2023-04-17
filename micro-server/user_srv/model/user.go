package model

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel 基础model类，定义了所有表基本的字段
type BaseModel struct {
	ID       int32     `gorm:"primaryKey"`
	CreateAt time.Time `gorm:"column:add_time"`
	UpdateAt time.Time `gorm:"column:update_time"`
	DeleteAt gorm.DeletedAt
	isDelete bool
}

type User struct {
	BaseModel            // 继承基本字段
	Mobile    string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password  string     `gorm:"column:password;not null;type:varchar(100)"`
	NickName  string     `gorm:"type:varchar(20)"`
	Birthday  *time.Time `gorm:"type:datetime"`
	Gender    string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female女,male男'"`
	Role      int        `gorm:"column:role;default:1;type:int comment '1表示普通用户 2表示管理员'"`
}
