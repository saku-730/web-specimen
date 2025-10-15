// internal/entity/user_entity.go
package entity

import (
	"time"
)

// User は public.users テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type User struct {
	// --- Table Columns ---
	UserID      uint      `gorm:"primaryKey;column:user_id"`
	UserName    string    `gorm:"column:user_name;not null"`
	DisplayName string    `gorm:"column:display_name;not null"`
	MailAddress *string   `gorm:"column:mail_address;unique"`
	Password    *string   `gorm:"column:password"`
	RoleID      *int      `gorm:"column:role_id"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	Timezone    int16     `gorm:"column:timezone;not null"`

	// --- Relationships ---

	// ◆ Belongs To (所属)の関係 ◆
	// usersテーブルが外部キー(role_id)を持っている関係なのだ ➡️
	UserRole UserRole `gorm:"foreignKey:RoleID"`

	// ◆ Has One / Has Many (所有)の関係 ◆
	// 他のテーブルからuser_idで参照されている関係なのだ ⬅️
	
	// Userは一つのUserDefaultsを持つ (Has One)
	UserDefault UserDefault `gorm:"foreignKey:UserID"`

	// Userは多くのAttachmentを持つ (Has Many)
	Attachments []Attachment `gorm:"foreignKey:UserID"`

	// Userは多くのChangeLogを持つ (Has Many)
	ChangeLogs []ChangeLog `gorm:"foreignKey:UserID"`

	// Userは多くのIdentificationを持つ (Has Many)
	Identifications []Identification `gorm:"foreignKey:UserID"`

	// Userは多くのMakeSpecimenを持つ (Has Many)
	MakeSpecimens []MakeSpecimen `gorm:"foreignKey:UserID"`

	// Userは多くのObservationを持つ (Has Many)
	Observations []Observation `gorm:"foreignKey:UserID"`

	// Userは多くのOccurrenceを持つ (Has Many)
	Occurrences []Occurrence `gorm:"foreignKey:UserID"`

	// Userは多くのProjectMemberレコードを持つ (Has Many)
	ProjectMembers []ProjectMember `gorm:"foreignKey:UserID"`

	// Userは多くのWikiPageを持つ (Has Many)
	WikiPages []WikiPage `gorm:"foreignKey:UserID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (User) TableName() string {
	return "users"
}
