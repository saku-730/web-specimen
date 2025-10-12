// internal/entity/user_defaults_entity.go
package entity

// UserDefault は users_defaults テーブルのレコードをマッピングするための構造体なのだ
type UserDefault struct {
	UserID                      int     `gorm:"primaryKey;column:user_id"`
	ProjectID                   *int    `gorm:"column:project_id"`
	ProjectName                 *string `gorm:"column:project_name"`
	IndividualID                *int    `gorm:"column:individual_id"`
	Lifestage                   *string `gorm:"column:lifestage"`
	Sex                         *string `gorm:"column:sex"`
	LanguageID                  *int    `gorm:"column:language_id"`
	LanguageCommon              *string `gorm:"column:language_common"`
	PlaceName                   *string `gorm:"column:place_name"`
	Note                        *string `gorm:"column:note"`
	ClassificationSpecies       *string `gorm:"column:classification_species"`
	ClassificationGenus         *string `gorm:"column:classification_genus"`
	ClassificationFamily        *string `gorm:"column:classification_family"`
	ClassificationOrder         *string `gorm:"column:classification_order"`
	ClassificationClass         *string `gorm:"column:classification_class"`
	ClassificationPhylum        *string `gorm:"column:classification_phylum"`
	ClassificationKingdom       *string `gorm:"column:classification_kingdom"`
	ClassificationOthers        *string `gorm:"column:classification_others"`
	ObservationUserID           *int    `gorm:"column:observation_user_id"`
	ObservationUserName         *string `gorm:"column:observation_user_name"`
	ObservationMethodID         *int    `gorm:"column:observation_method_id"`
	ObservationMethodName       *string `gorm:"column:observation_method_name"`
	ObservationBehavior         *string `gorm:"column:observation_behavior"`
	ObservationObservedAt       *string `gorm:"column:observation_observed_at"`
	SpecimenUserID              *int    `gorm:"column:specimen_user_id"`
	SpecimenUserName            *string `gorm:"column:specimen_user_name"`
	SpecimenMethodID            *int    `gorm:"column:specimen_method_id"`
	SpecimenMethodName          *string `gorm:"column:specimen_method_name"`
	IdentificationUserID        *int    `gorm:"column:identification_user_id"`
	IdentificationUserName      *string `gorm:"column:identification_user_name"`
	IdentificationIdentifiedAt  *string `gorm:"column:identification_identified_at"`
	IdentificationSourceInfo    *string `gorm:"column:identification_source_info"`
	UserName                    *string `gorm:"column:user_name"` // これはuser_idに紐づくuser_name
}

// TableName メソッドで、GORMにこの構造体がどのテーブルに対応するかを教えるのだ
func (UserDefault) TableName() string {
	return "users_defaults"
}
