// internal/entity/place_entity.go

package entity

import (
	"database/sql/driver"
	"fmt"
)

// Point は PostGISの geography(Point,4326) 型をGORMで扱うためのカスタム型
type Point struct {
	Lat *float64
	Lng *float64
}

// Scan メソッドはデータベースから値（例: "SRID=4326;POINT(139.7 35.6)")を読み込む時に呼ばれるのだ
func (p *Point) Scan(value interface{}) error {
	var data []byte
	switch v := value.(type) {
	case []byte:
		data = v
	case string:
		data = []byte(v)
	default:
		return fmt.Errorf("unsupported type for Point: %T", value)
	}
	var lng, lat float64
	_, err := fmt.Sscanf(string(data), "SRID=4326;POINT(%f %f)", &lng, &lat)
	if err != nil {
		return fmt.Errorf("failed to scan point: %w", err)
	}
	p.Lat = &lat
	p.Lng = &lng
	return nil
}

// Value メソッドはデータベースに値を書き込む時に呼ばれるのだ
func (p Point) Value() (driver.Value, error) {
	if p.Lat == nil && p.Lng == nil {
		return nil, nil
	}
	pointString :=  fmt.Sprintf("SRID=4326;POINT(%f %f)", *p.Lng, *p.Lat)
	fmt.Printf("--- DEBUG: Sending to PostGIS ---> %s\n", pointString)
	return pointString, nil
}



// Place は public.places テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type Place struct {
	// --- Table Columns ---
	PlaceID      uint     `gorm:"primaryKey;column:place_id"`
	Coordinates  *Point   `gorm:"type:geography(Point,4326);column:coordinates"`
	PlaceNameID  *uint     `gorm:"column:place_name_id"`
	Accuracy     *float64 `gorm:"column:accuracy"`

	// --- Relationships ---

	// ◆ Belongs To (所属)の関係 ◆
	// placesテーブルが外部キー(place_name_id)を持っている関係なのだ ➡️
	PlaceNamesJSON PlaceNamesJSON `gorm:"foreignKey:PlaceNameID"`

	// ◆ Has Many (所有)の関係 ◆
	// 'occurrence'テーブルからplace_idで参照されている関係なのだ ⬅️
	Occurrences []Occurrence `gorm:"foreignKey:PlaceID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (Place) TableName() string {
	return "places"
}
