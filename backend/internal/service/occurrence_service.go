// internal/service/occurrence_service.go
package service

import (
	"errors" // エラーハンドリングで使うのだ
	"gorm.io/gorm"

//	"github.com/saku-730/web-specimen/backend/internal/entity"
	"github.com/saku-730/web-specimen/backend/internal/model"
	"github.com/saku-730/web-specimen/backend/internal/repository"
)

// OccurrenceServiceのインターフェース。役割をはっきり分けたのだ。
type OccurrenceService interface {
	PrepareCreatePage() (*model.Dropdowns, error)
	GetDefaultValues(userID int) (*model.DefaultValues, error)
}

// occurrenceService構造体。必要なリポジトリを全部持たせるのだ。
type occurrenceService struct {
	dropRepo     repository.DropdownRepository
	defaultsRepo repository.UserDefaultsRepository
}

// NewOccurrenceService は、必要なリポジトリを全部引数で受け取るのだ！
func NewOccurrenceService(
	dropRepo repository.DropdownRepository,
	defaultsRepo repository.UserDefaultsRepository,
) OccurrenceService {
	return &occurrenceService{
		dropRepo:     dropRepo,
		defaultsRepo: defaultsRepo,
	}
}

// PrepareCreatePage get dropdown list for create page
func (s *occurrenceService) PrepareCreatePage() (*model.DropdownListResponse, error) {
	return s.dropRepo.GetDropdownLists()
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

	response := &model.DefaultValueResponse{
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

