// internal/entity/attachments_entity.go

package entity

import (
	"time"
)

// Attachment は public.attachments テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type Attachment struct {
	// --- Table Columns ---
	AttachmentID uint       `gorm:"primaryKey;column:attachment_id"`
	FilePath     string     `gorm:"column:file_path;not null"`
	ExtensionID  *int       `gorm:"column:extension_id"`
	UserID       *int       `gorm:"column:user_id"`
	Uploaded     *time.Time `gorm:"column:uploaded"`
	Note         *string    `gorm:"column:note"`

	// --- Relationships ---

	// ◆ Belongs To (所属)の関係 ◆
	// attachmentsテーブルが外部キーを持っている関係なのだ ➡️
	FileExtension FileExtension `gorm:"foreignKey:ExtensionID"`
	User          User          `gorm:"foreignKey:UserID"`

	// ◆ Has Many (所有)の関係 ◆
	// 他のテーブルからattachment_idで参照されている関係なのだ ⬅️
	AttachmentGroups []AttachmentGroup `gorm:"foreignKey:AttachmentID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (Attachment) TableName() string {
	return "attachments"
}
