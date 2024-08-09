package model_type

import (
    "time"

    "gorm.io/gorm"
)

type User struct {
    UUID      uint64 `gorm:"column:uuid;type:bigint unsigned;primaryKey;comment:用户uuid" json:"uuid,omitempty"`
    Username  string `gorm:"column:username;type:varchar(128);not null;comment:用户名" json:"username,omitempty"`
    Email     string `gorm:"column:email;type:varchar(32);comment:用户邮箱" json:"email,omitempty"`
    Password  string `gorm:"column:password;type:varchar(128);not null;comment:用户密码" json:"password,omitempty"`
    Avatar    string `gorm:"column:avatar;type:varchar(128);comment:用户头像" json:"avatar,omitempty"`
    Nickname  string `gorm:"column:nickname;type:varchar(128);comment:用户昵称" json:"nickname,omitempty"`
    UnionID   string `gorm:"column:union_id;type:varchar(32);comment:微信开放平台ID" json:"union_id,omitempty"`
    MpOpenID  string `gorm:"column:mp_open_id;type:varchar(32);comment:公众号OpenID" json:"mp_open_id,omitempty"`
    Profile   string `gorm:"column:profile;type:text;comment:用户简介" json:"profile,omitempty"`
    Role      string `gorm:"column:role;type:varchar(32);not null;default:user;comment:用户角色" json:"role,omitempty"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) TableName() string {
    return TableNameUser
}
