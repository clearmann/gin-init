package model_type

import (
    "gorm.io/gorm"
)

type Post struct {
    gorm.Model
    Title     string `gorm:"column:title;type:varchar(128);not null;comment:标题"`
    Content   string `gorm:"column:content;type:text;comment:内容"`
    Tags      string `gorm:"column:tags;type:varchar(512);标签列表(json数组)"`
    ThumbNum  int    `gorm:"column:thumb_num;type:int;default:0;not null;comment:点赞数"`
    FavourNum int    `gorm:"column:favour_num;type:int;default:0;not null;comment:收藏数"`
    UUID      uint64 `gorm:"column:uuid;type:bigint unsigned;comment:创建用户的uuid"`
}

func (u *Post) TableName() string {
    return TableNamePost
}
