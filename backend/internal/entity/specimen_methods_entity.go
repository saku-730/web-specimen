// internal/entity/specimen_methods_entity.go

package entity

// SpecimenMethod は public.specimen_methods テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type SpecimenMethod struct {
	// --- Table Columns ---
	SpecimenMethodsID uint    `gorm:"primaryKey;column:specimen_methods_id"`
	MethodCommonName  *string `gorm:"column:method_common_name"`
	PageID            *uint    `gorm:"column:page_id"`

	// --- Relationships ---

	// ◆ Belongs To (所属)の関係 ◆
	// specimen_methodsテーブルが外部キー(page_id)を持っている関係なのだ ➡️
	WikiPage WikiPage `gorm:"foreignKey:PageID"`

	// ◆ Has Many (所有)の関係 ◆
	// 他のテーブルからspecimen_method_idで参照されている関係なのだ ⬅️

	// この標本作成方法は多くのMakeSpecimenレコードで使われる (Has Many)
	MakeSpecimens []MakeSpecimen `gorm:"foreignKey:SpecimenMethodID"`

	// この標本作成方法は多くのSpecimenレコードで使われる (Has Many)
	Specimens []Specimen `gorm:"foreignKey:SpecimenMethodID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (SpecimenMethod) TableName() string {
	return "specimen_methods"
}
