package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID        int64  `gorm:"column:id; type:int; primaryKey; autoIncrement; comment:'primaryKey is ID"`
	RoleName  string `gorm:"column:role_name"`
	RolesNote string `gorm:"column:role_note; type:text"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
