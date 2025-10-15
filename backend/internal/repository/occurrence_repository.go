// internal/repository/occurrence_repository.go
package repository

import (
	"github.com/saku-730/web-specimen/backend/internal/entity"
	"github.com/saku-730/web-specimen/backend/internal/model"
	"gorm.io/gorm"
)

// DropdownRepository はドロップダウンリストのデータ取得を定義するインターフェースなのだ
type OccurrenceRepository interface {
	GetDropdownLists() (*model.Dropdowns, error)
	CreateOccurrence(tx *gorm.DB, occurrence *entity.Occurrence, classification *entity.ClassificationJSON, place *entity.Place, placeName *entity.PlaceNamesJSON, observation *entity.Observation, specimen *entity.Specimen, makeSpecimen *entity.MakeSpecimen, identification *entity.Identification) (*entity.Occurrence, error)
	Search(query *model.SearchQuery) ([]entity.Occurrence, int64, error)
	FindByID(id uint) (*entity.Occurrence, error)
}

type occurrenceRepository struct {
	db *gorm.DB
}

// NewOccurrenceRepository は新しいリポジトリのインスタンスを作成するのだ
func NewOccurrenceRepository(db *gorm.DB) OccurrenceRepository {
	return &occurrenceRepository{db: db}
}

// GetDropdownLists は各テーブルからリスト作成に必要な情報を取得してくるのだ
func (r *occurrenceRepository) GetDropdownLists() (*model.Dropdowns, error) {
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
	response := &model.Dropdowns{
		Users:              users,
		Projects:           projects,
		Languages:          languages,
		ObservationMethods: obsMethods,
		SpecimenMethods:    specMethods,
		Institutions:       institutions,
	}

	return response, nil
}

func (r *occurrenceRepository) CreateOccurrence(tx *gorm.DB, occurrence *entity.Occurrence, classification *entity.ClassificationJSON, place *entity.Place, placeName *entity.PlaceNamesJSON, observation *entity.Observation, specimen *entity.Specimen, makeSpecimen *entity.MakeSpecimen, identification *entity.Identification) (*entity.Occurrence, error) {
	if err := tx.Create(classification).Error; err != nil { return nil, err }

	if err := tx.Create(placeName).Error; err != nil { return nil, err }
	place.PlaceNameID = &placeName.PlaceNameID
	if err := tx.Create(place).Error; err != nil { return nil, err }

	occurrence.ClassificationID = &classification.ClassificationID
	occurrence.PlaceID = &place.PlaceID
	if err := tx.Create(occurrence).Error; err != nil { return nil, err }

	observation.OccurrenceID = &occurrence.OccurrenceID
	if err := tx.Create(observation).Error; err != nil { return nil, err }

	if specimen != nil && makeSpecimen != nil {
		specimen.OccurrenceID = &occurrence.OccurrenceID
		if err := tx.Create(specimen).Error; err != nil { return nil, err }
		makeSpecimen.OccurrenceID = &occurrence.OccurrenceID
		makeSpecimen.SpecimenID = &specimen.SpecimenID
		if err := tx.Create(makeSpecimen).Error; err != nil { return nil, err }
	}

	if identification != nil {
		identification.OccurrenceID = &occurrence.OccurrenceID
		if err := tx.Create(identification).Error; err != nil { return nil, err }
	}

	return occurrence, nil
}


