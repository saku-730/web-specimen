// internal/entity/classification_json_entity.go

package entity

import (
	"gorm.io/datatypes"
)

// ClassificationJSON は public.classification_json テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type ClassificationJSON struct {
	// --- Table Columns ---
	ClassificationID   uint           `gorm:"primaryKey;column:classification_id"`
	ClassClassification datatypes.JSON `gorm:"column:class_classification"`

	// --- Relationships ---

	// 'occurrence'テーブルからclassification_idで参照されている関係なのだ ⬅️
	Occurrence []Occurrence `gorm:"foreignKey:ClassificationID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (ClassificationJSON) TableName() string {
	return "classification_json"
}
