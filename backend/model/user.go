package model

import "time"

type User struct {
	ID        uint      `gorm:"primarykey;autoIncrement" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Age       uint      `gorm:"column:age" json:"age"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli" json:"created_at"`
}
