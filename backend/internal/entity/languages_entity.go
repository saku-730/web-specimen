// internal/entity/languages_entity.go
package entity

// Language は public.languages テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type Language struct {
	// --- Table Columns ---
	LanguageID     uint    `gorm:"primaryKey;column:language_id"`
	LanguageCommon *string `gorm:"column:language_common"`

	// --- Relationships ---

	// ◆ Has Many (所有)の関係 ◆
	// 他のテーブルからlanguage_idで参照されている関係なのだ ⬅️

	// Languageは多くのOccurrenceで使われる (Has Many)
	Occurrences []Occurrence `gorm:"foreignKey:LanguageID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (Language) TableName() string {
	return "languages"
}
