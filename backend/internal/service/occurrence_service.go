// internal/service/occurrence_service.go
package service

import (
	"errors" // エラーハンドリングで使う
	"gorm.io/gorm"
	"fmt"
	"encoding/json"
	"time"
	"os"
	"mime/multipart"
	"path/filepath"
	"io"
	"strings"

	"github.com/saku-730/web-specimen/backend/internal/entity"
	"github.com/saku-730/web-specimen/backend/internal/model"
	"github.com/saku-730/web-specimen/backend/internal/repository"
)

// OccurrenceServiceのインターフェース。役割をはっきり分けたのだ。
type OccurrenceService interface {
	PrepareCreatePage() (*model.Dropdowns, error)
	GetDefaultValues(userID int) (*model.DefaultValues, error)
	CreateOccurrence(req *model.OccurrenceCreate)(*entity.Occurrence, error)
	AttachFiles (occurrenceID uint, userID uint, files []*multipart.FileHeader) ([]string, error)
}

// occurrenceService構造体。必要なリポジトリを全部持たせるのだ。
type occurrenceService struct {
	db	     *gorm.DB
	occRepo      repository.OccurrenceRepository
	defaultsRepo repository.UserDefaultsRepository
	attachmentRepo    repository.AttachmentRepository
	attachmentGroupRepo repository.AttachmentGroupRepository
	fileExtRepo	repository.FileExtensionRepository
}

// NewOccurrenceService は、必要なリポジトリを全部引数で受け取るのだ！
func NewOccurrenceService(
	db	*gorm.DB,
	occRepo repository.OccurrenceRepository,
	defaultsRepo repository.UserDefaultsRepository,
	attRepo repository.AttachmentRepository, 
	attGroupRepo repository.AttachmentGroupRepository,
	fileExtRepo	repository.FileExtensionRepository
) OccurrenceService {
	return &occurrenceService{
		db:	      db,
		occRepo:      occRepo,
		defaultsRepo: defaultsRepo,
		attachmentRepo: attRepo,
		attachmentGroupRepo: attGroupRepo,
		fileExtRepo: fileExtRepo,
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


func (s *occurrenceService) CreateOccurrence(req *model.OccurrenceCreate) (*entity.Occurrence, error) {

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

	var createdOccurrence *entity.Occurrence

	// --- 2. トランザクションを開始してRepositoryを呼び出す ---
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var err error
		createdOccurrence, err = s.occRepo.CreateOccurrence(tx, occurrence, classification, place, placeName, observation, specimen, makeSpecimen, identification)
		return err
	})

	if err != nil {
		return nil, err
	}

	return createdOccurrence, nil
}


func (s *occurrenceService) AttachFiles(occurrenceID uint, userID uint, files []*multipart.FileHeader) ([]string, error) {
	//prepare dir
	uploadDir := os.Getenv("UPLOAD_DIR")
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return nil, err
	}

	var savedFileNames []string

	// --- save file and file info to database ---
	err := s.db.Transaction(func(tx *gorm.DB) error {
		for _, fileHeader := range files {
			// --- ファイルをサーバーに保存 ---
			// 衝突を避けるためにユニークなファイル名を生成
			uniqueFileName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), fileHeader.Filename)
			destPath := filepath.Join(uploadDir, uniqueFileName)
			
			src, err := fileHeader.Open()
			if err != nil { return err }
			defer src.Close()

			dst, err := os.Create(destPath)
			if err != nil { return err }
			defer dst.Close()

			if _, err := io.Copy(dst, src); err != nil { return err }

			ext := filepath.Ext(fileHeader.Filename)
			// unified to lowercase
			ext = strings.ToLower(ext)

			var extensionID *int
	
			if ext != "" {
				fileExtEntity, err := s.fileExtRepo.FindByText(tx, ext)
				// gorm.ErrRecordNotFound の場合は、見つからなかっただけなので処理を続ける
				// それ以外のDBエラーの場合は、トランザクションを失敗させる
				if err != nil && err != gorm.ErrRecordNotFound {
					return err
				}
				// もし見つかったら、IDをセットする
				if fileExtEntity != nil {
					extensionID = &fileExtEntity.ExtensionID
				}
			}
			
			// --- save to database ---
			//attachment table
			attachment := &entity.Attachment{
				FilePath:         destPath,
				OriginalFilename: &fileHeader.Filename,
				ExtensionID:      extensionID, 
				UserID:           &userID,
				Uploaded:         time.Now(),
			}
			//Repository 
			if err := s.attachmentRepo.Create(tx, attachment); err != nil {
				return err
			}

			//attachment group table
			group := &entity.AttachmentGroup{
				OccurrenceID: occurrenceID,
				AttachmentID: attachment.AttachmentID,
			}
			//Repository 
			if err := s.attachmentGroupRepo.Create(tx, group); err != nil {
				return err
			}
			
			savedFileNames = append(savedFileNames, uniqueFileName)
			}
			if err := s.attachmentRepo.Create(tx, attachment); err != nil {
				return err


		}
		return nil
	})


	if err != nil {
		return nil, err
	}

	return savedFileNames, nil
}
