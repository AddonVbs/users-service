package user

import (
	//t "BackEnd/internal/taskservice"
	"time"

	"gorm.io/gorm"
)

type User struct {
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Id        int            `gorm:"primaryKey" json:"id"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	UpdatedAt time.Time      `json:"updated_at"`
	//ALTER
	//Tasks []t.Task `gorm:"foreignKey:UserID"`
}
