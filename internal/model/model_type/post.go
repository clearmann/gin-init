package model_type

import (
    "gorm.io/gorm"
)

type Post struct {
    gorm.Model
    Title     string `gorm:"column:title;type:varchar(128);not null;comment:标题" json:"title"`
    Content   string `gorm:"column:content;type:text;comment:内容" json:"content"`
    Tags      string `gorm:"column:tags;type:varchar(512);标签列表(json数组)" json:"tags"`
    ThumbNum  int    `gorm:"column:thumb_num;type:int;default:0;not null;comment:点赞数" json:"thumb_num"`
    FavourNum int    `gorm:"column:favour_num;type:int;default:0;not null;comment:收藏数" json:"favour_num"`
    UUID      uint64 `gorm:"column:uuid;type:bigint unsigned;comment:创建用户的uuid" json:"uuid"`
}

func (u *Post) TableName() string {
    return TableNamePost
}