// Searchメソッドを実装
func (r *occurrenceRepository) Search(query *model.SearchQuery) ([]entity.Occurrence, int64, error) {
	var occurrences []entity.Occurrence
	var total int64

	// ベースとなるクエリを組み立てる。
	tx := r.db.Model(&entity.Occurrence{}).
		Joins("LEFT JOIN users ON users.user_id = occurrence.user_id").
		Joins("LEFT JOIN projects ON projects.project_id = occurrence.project_id").
		Joins("LEFT JOIN places ON places.place_id = occurrence.place_id").
		Joins("LEFT JOIN place_names_json ON place_names_json.place_name_id = places.place_name_id").
		Joins("LEFT JOIN classification_json ON classification_json.classification_id = occurrence.classification_id").
		Joins("LEFT JOIN observations ON observations.occurrence_id = occurrence.occurrence_id").
		Joins("LEFT JOIN make_specimen ON make_specimen.occurrence_id = occurrence.occurrence_id").
		Joins("LEFT JOIN specimen ON specimen.specimen_id = make_specimen.specimen_id").
		Joins("LEFT JOIN identifications ON identifications.occurrence_id = occurrence.occurrence_id")


	// --- クエリパラメータに基づいて動的にWHERE句を追加 ---
	if query.UserID != "" { tx = tx.Where("occurrence.user_id = ?", query.UserID) }
	if query.OccurrenceID != "" { tx = tx.Where("occurrence.occurrence_id = ?", query.OccurrenceID) }
	if query.ProjectID != "" { tx = tx.Where("occurrence.project_id = ?", query.ProjectID) }
	if query.IndividualID != "" { tx = tx.Where("occurrence.individual_id = ?", query.IndividualID) }
	if query.Lifestage != "" { tx = tx.Where("occurrence.lifestage LIKE ?", "%"+query.Lifestage+"%") }
	if query.Sex != "" { tx = tx.Where("occurrence.sex LIKE ?", "%"+query.Sex+"%") }
	if query.BodyLength != "" { tx = tx.Where("occurrence.body_length = ?", query.BodyLength) }
	if query.Note != "" { tx = tx.Where("occurrence.note LIKE ?", "%"+query.Note+"%") }
	if query.CreatedStart != "" && query.CreatedEnd != "" { tx = tx.Where("occurrence.created_at BETWEEN ? AND ?", query.CreatedStart, query.CreatedEnd) }
	if query.PlaceName != "" { tx = tx.Where("place_names_json.class_place_name ->> 'name' LIKE ?", "%"+query.PlaceName+"%") }
	if query.Species != "" { tx = tx.Where("classification_json.class_classification ->> 'species' LIKE ?", "%"+query.Species+"%") }
	if query.Genus != "" { tx = tx.Where("classification_json.class_classification ->> 'genus' LIKE ?", "%"+query.Genus+"%") }
	if query.Family != "" { tx = tx.Where("classification_json.class_classification ->> 'family' LIKE ?", "%"+query.Family+"%") }
	if query.Order != "" { tx = tx.Where("classification_json.class_classification ->> 'order' LIKE ?", "%"+query.Order+"%") }
	if query.Class != "" { tx = tx.Where("classification_json.class_classification ->> 'class' LIKE ?", "%"+query.Class+"%") }
	if query.Phylum != "" { tx = tx.Where("classification_json.class_classification ->> 'phylum' LIKE ?", "%"+query.Phylum+"%") }
	if query.Kingdom != "" { tx = tx.Where("classification_json.class_classification ->> 'kingdom' LIKE ?", "%"+query.Kingdom+"%") }
	if query.Others != "" { tx = tx.Where("classification_json.class_classification ->> 'others' LIKE ?", "%"+query.Others+"%") }
	if query.ObservationUserID != "" { tx = tx.Where("observations.user_id = ?", query.ObservationUserID) }
	if query.ObservationMethodID != "" { tx = tx.Where("observations.observation_method_id = ?", query.ObservationMethodID) }
	if query.ObservedStart != "" && query.ObservedEnd != "" { tx = tx.Where("observations.observed_at BETWEEN ? AND ?", query.ObservedStart, query.ObservedEnd) }
	if query.Behavior != "" { tx = tx.Where("observations.behavior LIKE ?", "%"+query.Behavior+"%") }
	if query.SpecimenUserID != "" { tx = tx.Where("make_specimen.user_id = ?", query.SpecimenUserID) }
	if query.SpecimenMethodsID != "" { tx = tx.Where("specimen.specimen_method_id = ?", query.SpecimenMethodsID) }
	if query.InstitutionID != "" { tx = tx.Where("specimen.institution_id = ?", query.InstitutionID) }
	if query.CollectionID != "" { tx = tx.Where("specimen.collection_id LIKE ?", "%"+query.CollectionID+"%") }
	if query.IdentificationUserID != "" { tx = tx.Where("identifications.user_id = ?", query.IdentificationUserID) }
	if query.IdentifiedStart != "" && query.IdentifiedEnd != "" { tx = tx.Where("identifications.identificated_at BETWEEN ? AND ?", query.IdentifiedStart, query.IdentifiedEnd) }
	

	//if err := tx.Count(&total).Error; err != nil {
	//	return nil, 0, err
	//}

	countTx := tx.Session(&gorm.Session{})
	if err := countTx.Select("COUNT(DISTINCT occurrence.occurrence_id)").Count(&total).Error; err != nil {
	    return nil, 0, err
	}

	//pagenation
	offset := (query.Page - 1) * query.PerPage
	dataTx := tx.Session(&gorm.Session{})
	err := dataTx.Limit(query.PerPage).Offset(offset).
		Preload("User").
		Preload("Project").
		Preload("Place.PlaceNamesJSON").
		Preload("ClassificationJSON").
		Preload("Observations.User").
		Preload("Observations.ObservationMethod").
		Preload("Specimens.SpecimenMethod").
		Preload("Specimens.InstitutionIDCode").
		Preload("MakeSpecimens.User").
		Preload("Identifications.User").
		Order("occurrence.occurrence_id DESC").
		Find(&occurrences).Error

	return occurrences, total, err
}


func (r *occurrenceRepository) FindByID(id uint) (*entity.Occurrence, error) {
	var occurrence entity.Occurrence

	err := r.db.
		Preload("User").
		Preload("Project").
		Preload("Place.PlaceNamesJSON").
		Preload("ClassificationJSON").
		Preload("Observations.User").
		Preload("Observations.ObservationMethod").
		Preload("Specimens.SpecimenMethod").
		Preload("Specimens.InstitutionIDCode").
		Preload("MakeSpecimens.User").
		Preload("Identifications.User").
		Preload("AttachmentGroups.Attachment"). // 中間テーブル経由でAttachmentを取得
		First(&occurrence, id).Error

	return &occurrence, err
}
