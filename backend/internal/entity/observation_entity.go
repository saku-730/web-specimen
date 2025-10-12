// internal/entity/observation_entity.go

package entity

import (
	"time"
)

// Observation は public.observations テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type Observation struct {
	// --- Table Columns ---
	ObservationsID      uint      `gorm:"primaryKey;column:observations_id"`
	UserID              *int      `gorm:"column:user_id"`
	OccurrenceID        *int      `gorm:"column:occurrence_id"`
	ObservationMethodID *int      `gorm:"column:observation_method_id"`
	Behavior            *string   `gorm:"column:behavior"`
	ObservedAt          time.Time `gorm:"column:observed_at;autoCreateTime"`
	Timezone            int16     `gorm:"column:timezone;not null"`

	// --- Relationships ---

	// ◆ Belongs To (所属)の関係 ◆
	// observationsテーブルが外部キーを持っている関係なのだ ➡️
	User              User              `gorm:"foreignKey:UserID"`
	Occurrence        Occurrence        `gorm:"foreignKey:OccurrenceID"`
	ObservationMethod ObservationMethod `gorm:"foreignKey:ObservationMethodID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (Observation) TableName() string {
	return "observations"
}
