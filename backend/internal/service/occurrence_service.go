// internal/service/occurrence_service.go
package service

import (
	"errors" 
	"gorm.io/gorm"
	"fmt"
	"encoding/json"
	"time"
	"os"
	"mime/multipart"
	"path/filepath"
	"io"
	"strings"
	"math"

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
	Search(query *model.SearchQuery) (*model.SearchResponse, error)
	GetOccurrenceDetail(id uint) (*model.OccurrenceDetailResponse, error)
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
	fileExtRepo	repository.FileExtensionRepository,
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

// PrepareCreatePage get dropdown list for create,search page
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
		UserName:       entity.UserName, // ポインタ型なのでデリファレンスするのだ ProjectID:      entity.ProjectID, ProjectName:    entity.ProjectName,
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
    // --- 1. リクエストDTOを各Entityオブジェクトに変換 ---

	// まず、nilになる可能性のある変数をすべてnilで宣言しておくのだ。これが安全の基本！
	var classification *entity.ClassificationJSON
	var placeName *entity.PlaceNamesJSON
	var place *entity.Place
	var observation *entity.Observation
	var specimen *entity.Specimen
	var makeSpecimen *entity.MakeSpecimen
	var identification *entity.Identification


	// 1. Classification: フロントエンドからデータが送られてきた場合のみ、entityを作成する。
	if req.Classification != nil {
		classMap := map[string]interface{}{
			"species": req.Classification.Species, "genus": req.Classification.Genus, "family": req.Classification.Family, 
			"order": req.Classification.Order, "class": req.Classification.Class, "phylum": req.Classification.Phylum, 
			"kingdom": req.Classification.Kingdom, "others": req.Classification.Others,
		}
		classJSON, _ := json.Marshal(classMap)
		classification = &entity.ClassificationJSON{ClassClassification: classJSON}
	}

	// 2. Place: 場所に関する情報が何か一つでも送られてきた場合のみ、entityを作成する。
	if (req.PlaceName != nil && *req.PlaceName != "") || 
	   (req.Latitude != nil && *req.Latitude != 0) || 
	   (req.Longitude != nil && *req.Longitude != 0) {
		
		var name string
		if req.PlaceName != nil {
			name = *req.PlaceName
		}
		placeNameMap := map[string]interface{}{"name": name}
		placeNameJSON, _ := json.Marshal(placeNameMap)
		placeName = &entity.PlaceNamesJSON{ClassPlaceName: placeNameJSON}

		place = &entity.Place{Coordinates: &entity.Point{Lat: req.Latitude, Lng: req.Longitude}}
	}

	// 3. Observation: データが送られてきた場合のみ、entityを作成する。
	if req.Observation != nil {
		observation = &entity.Observation{
			UserID:              req.Observation.ObservationUserID,
			ObservationMethodID: req.Observation.ObservationMethodID,
			Behavior:            req.Observation.Behavior,
			ObservedAt:          req.Observation.ObservedAt,
			Timezone:            formatTimezone(req.Observation.ObservedAt), // formatTimezone内でnilチェック済み
		}
	}
	
	// 4. Specimen: データが送られてきた場合のみ、entityを作成する。
	if req.Specimen != nil {
		specimen = &entity.Specimen{
			SpecimenMethodID: req.Specimen.SpecimenMethodsID,
			InstitutionID:    req.Specimen.InstitutionID,
			CollectionID:     req.Specimen.CollectionID,
		}
		makeSpecimen = &entity.MakeSpecimen{
			UserID:           req.Specimen.SpecimenUserID,
			SpecimenMethodID: req.Specimen.SpecimenMethodsID,
			Date:             req.Specimen.CreatedAt,
			Timezone:         formatTimezone(req.Specimen.CreatedAt), // formatTimezone内でnilチェック済み
		}
	}

	// 5. Identification: データが送られてきた場合のみ、entityを作成する。
	if req.Identification != nil {
		identification = &entity.Identification{
			UserID:          req.Identification.IdentificationUserID,
			SourceInfo:      req.Identification.SourceInfo,
			IdentificatedAt: req.Identification.IdentifiedAt,
			Timezone:        formatTimezone(req.Identification.IdentifiedAt), // formatTimezone内でnilチェック済み
		}
	}

    // 6. 最後に、必須項目とトップレベルの任意項目でoccurrenceを作る。
    //    reqの各フィールドはポインタなので、そのまま代入すればOK。
	occurrence := &entity.Occurrence{
		ProjectID:    req.ProjectID,
		UserID:       &req.UserID, // UserIDは必須項目なのでポインタではない
		IndividualID: req.IndividualID,
		Lifestage:    req.Lifestage,
		Sex:          req.Sex,
		BodyLength:   req.BodyLength,
		LanguageID:   req.LanguageID,
		Note:         req.Note,
		CreatedAt:    req.CreatedAt,
		Timezone:     formatTimezone(req.CreatedAt), // CreatedAtは必須と仮定
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

			var extensionID *uint
	
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
		
			now := time.Now()
			// --- save to database ---
			//attachment table
			attachment := &entity.Attachment{
				FilePath:         destPath,
				OriginalFilename: &fileHeader.Filename,
				ExtensionID:      extensionID, 
				UserID:           &userID,
				Uploaded:         &now,
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
		return nil
	})


	if err != nil {
		return nil, err
	}

	return savedFileNames, nil
}


