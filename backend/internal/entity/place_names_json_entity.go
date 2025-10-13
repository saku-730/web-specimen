// internal/entity/place_names_json_entity.go

package entity

import (
	"gorm.io/datatypes"
)

// PlaceNamesJSON は public.place_names_json テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type PlaceNamesJSON struct {
	// --- Table Columns ---
	PlaceNameID    uint           `gorm:"primaryKey;column:place_name_id"`
	ClassPlaceName datatypes.JSON `gorm:"column:class_place_name"`

	// --- Relationships ---

	// ◆ Has One (所有)の関係 ◆
	// 'places'テーブルからplace_name_idで参照されている関係なのだ ⬅️
	// 通常、1つの地名情報JSONは1つの場所に紐づくのでHas Oneの関係になるのだ
	Place []Place `gorm:"foreignKey:PlaceNameID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (PlaceNamesJSON) TableName() string {
	return "place_names_json"
}
