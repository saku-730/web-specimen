// internal/entity/change_log_entity.go
package entity

import (
	"time"
)

// ChangeLog は public.change_logs テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type ChangeLog struct {
	// --- Table Columns ---
	LogID       uint      `gorm:"primaryKey;column:log_id"`
	Type        *string   `gorm:"column:type"`
	ChangedID   *int      `gorm:"column:changed_id"`
	BeforeValue *string   `gorm:"column:before_value"`
	AfterValue  *string   `gorm:"column:after_value"`
	UserID      *int      `gorm:"column:user_id"`
	Date        time.Time `gorm:"column:date;autoCreateTime"`
	Row         *string   `gorm:"column:row"` // "row"はGoの予約語ではないのでそのまま使えるのだ

	// --- Relationships ---

	// ◆ Belongs To (所属)の関係 ◆
	// change_logsテーブルが外部キー(user_id)を持っている関係なのだ ➡️
	User User `gorm:"foreignKey:UserID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (ChangeLog) TableName() string {
	return "change_logs"
}