// Searchメソッドを実装
func (s *occurrenceService) Search(query *model.SearchQuery) (*model.SearchResponse, error) {
	// ページネーションのデフォルト値を設定
	if query.Page <= 0 { query.Page = 1 }
	if query.PerPage <= 0 { query.PerPage = 30 }

	occurrences, total, err := s.occRepo.Search(query)
	if err != nil {
		return nil, err
	}

	// --- entityからレスポンス用のmodelに変換する ---
	var results []model.OccurrenceResult
	for _, occ := range occurrences {
		result := model.OccurrenceResult{
			UserID:       occ.UserID,
			UserName:     occ.User.UserName,
			ProjectID:    occ.ProjectID,
			ProjectName:  occ.Project.ProjectName,
			IndividualID: occ.IndividualID,
			Lifestage:    occ.Lifestage,
			Sex:          occ.Sex,
			BodyLength:   occ.BodyLength,
			CreatedAt:    occ.CreatedAt,
			LanguageID:   occ.LanguageID,
			Note:         occ.Note,
		}

		if occ.Place != nil && occ.Place.Coordinates != nil {
			result.Latitude = occ.Place.Coordinates.Lat
			result.Longitude = occ.Place.Coordinates.Lng
		}

		if occ.Place != nil && occ.Place.PlaceNamesJSON != nil {
			var placeNameData map[string]string
			if err := json.Unmarshal(occ.Place.PlaceNamesJSON.ClassPlaceName, &placeNameData); err == nil {
				name := placeNameData["name"]
				result.PlaceName = &name
			}
		}

		if occ.ClassificationJSON != nil {
			var classData map[string]string
			if err := json.Unmarshal(occ.ClassificationJSON.ClassClassification, &classData); err == nil {
				species := classData["species"]
				genus := classData["genus"]
				family := classData["family"]
				order := classData["order"]
				class := classData["class"]
				phylum := classData["phylum"]
				kingdom := classData["kingdom"]
				others := classData["others"]


				result.Classification = &model.ClassificationResult{
					ClassificationID: &occ.ClassificationJSON.ClassificationID,
					Species:          &species,
					Genus:            &genus,
					Family:           &family,
					Order:            &order,
					Class:            &class,
					Phylum:           &phylum,
					Kingdom:          &kingdom,
					Others:           &others,
				}
			}
		}

		if len(occ.Observations) > 0 {
			obs := occ.Observations[0] // 代表して最初の1件を取得
			result.Observation = &model.ObservationResult{
				ObservationID:         &obs.ObservationsID,
				ObservationUserID:     obs.UserID,
				ObservationUser:       &obs.User.UserName,
				ObservationMethodID:   obs.ObservationMethodID,
				ObservationMethodName: obs.ObservationMethod.MethodCommonName,
				PageID:                obs.ObservationMethod.PageID,
				Behavior:              obs.Behavior,
				ObservedAt:            obs.ObservedAt,
			}
		}

		if len(occ.Specimens) > 0 && len(occ.MakeSpecimens) > 0 {
			spec := occ.Specimens[0]         // 代表して最初の標本を取得
			makeSpec := occ.MakeSpecimens[0] // 代表して最初の標本作成記録を取得

			result.Specimen = &model.SpecimenResult{
				SpecimenID:            &spec.SpecimenID,
				SpecimenUserID:        makeSpec.UserID,
				SpecimenUser:          &makeSpec.User.UserName, 
				SpecimenMethodsID:     spec.SpecimenMethodID,
				SpecimenMethodsCommon: spec.SpecimenMethod.MethodCommonName, 
				PageID:                spec.SpecimenMethod.PageID,
				InstitutionID:         spec.InstitutionID,
				InstitutionCode:       spec.InstitutionIDCode.InstitutionCode,
				CollectionID:          spec.CollectionID,
			}
		}

		// Identification の情報をマッピングするのだ
		if len(occ.Identifications) > 0 {
			ident := occ.Identifications[0] // 代表して最初の同定記録を取得
			
			result.Identification = &model.IdentificationResult{
				IdentificationID:     &ident.IdentificationID,
				IdentificationUserID: ident.UserID,
				IdentificationUser:   &ident.User.UserName, 
				IdentifiedAt:         ident.IdentificatedAt,
				SourceInfo:           ident.SourceInfo,
			}
		}

		results = append(results, result)
	}

	// メタデータを計算
	totalPages := 0
	if total > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(query.PerPage)))
	}

	// 最終的なレスポンスを組み立てる
	response := &model.SearchResponse{
		Results: results,
		Metadata: model.Metadata{
			TotalResults: int(total),
			CurrentPage:  query.Page,
			PerPage:      query.PerPage,
			TotalPages:   totalPages,
		},
	}

	return response, nil
}

