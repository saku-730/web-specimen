package entity

import (
	"time"
)

// Occurrence は public.occurrence テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type Occurrence struct {
	// --- Table Columns ---
	OccurrenceID      uint       `gorm:"primaryKey;column:occurrence_id"`
	ProjectID         *uint       `gorm:"column:project_id"`
	UserID            *uint       `gorm:"column:user_id"`
	IndividualID      *int       `gorm:"column:individual_id"`
	Lifestage         *string    `gorm:"column:lifestage"`
	Sex               *string    `gorm:"column:sex"`
	ClassificationID  *uint       `gorm:"column:classification_id"`
	PlaceID           *uint       `gorm:"column:place_id"`
	AttachmentGroupID *uint       `gorm:"column:attachment_group_id"`
	BodyLength        *string    `gorm:"column:body_length"`
	LanguageID        *uint       `gorm:"column:language_id"`
	Note              *string    `gorm:"column:note"`
	CreatedAt         *time.Time  `gorm:"column:created_at;autoCreateTime"`
	Timezone          *string      `gorm:"column:timezone;not null"`

	// --- Relationships ---

	// ◆ Belongs To (所属)の関係 ◆
	// occurrenceテーブルが外部キーを持っている関係なのだ ➡️
	ClassificationJSON ClassificationJSON `gorm:"foreignKey:ClassificationID"`
	Language           Language           `gorm:"foreignKey:LanguageID"`
	Place              *Place              `gorm:"foreignKey:PlaceID"`
	Project            Project            `gorm:"foreignKey:ProjectID"`
	User               User               `gorm:"foreignKey:UserID"`

	// ◆ Has Many / Has One (所有)の関係 ◆
	// 他のテーブルからoccurrence_idで参照されている関係なのだ ⬅️

	// Occurrenceは多くのAttachmentGroupレコードを持つ (Has Many)
	AttachmentGroups []AttachmentGroup `gorm:"foreignKey:OccurrenceID"`

	// Occurrenceは多くのIdentificationを持つ (Has Many)
	Identifications []Identification `gorm:"foreignKey:OccurrenceID"`

	// Occurrenceは多くのMakeSpecimenレコードを持つ (Has Many)
	MakeSpecimens []MakeSpecimen `gorm:"foreignKey:OccurrenceID"`

	// Occurrenceは多くのObservationレコードを持つ (Has Many)
	Observations []Observation `gorm:"foreignKey:OccurrenceID"`

	// Occurrenceは多くのSpecimenレコードを持つ (Has Many)
	Specimens []Specimen `gorm:"foreignKey:OccurrenceID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (Occurrence) TableName() string {
	return "occurrence"
}
