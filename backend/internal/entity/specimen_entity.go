// internal/entity/specimen_entity.go
package entity

// Specimen は public.specimen テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type Specimen struct {
	// --- Table Columns ---
	SpecimenID       uint    `gorm:"primaryKey;column:specimen_id"`
	OccurrenceID     *uint    `gorm:"column:occurrence_id"`
	SpecimenMethodID *uint    `gorm:"column:specimen_method_id"`
	InstitutionID    *uint    `gorm:"column:institution_id"`
	CollectionID     *string `gorm:"column:collection_id"`

	// --- Relationships ---

	// ◆ Belongs To (所属)の関係 ◆
	// specimenテーブルが外部キーを持っている関係なのだ ➡️
	InstitutionIDCode InstitutionIDCode `gorm:"foreignKey:InstitutionID"`
	Occurrence        Occurrence        `gorm:"foreignKey:OccurrenceID"`
	SpecimenMethod    SpecimenMethod    `gorm:"foreignKey:SpecimenMethodID"`

	// ◆ Has One (所有)の関係 ◆
	// 他のテーブルからspecimen_idで参照されている関係なのだ ⬅️
	// 1つの標本に対して、作成記録は1つだけなので Has One になるのだ
	MakeSpecimen []MakeSpecimen `gorm:"foreignKey:SpecimenID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (Specimen) TableName() string {
	return "specimen"
}
