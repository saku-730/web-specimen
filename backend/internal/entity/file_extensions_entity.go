package entity

// FileExtension は public.file_extensions テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type FileExtension struct {
	// --- Table Columns ---
	ExtensionID   uint    `gorm:"primaryKey;column:extension_id"`
	ExtensionText *string `gorm:"column:extension_text"`
	FileTypeID    *int    `gorm:"column:file_type_id"`

	// --- Relationships ---

	// ◆ Belongs To (所属)の関係 ◆
	// file_extensionsテーブルが外部キー(file_type_id)を持っている関係なのだ ➡️
	FileType FileType `gorm:"foreignKey:FileTypeID"`

	// ◆ Has Many (所有)の関係 ◆
	// 他のテーブルからextension_idで参照されている関係なのだ ⬅️
	Attachments []Attachment `gorm:"foreignKey:ExtensionID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (FileExtension) TableName() string {
	return "file_extensions"
}
