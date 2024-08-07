package model_type

import (
    "time"

    "gorm.io/gorm"
)

type User struct {
    UUID      uint64 `gorm:"column:uuid;type:bigint unsigned;primaryKey"`
    Nickname  string `gorm:"column:nickname;type:varchar(32);not null"`
    Password  string `gorm:"column:password;type:varchar(128);not null"`
    Email     string `gorm:"column:email;type:varchar(32);not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) TableName() string {
    return TableNameUser
}
