// internal/entity/projects_entity.go
package entity

import (
	"time"
)

// Project は public.projects テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type Project struct {
	// --- Table Columns ---
	ProjectID    uint       `gorm:"primaryKey;column:project_id"`
	ProjectName  string     `gorm:"column:project_name;not null"`
	Disscription *string    `gorm:"column:disscription"`
	StartDay     *time.Time `gorm:"column:start_day"`
	FinishedDay  *time.Time `gorm:"column:finished_day"`
	UpdatedDay   *time.Time `gorm:"column:updated_day"`
	Note         *string    `gorm:"column:note"`

	// --- Relationships ---

	// ◆ Has Many (所有)の関係 ◆
	// 他のテーブルからproject_idで参照されている関係なのだ ⬅️

	// Projectは多くのOccurrenceを持つ (Has Many)
	Occurrences []Occurrence `gorm:"foreignKey:ProjectID"`

	// Projectは多くのProjectMemberレコードを持つ (Has Many)
	ProjectMembers []ProjectMember `gorm:"foreignKey:ProjectID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (Project) TableName() string {
	return "projects"
}
