// internal/entity/user_entity.go
package entity

import "time"

// users table
type User struct {
	UserID      int       `gorm:"primaryKey;column:user_id"`
	UserName    string    `gorm:"column:user_name"`
	DisplayName string    `gorm:"column:display_name"`
	MailAddress *string   `gorm:"column:mail_address"`
	Password    *string   `gorm:"column:password"` // hashed
	RoleID      *int      `gorm:"column:role_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	Timezone    int16     `gorm:"column:timezone"`
}

// TableName メソッドでGORMにテーブル名を明示的に教えるのだ
func (User) TableName() string {
	return "users"
}
