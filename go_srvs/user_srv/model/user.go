package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool `gorm:"column:is_deleted"`
}

// User /*
type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null comment '手机号唯一'"` // 加索引
	Password string     `gorm:"type:varchar(100);not null;comment '密码 md5加密'"`
	NickName string     `gorm:"type:varchar(20) comment '昵称'"`
	Birthday *time.Time `gorm:"type:datetime comment '生日'"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female表示女性 male表示男性'"`
	Role     int        `gorm:"default:1;type:int comment '1表示普通用户2表示管理员'"`
}
