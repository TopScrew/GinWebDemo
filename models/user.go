package models

import "time"

/**
 * @Author: Screw
 * @Date: 2020/9/30 11:13 下午
 * @Desc:
 */

type User struct {
	ID        uint      `gorm:"primary_key;auto_increment"` // ID
	Name      string    `gorm:"type:varchar(64);not null"`
	Passwd    string    `gorm:"type:varchar(64);not null"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"` // 创建日期
	UpdatedAt time.Time `gorm:"not null;default:current_timestamp"` // 更新日期
}

// 设置表名
func (User) TableName() string {
	return "users"
}


