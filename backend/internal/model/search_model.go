//internal/model/search_model.go
package model

import "time"

type SearchQuery struct {
	// Pagination
	Page    int `form:"page"`
	PerPage int `form:"per_page"`

	// Occurrence
	UserID       string `form:"user_id"`
	OccurrenceID string `form:"occurrence_id"`
	ProjectID    string `form:"project_id"`
	IndividualID string `form:"individual_id"`
	Lifestage    string `form:"lifestage"`
	Sex          string `form:"sex"`
	BodyLength   string `form:"body_length"`
	CreatedStart string `form:"created_start"`
	CreatedEnd   string `form:"created_end"`
	Note         string `form:"note"`

	// Place
	PlaceName string `form:"place_name"`

	// Classification
	Species string `form:"species"`
	Genus   string `form:"genus"`
	Family  string `form:"family"`
	Order   string `form:"order"`
	Class   string `form:"class"`
	Phylum  string `form:"phylum"`
	Kingdom string `form:"kingdom"`
	Others  string `form:"others"`

	// Observation
	ObservationUserID   string `form:"observation_user_id"`
	ObservationMethodID string `form:"observation_method_id"`
	ObservedStart       string `form:"observed_start"`
	ObservedEnd         string `form:"observed_end"`
	Behavior            string `form:"behavior"`

	// Specimen / MakeSpecimen
	SpecimenUserID       string `form:"specimen_user_id"`
	SpecimenMethodsID    string `form:"specimen_methods_id"`
	SpecimenCreatedStart string `form:"specimen_created_start"`
	SpecimenCreatedEnd   string `form:"specimen_created_end"`
	InstitutionID        string `form:"institution_id"`
	CollectionID         string `form:"collection_id"`

	// Identification
	IdentificationUserID string `form:"identification_user_id"`
	IdentifiedStart string `form:"identified_start"`
	IdentifiedEnd   string `form:"identified_end"`
}

// SearchResponse は検索結果のレスポンス全体の構造なのだ
type SearchResponse struct {
	Results  []OccurrenceResult `json:"occurrence_results"`
	Metadata Metadata           `json:"metadata"`
}

// Metadata はページネーション情報の構造なのだ
type Metadata struct {
	TotalResults int `json:"total_results"`
	CurrentPage  int `json:"current_page"`
	PerPage      int `json:"per_page"`
	TotalPages   int `json:"total_pages"`
}

// OccurrenceResult は検索結果の各項目の詳細な構造なのだ
type OccurrenceResult struct {
	UserID         int                   `json:"user_id"`
	UserName       string                `json:"user_name"`
	ProjectID      int                   `json:"project_id"`
	ProjectName    string                `json:"project_name"`
	IndividualID   *int                  `json:"individual_id,omitempty"`
	Lifestage      *string               `json:"lifestage,omitempty"`
	Sex            *string               `json:"sex,omitempty"`
	BodyLength     *string               `json:"body_length,omitempty"`
	CreatedAt      time.Time             `json:"created_at"`
	LanguageID     *int                  `json:"language_id,omitempty"`
	Latitude       *float64              `json:"latitude,omitempty"`
	Longitude      *float64              `json:"longitude,omitempty"`
	PlaceName      string                `json:"place_name,omitempty"`
	Note           *string               `json:"note,omitempty"`
	Classification *ClassificationResult `json:"classification,omitempty"`
	Observation    *ObservationResult    `json:"observation,omitempty"`
	Specimen       *SpecimenResult       `json:"specimen,omitempty"`
	Identification *IdentificationResult `json:"identification,omitempty"`
}

// ClassificationResult は分類情報のレスポンス構造なのだ
type ClassificationResult struct {
	ClassificationID uint   `json:"classification_id"`
	Species          string `json:"species,omitempty"`
	Genus            string `json:"genus,omitempty"`
	Family           string `json:"family,omitempty"`
	Order            string `json:"order,omitempty"`
	Class            string `json:"class,omitempty"`
	Phylum           string `json:"phylum,omitempty"`
	Kingdom          string `json:"kingdom,omitempty"`
	Others           string `json:"others,omitempty"`
}

// ObservationResult は観察情報のレスポンス構造なのだ
type ObservationResult struct {
	ObservationID         uint      `json:"observation_id"`
	ObservationUserID     uint       `json:"observation_user_id"`
	ObservationUser       string    `json:"observation_user"`
	ObservationMethodID   uint       `json:"observation_method_id"`
	ObservationMethodName string    `json:"observation_method_name"`
	PageID                *int      `json:"page_id,omitempty"`
	Behavior              *string   `json:"behavior,omitempty"`
	ObservedAt            time.Time `json:"observed_at"`
}

// SpecimenResult は標本情報のレスポンス構造なのだ
type SpecimenResult struct {
	SpecimenID            uint    `json:"specimen_id"`
	SpecimenUserID        uint     `json:"specimen_user_id"`
	SpecimenUser          string  `json:"specimen_user"`
	SpecimenMethodsID     uint     `json:"specimen_methods_id"`
	SpecimenMethodsCommon string  `json:"specimen_methods_common"`
	PageID                *uint    `json:"page_id,omitempty"`
	InstitutionID         uint     `json:"institution_id"`
	InstitutionCode       string  `json:"institution_code"`
	CollectionID          *string `json:"collection_id,omitempty"`
}

// IdentificationResult は同定情報のレスポンス構造なのだ
type IdentificationResult struct {
	IdentificationID     uint      `json:"identification_id"`
	IdentificationUserID int       `json:"identification_user_id"`
	IdentificationUser   string    `json:"identification_user"`
	IdentifiedAt         time.Time `json:"identified_at"`
	SourceInfo           *string   `json:"source_info,omitempty"`
}
