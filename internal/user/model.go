package user

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	Role      string    `json:"role"` // Roles: superadmin, admin, user
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