func (s *occurrenceService) GetOccurrenceDetail(id uint) (*model.OccurrenceDetailResponse, error) {
	occ, err := s.occRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// --- entityからレスポンス用のmodelに変換する ---
	response := &model.OccurrenceDetailResponse{
		UserID:       *occ.UserID,
		UserName:     occ.User.UserName,
		ProjectID:    occ.ProjectID,
		ProjectName:  occ.Project.ProjectName,
		IndividualID: occ.IndividualID,
		Lifestage:    occ.Lifestage,
		Sex:          occ.Sex,
		BodyLength:   occ.BodyLength,
		CreatedAt:    *occ.CreatedAt,
		LanguageID:   occ.LanguageID,
		Note:         occ.Note,
	}

	if occ.Place != nil && occ.Place.Coordinates != nil {
		response.Latitude = occ.Place.Coordinates.Lat
		response.Longitude = occ.Place.Coordinates.Lng
	}
	if occ.Place != nil && occ.Place.PlaceNamesJSON != nil {
		var placeNameData map[string]string
		if err := json.Unmarshal(occ.Place.PlaceNamesJSON.ClassPlaceName, &placeNameData); err == nil {
			placeName := placeNameData["name"]
			response.PlaceName = &placeName
		}
	}

	if occ.ClassificationJSON != nil {
		var classData map[string]string
		json.Unmarshal(occ.ClassificationJSON.ClassClassification, &classData)
		
		species := classData["species"]
		genus := classData["genus"]
		family := classData["family"]
		order := classData["order"]
		class := classData["class"]
		phylum := classData["phylum"]
		kingdom := classData["kingdom"]
		others := classData["order"]

		response.Classification = &model.ClassificationDetail{
			ClassificationID: &occ.ClassificationJSON.ClassificationID,
			Species:          &species,
			Genus:      	  &genus,      
			Family:           &family,
			Order:            &order,
			Class:            &class,
			Phylum:           &phylum,
			Kingdom:          &kingdom,
			Others:           &others,
		}
	}

	// Observations (リスト) の変換
	for _, obs := range occ.Observations {
		response.Observations = append(response.Observations, model.ObservationDetail{
			ObservationID:         &obs.ObservationsID,
			ObservationUserID:     obs.UserID,
			ObservationUser:       &obs.User.UserName,
			ObservationMethodID:   obs.ObservationMethodID,
			ObservationMethodName: obs.ObservationMethod.MethodCommonName,
			PageID:                obs.ObservationMethod.PageID,
			Behavior:              obs.Behavior,
			ObservedAt:            obs.ObservedAt,
		})
	}

	// Specimens (リスト) の変換
	for _, spec := range occ.Specimens {
		// make_specimenから対応するレコードを探す
		var makeSpecUser entity.User
		var makeSpecCreatedAt time.Time
		for _, ms := range occ.MakeSpecimens {
			if ms.SpecimenID != nil && *ms.SpecimenID == spec.SpecimenID {
				makeSpecUser = ms.User
				makeSpecCreatedAt = *ms.CreatedAt
				break
			}
		}
		response.Specimens = append(response.Specimens, model.SpecimenDetail{
			SpecimenID:            &spec.SpecimenID,
			SpecimenUserID:        &makeSpecUser.UserID,
			SpecimenUser:          &makeSpecUser.UserName,
			SpecimenMethodsID:     spec.SpecimenMethodID,
			SpecimenMethodsCommon: spec.SpecimenMethod.MethodCommonName,
			CreatedAt:             &makeSpecCreatedAt,
			PageID:                spec.SpecimenMethod.PageID,
			InstitutionID:         spec.InstitutionID,
			InstitutionCode:       spec.InstitutionIDCode.InstitutionCode,
			CollectionID:          spec.CollectionID,
		})
	}

	// Identifications (リスト) の変換
	for _, ident := range occ.Identifications {
		response.Identifications = append(response.Identifications, model.IdentificationDetail{
			IdentificationID:     &ident.IdentificationID,
			IdentificationUserID: ident.UserID,
			IdentificationUser:   &ident.User.UserName,
			IdentifiedAt:         ident.IdentificatedAt,
			SourceInfo:           ident.SourceInfo,
		})
	}
	
	// Attachments (リスト) の変換
	for _, group := range occ.AttachmentGroups {
		if group.Attachment != nil {
			response.Attachments = append(response.Attachments, model.AttachmentDetail{
				AttachmentID: &group.Attachment.AttachmentID,
				FilePath:     &group.Attachment.FilePath,
			})
		}
	}

	return response, nil
}
