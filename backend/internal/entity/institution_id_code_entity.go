// internal/entity/institution_id_code_entity.go

package entity

// InstitutionIDCode は public.institution_id_code テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type InstitutionIDCode struct {
	// --- Table Columns ---
	InstitutionID   uint    `gorm:"primaryKey;column:institution_id"`
	InstitutionCode *string `gorm:"column:institution_code"`

	// --- Relationships ---

	// ◆ Has Many (所有)の関係 ◆
	// 他のテーブルからinstitution_idで参照されている関係なのだ ⬅️

	// この機関は多くのSpecimenレコードを持つ (Has Many)
	Specimens []Specimen `gorm:"foreignKey:InstitutionID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (InstitutionIDCode) TableName() string {
	return "institution_id_code"
}
