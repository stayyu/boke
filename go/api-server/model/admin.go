package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32     `gorm:"primarykey;type:int" json:"id"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

type Admin struct {
	BaseModel
	NickName string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Role     int    `gorm:"column:role;default:2;type:int comment'1表示普通，2表示管理员'"`
}
type Blog struct {
	BaseModel
	Title      string `gorm:"type:varchar(200)"`
	Content    string `gorm:"type:text"`
	CategoryID int32  `gorm:"type:int;index:idx_category"`
	Category   Category
}
type Category struct {
	BaseModel
	Name string `gorm:"type:varchar(50)"`
}

type User struct {
	BaseModel
	Password string `gorm:"type:varchar(100);not null"`
	NickName string `gorm:"type:varchar(20)"`
	Role     int    `gorm:"column:role;default:1;type:int comment '1表示普通用户, 2表示管理员'"`
}
