package model_type

import "gorm.io/gorm"

type PostFavour struct {
    gorm.DB
    PostID int    `gorm:"column:post_id;type:int;not null;comment:帖子id"`
    UUID   uint64 `gorm:"column:uuid;type:bigint unsigned;not null;comment:用户uuid"`
}

func (u *PostFavour) TableName() string {
    return TableNamePostFavour
}
