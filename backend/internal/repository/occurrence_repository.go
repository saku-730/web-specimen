// internal/repository/occurrence_repository.go
package repository

import (
	"github.com/saku-730/web-specimen/backend/internal/entity"
	"github.com/saku-730/web-specimen/backend/internal/model"
	"gorm.io/gorm"
)

// DropdownRepository はドロップダウンリストのデータ取得を定義するインターフェースなのだ
type DropdownRepository interface {
	GetDropdownLists() (*model.DropdownListResponse, error)
}

type dropdownRepository struct {
	db *gorm.DB
}

// NewDropdownRepository は新しいリポジトリのインスタンスを作成するのだ
func NewDropdownRepository(db *gorm.DB) DropdownRepository {
	return &dropdownRepository{db: db}
}

// GetDropdownLists は各テーブルからリスト作成に必要な情報を取得してくるのだ
func (r *dropdownRepository) GetDropdownLists() (*model.Dropdowns, error) {
	var users []model.DropdownUser
	var projects []model.DropdownProject
	var languages []model.DropdownLanguage
	var obsMethods []model.DropdownObservationMethod
	var specMethods []model.DropdownSpecimenMethod
	var institutions []model.DropdownInstitution

	// Users テーブルから取得
	if err := r.db.Model(&entity.User{}).Select("user_id, user_name").Find(&users).Error; err != nil {
		return nil, err
	}

	// Projects テーブルから取得
	if err := r.db.Model(&entity.Project{}).Select("project_id, project_name").Find(&projects).Error; err != nil {
		return nil, err
	}

	// Languages テーブルから取得
	if err := r.db.Model(&entity.Language{}).Select("language_id, language_common").Find(&languages).Error; err != nil {
		return nil, err
	}

	// ObservationMethods テーブルから取得 (カラム名をモデルのフィールド名に合わせるのだ)
	if err := r.db.Model(&entity.ObservationMethod{}).Select("observation_method_id, method_common_name AS observation_method_name").Find(&obsMethods).Error; err != nil {
		return nil, err
	}

	// SpecimenMethods テーブルから取得 (こちらもカラム名を合わせるのだ)
	if err := r.db.Model(&entity.SpecimenMethod{}).Select("specimen_methods_id, method_common_name AS specimen_methods_common").Find(&specMethods).Error; err != nil {
		return nil, err
	}

	// InstitutionIDCode テーブルから取得
	if err := r.db.Model(&entity.InstitutionIDCode{}).Select("institution_id, institution_code").Find(&institutions).Error; err != nil {
		return nil, err
	}

	// 取得した各リストを一つのレスポンス構造体にまとめるのだ
	response := &model.DropdownListResponse{
		Users:              users,
		Projects:           projects,
		Languages:          languages,
		ObservationMethods: obsMethods,
		SpecimenMethods:    specMethods,
		Institutions:       institutions,
	}

	return response, nil
}
