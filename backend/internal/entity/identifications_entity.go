// internal/entity/identifications_entity.go

package entity

import (
	"time"
)

// Identification は public.identifications テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type Identification struct {
	// --- Table Columns ---
	IdentificationID uint      `gorm:"primaryKey;column:identification_id"`
	UserID           *uint      `gorm:"column:user_id"`
	OccurrenceID     *uint      `gorm:"column:occurrence_id"`
	SourceInfo       *string   `gorm:"column:source_info"`
	IdentificatedAt  *time.Time `gorm:"column:identificated_at;autoCreateTime"`
	Timezone         *string     `gorm:"column:timezone;not null"`

	// --- Relationships ---

	// ◆ Belongs To (所属)の関係 ◆
	// identificationsテーブルが外部キーを持っている関係なのだ ➡️
	Occurrence Occurrence `gorm:"foreignKey:OccurrenceID"`
	User       User       `gorm:"foreignKey:UserID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (Identification) TableName() string {
	return "identifications"
}
