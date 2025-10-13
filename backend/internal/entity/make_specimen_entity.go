// internal/entity/make_specimen_entity.go

package entity

import (
	"time"
)

// MakeSpecimen は public.make_specimen テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type MakeSpecimen struct {
	// --- Table Columns ---
	MakeSpecimenID   uint       `gorm:"primaryKey;column:make_specimen_id"`
	OccurrenceID     *uint       `gorm:"column:occurrence_id"`
	UserID           *uint       `gorm:"column:user_id"`
	SpecimenID       *uint       `gorm:"column:specimen_id"`
	Date             *time.Time `gorm:"column:date"`
	SpecimenMethodID *uint       `gorm:"column:specimen_method_id"`
	CreatedAt        *time.Time  `gorm:"column:created_at;autoCreateTime"`
	Timezone         *string      `gorm:"column:timezone;not null"`

	// --- Relationships ---

	// ◆ Belongs To (所属)の関係 ◆
	// make_specimenテーブルが外部キーを持っている関係なのだ ➡️
	Occurrence     Occurrence     `gorm:"foreignKey:OccurrenceID"`
	User           User           `gorm:"foreignKey:UserID"`
	Specimen       Specimen       `gorm:"foreignKey:SpecimenID"`
	SpecimenMethod SpecimenMethod `gorm:"foreignKey:SpecimenMethodID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (MakeSpecimen) TableName() string {
	return "make_specimen"
}
