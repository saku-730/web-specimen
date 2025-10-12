// internal/entity/wiki_pages_entity.go

package entity

import (
	"time"
)

// WikiPage は public.wiki_pages テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type WikiPage struct {
	// --- Table Columns ---
	PageID      uint       `gorm:"primaryKey;column:page_id"`
	Title       *string    `gorm:"column:title"`
	UserID      *int       `gorm:"column:user_id"`
	CreatedDate time.Time  `gorm:"column:created_date;autoCreateTime"`
	UpdatedDate *time.Time `gorm:"column:updated_date"`
	ContentPath *string    `gorm:"column:content_path"`

	// --- Relationships ---

	// ◆ Belongs To (所属)の関係 ◆
	// wiki_pagesテーブルが外部キー(user_id)を持っている関係なのだ ➡️
	User User `gorm:"foreignKey:UserID"`

	// ◆ Has Many (所有)の関係 ◆
	// 他のテーブルからpage_idで参照されている関係なのだ ⬅️

	// このWikiページは多くのObservationMethodで参照される (Has Many)
	ObservationMethods []ObservationMethod `gorm:"foreignKey:PageID"`

	// このWikiページは多くのSpecimenMethodで参照される (Has Many)
	SpecimenMethods []SpecimenMethod `gorm:"foreignKey:PageID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (WikiPage) TableName() string {
	return "wiki_pages"
}
