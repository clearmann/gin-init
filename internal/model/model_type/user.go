package model_type

import (
    "time"

    "gorm.io/gorm"
)

type User struct {
    UUID      uint64 `gorm:"column:uuid;type:bigint unsigned;primaryKey;comment:用户uuid"`
    Username  string `gorm:"column:username;type:varchar(128);not null;comment:用户名"`
    Email     string `gorm:"column:email;type:varchar(32);comment:用户邮箱"`
    Password  string `gorm:"column:password;type:varchar(128);not null;comment:用户密码"`
    Avatar    string `gorm:"column:avatar;type:varchar(128);comment:用户头像"`
    Nickname  string `gorm:"column:nickname;type:varchar(128);comment:用户昵称"`
    UnionID   string `gorm:"column:union_id;type:varchar(32);comment:微信开放平台ID"`
    MpOpenID  string `gorm:"column:mp_open_id;type:varchar(32);comment:公众号OpenID"`
    Profile   string `gorm:"column:profile;type:text;comment:用户简介"`
    Role      string `gorm:"column:role;type:varchar(32);not null;default:user;comment:用户角色"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) TableName() string {
    return TableNameUser
}
