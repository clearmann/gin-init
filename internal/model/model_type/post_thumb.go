package model_type

import "gorm.io/gorm"

type PostThumb struct {
    gorm.Model
    PostID int    `gorm:"column:post_id;type:int;not null;comment:帖子id"`
    UUID   uint64 `gorm:"column:uuid;type:bigint unsigned;not null;comment:用户uuid"`
}

func (u *PostThumb) TableName() string {
    return TableNamePostThumb
}
