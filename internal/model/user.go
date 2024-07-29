package model

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    UUID     uint64 `gorm:"column:uuid;type:uuid"`
    Nickname string `gorm:"column:nickname;type:varchar(32)"`
    Password string `gorm:"column:password;type:varchar(32)"`
    Email    string `gorm:"column:email;type:varchar(32)"`
}

func (u *User) TableName() string {
    return "user"
}
