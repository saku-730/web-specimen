// internal/entity/attachment_group_entity.go
package entity

// AttachmentGroup は public.attachment_goup テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type AttachmentGroup struct {
	// --- Table Columns ---
	OccurrenceID uint `gorm:"primaryKey;column:occurrence_id"`
	AttachmentID uint `gorm:"primaryKey;column:attachment_id"`
	Priority     *int `gorm:"column:priority"`

	// --- Relationships ---

	// ◆ Belongs To (所属)の関係 ◆
	// attachment_goupテーブルが外部キーを持っている関係なのだ ➡️
	Occurrence Occurrence `gorm:"foreignKey:OccurrenceID"`
	Attachment Attachment `gorm:"foreignKey:AttachmentID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (AttachmentGroup) TableName() string {
	// psqlの出力に合わせて "attachment_goup" としているのだ
	return "attachment_goup"
}
