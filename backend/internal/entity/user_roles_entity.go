// internal/entity/user_roles_entity.go

package entity

// UserRole は public.user_roles テーブルのレコードをマッピングするための構造体なのだ
// データベースの定義に沿って、すべてのカラムと関係性を定義しているのだ
type UserRole struct {
	// --- Table Columns ---
	RoleID   uint   `gorm:"primaryKey;column:role_id"`
	RoleName string `gorm:"column:role_name;not null;unique"`

	// --- Relationships ---

	// ◆ Has Many (所有)の関係 ◆
	// 他のテーブルからrole_idで参照されている関係なのだ ⬅️

	// この役割(Role)は多くのUserに割り当てられる (Has Many)
	Users []User `gorm:"foreignKey:RoleID"`
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (UserRole) TableName() string {
	return "user_roles"
}
