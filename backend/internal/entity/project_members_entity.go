// internal/entity/project_members_entity.go

package entity

import (
	"time"
)

// ProjectMember は public.project_members テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type ProjectMember struct {
	// --- Table Columns ---
	ProjectMemberID uint       `gorm:"primaryKey;column:project_member_id"`
	ProjectID       *int       `gorm:"column:project_id"`
	UserID          *int       `gorm:"column:user_id"`
	JoinDay         *time.Time `gorm:"column:join_day"`
	FinishDay       *time.Time `gorm:"column:finish_day"`

	// --- Relationships ---

	// ◆ Belongs To (所属)の関係 ◆
	// project_membersテーブルが外部キーを持っている関係なのだ ➡️
	Project Project `gorm:"foreignKey:ProjectID"`
	User    User    `gorm:"foreignKey:UserID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (ProjectMember) TableName() string {
	return "project_members"
}
