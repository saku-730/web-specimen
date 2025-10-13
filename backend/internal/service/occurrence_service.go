// internal/service/occurrence_service.go
package service

import (
	"errors" // エラーハンドリングで使う
	"gorm.io/gorm"
	"fmt"
	"encoding/json"
	"time"

	"github.com/saku-730/web-specimen/backend/internal/entity"
	"github.com/saku-730/web-specimen/backend/internal/model"
	"github.com/saku-730/web-specimen/backend/internal/repository"
)

// OccurrenceServiceのインターフェース。役割をはっきり分けたのだ。
type OccurrenceService interface {
	PrepareCreatePage() (*model.Dropdowns, error)
	GetDefaultValues(userID int) (*model.DefaultValues, error)
	CreateOccurrence(req *model.OccurrenceCreate)(*model.OccurrenceCreate, error)
}

// occurrenceService構造体。必要なリポジトリを全部持たせるのだ。
type occurrenceService struct {
	db	     *gorm.DB
	occRepo      repository.OccurrenceRepository
	defaultsRepo repository.UserDefaultsRepository
}

// NewOccurrenceService は、必要なリポジトリを全部引数で受け取るのだ！
func NewOccurrenceService(
	occRepo repository.OccurrenceRepository,
	defaultsRepo repository.UserDefaultsRepository,
) OccurrenceService {
	return &occurrenceService{
		occRepo:     occRepo,
		defaultsRepo: defaultsRepo,
	}
}

// PrepareCreatePage get dropdown list for create page
func (s *occurrenceService) PrepareCreatePage() (*model.Dropdowns, error) {
	return s.occRepo.GetDropdownLists()
}

// GetDefaultValues は、DBから取得したフラットなデータをネストされたモデルに組み立てるのだ！
func (s *occurrenceService) GetDefaultValues(userID int) (*model.DefaultValues, error) {
	// 1. リポジトリからフラットなentity(部品)を取得
	entity, err := s.defaultsRepo.FindDefaultsByUserID(userID)
	if err != nil {
		// もしユーザーのデフォルト設定がDBに無かったら、空っぽのデフォルト値を返す
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.DefaultValues{UserID: userID}, nil // userIDだけセットして返す
		}
		return nil, err // それ以外のDBエラー
	}

	response := &model.DefaultValues{
		UserID:         entity.UserID,
		UserName:       entity.UserName, // ポインタ型なのでデリファレンスするのだ
		ProjectID:      entity.ProjectID,
		ProjectName:    entity.ProjectName,
		IndividualID:   entity.IndividualID,
		Lifestage:      entity.Lifestage,
		Sex:            entity.Sex,
		LanguageID:     entity.LanguageID,
		LanguageCommon: entity.LanguageCommon,
		PlaceName:      entity.PlaceName,
		Note:           entity.Note,
		Classification: model.Classification{
			Species: entity.ClassificationSpecies,
			Genus:   entity.ClassificationGenus,
			Family:  entity.ClassificationFamily,
			Order:   entity.ClassificationOrder,
			Class:   entity.ClassificationClass,
			Phylum:  entity.ClassificationPhylum,
			Kingdom: entity.ClassificationKingdom,
			Others:  entity.ClassificationOthers,
		},
		Observation: model.Observation{
			ObservationUserID:     entity.ObservationUserID,
			ObservationUser:       entity.ObservationUserName,
			ObservationMethodID:   entity.ObservationMethodID,
			ObservationMethodName: entity.ObservationMethodName,
			Behavior:              entity.ObservationBehavior,
			ObservedAt:            entity.ObservationObservedAt,
		},
		Specimen: model.Specimen{
			SpecimenUserID:          entity.SpecimenUserID,
			SpecimenUser:            entity.SpecimenUserName,
			SpecimenMethodsID:       entity.SpecimenMethodID,
			SpecimenMethodsCommon: entity.SpecimenMethodName,
		},
		Identification: model.Identification{
			IdentificationUserID: entity.IdentificationUserID,
			IdentificationUser:   entity.IdentificationUserName,
			IdentifiedAt:         entity.IdentificationIdentifiedAt,
			SourceInfo:           entity.IdentificationSourceInfo,
		},
	}

	return response, nil
}

func formatTimezone(t *time.Time) *string {
	if t == nil {
		return nil // 元がnilなら、nilを返す
	}
	_, offsetInSeconds := t.Zone()
	sign := "+"
	if offsetInSeconds < 0 {
		sign = "-"
		offsetInSeconds = -offsetInSeconds
	}
	hours := offsetInSeconds / 3600
	minutes := (offsetInSeconds % 3600) / 60
	timezoneStr := fmt.Sprintf("%s%02d:%02d", sign, hours, minutes)
	return &timezoneStr // 文字列へのポインタを返す
}


func (s *occurrenceService) CreateOccurrence(req *model.OccurrenceCreate) (*model.OccurrenceCreate, error) {

	classMap := map[string]interface{}{"species": req.Classification.Species, "genus": req.Classification.Genus, "family": req.Classification.Family, "order": req.Classification.Order, "class": req.Classification.Class, "phylum": req.Classification.Phylum, "kingdom": req.Classification.Kingdom, "others": req.Classification.Others}
	classJSON, _ := json.Marshal(classMap)
	classification := &entity.ClassificationJSON{ClassClassification: classJSON}

	placeNameMap := map[string]interface{}{"name": req.PlaceName}
	placeNameJSON, _ := json.Marshal(placeNameMap)
	placeName := &entity.PlaceNamesJSON{ClassPlaceName: placeNameJSON}
	place := &entity.Place{Coordinates: &entity.Point{Lat: req.Latitude, Lng: req.Longitude}}

	occurrence := &entity.Occurrence{
		ProjectID:    req.ProjectID,
		UserID:       &req.UserID,
		IndividualID: req.IndividualID,
		Lifestage:    req.Lifestage,
		Sex:          req.Sex,
		BodyLength:   req.BodyLength, 
		LanguageID:   req.LanguageID,
		Note:         req.Note,
		CreatedAt:    req.CreatedAt,  
		Timezone:     formatTimezone(req.CreatedAt),
	}
	
	observation := &entity.Observation{
		UserID: req.Observation.ObservationUserID, 
		ObservationMethodID: req.Observation.ObservationMethodID, 
		Behavior: req.Observation.Behavior, 
		ObservedAt: req.Observation.ObservedAt, 
		Timezone: formatTimezone(req.Observation.ObservedAt),
	}
	
	specimen := &entity.Specimen{
		SpecimenMethodID: req.Specimen.SpecimenMethodsID, 
		InstitutionID: req.Specimen.InstitutionID,
		CollectionID: req.Specimen.CollectionID,
	}

	makeSpecimen := &entity.MakeSpecimen{
		UserID: req.Specimen.SpecimenUserID, 
		SpecimenMethodID: req.Specimen.SpecimenMethodsID, 
		Date: req.Specimen.CreatedAt, 
		Timezone: formatTimezone(req.Specimen.CreatedAt),
	}

	identification := &entity.Identification{
		UserID: req.Identification.IdentificationUserID, 
		SourceInfo: req.Identification.SourceInfo, 
		IdentificatedAt: req.Identification.IdentifiedAt, 
		Timezone: formatTimezone(req.Identification.IdentifiedAt),
	}

	// --- 2. トランザクションを開始してRepositoryを呼び出す ---
	err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.occRepo.CreateOccurrence(tx, occurrence, classification, place, placeName, observation, specimen, makeSpecimen, identification)
	})

	if err != nil {
		return nil, err
	}

	return req, nil
}
