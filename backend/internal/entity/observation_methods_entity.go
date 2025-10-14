// internal/entity/observation_methods_entity.go

package entity

// ObservationMethod は public.observation_methods テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type ObservationMethod struct {
	// --- Table Columns ---
	ObservationMethodID uint    `gorm:"primaryKey;column:observation_method_id"`
	MethodCommonName    *string `gorm:"column:method_common_name"`
	PageID              *uint    `gorm:"column:pageid"`

	// --- Relationships ---

	// ◆ Belongs To (所属)の関係 ◆
	// observation_methodsテーブルが外部キー(pageid)を持っている関係なのだ ➡️
	WikiPage WikiPage `gorm:"foreignKey:PageID"`

	// ◆ Has Many (所有)の関係 ◆
	// 他のテーブルからobservation_method_idで参照されている関係なのだ ⬅️
	Observations []Observation `gorm:"foreignKey:ObservationMethodID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (ObservationMethod) TableName() string {
	return "observation_methods"
}
