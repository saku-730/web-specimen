// internal/entity/file_types_entity.go
package entity

// FileType は public.file_types テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type FileType struct {
	// --- Table Columns ---
	FileTypeID uint    `gorm:"primaryKey;column:file_type_id"`
	TypeName   *string `gorm:"column:type_name"`

	// --- Relationships ---

	// ◆ Has Many (所有)の関係 ◆
	// 他のテーブルからfile_type_idで参照されている関係なのだ ⬅️

	// このファイルタイプは多くのFileExtensionを持つ (Has Many)
	FileExtensions []FileExtension `gorm:"foreignKey:FileTypeID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (FileType) TableName() string {
	return "file_types"
}
