// backend/internal/model/occurrence_model.go
package model

import (
//	"mime/multipart"
	"time"
)

type CreatePageData struct {
    DropdownList Dropdowns `json:"dropdown_list"`
    DefaultValue DefaultValues `json:"default_value"`
}

// --- Dropdowns for create paga ---
type Dropdowns struct {
	Users              []DropdownUser              `json:"users"`
	Projects           []DropdownProject           `json:"projects"`
	Languages          []DropdownLanguage          `json:"languages"`
	ObservationMethods []DropdownObservationMethod `json:"observation_methods"`
	SpecimenMethods    []DropdownSpecimenMethod    `json:"specimen_methods"`
	Institutions       []DropdownInstitution       `json:"institutions"`
}

type DropdownUser struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
}

type DropdownProject struct {
	ProjectID   uint   `json:"project_id"`
	ProjectName string `json:"project_name"`
}

type DropdownLanguage struct {
	LanguageID     uint   `json:"language_id"`
	LanguageCommon string `json:"language_common"`
}

type DropdownObservationMethod struct {
	ObservationMethodID   uint   `json:"observation_method_id"`
	ObservationMethodName string `json:"observation_method_name"`
}

type DropdownSpecimenMethod struct {
	SpecimenMethodsID      uint   `json:"specimen_methods_id"`
	SpecimenMethodsCommon string `json:"specimen_methods_common"`
}

type DropdownInstitution struct {
	InstitutionID   uint   `json:"institution_id"`
	InstitutionCode string `json:"institution_code"`
}

// --- Default values for create paga ---
// DefaultValue
type DefaultValues struct {
	UserID         int            `json:"user_id"`
	UserName       *string         `json:"user_name"`
	ProjectID      *int           `json:"project_id"`
	ProjectName    *string        `json:"project_name"`
	IndividualID   *int           `json:"individual_id"`
	Lifestage      *string        `json:"lifestage"`
	Sex            *string        `json:"sex"`
	LanguageID     *int           `json:"language_id"`
	LanguageCommon *string        `json:"language_common"`
	PlaceName      *string        `json:"place_name"`
	Note           *string        `json:"note"`
	Classification Classification `json:"classification"`
	Observation    Observation    `json:"observation"`
	Specimen       Specimen       `json:"specimen"`
	Identification Identification `json:"identification"`
}

type Classification struct {
	Species *string `json:"species"`
	Genus   *string `json:"genus"`
	Family  *string `json:"family"`
	Order   *string `json:"order"`
	Class   *string `json:"class"`
	Phylum  *string `json:"phylum"`
	Kingdom *string `json:"kingdom"`
	Others  *string `json:"others"`
}

type Observation struct {
	ObservationUserID     *int    `json:"observation_user_id"`
	ObservationUser       *string `json:"observation_user"`
	ObservationMethodID   *int    `json:"observation_method_id"`
	ObservationMethodName *string `json:"observation_method_name"`
	Behavior              *string `json:"behavior"`
	ObservedAt            *string `json:"observed_at"`
}

type Specimen struct {
	SpecimenUserID          *int    `json:"specimen_user_id"`
	SpecimenUser            *string `json:"specimen_user"`
	SpecimenMethodsID       *int    `json:"specimen_methods_id"`
	SpecimenMethodsCommon *string `json:"specimen_methods_common"`
}

type Identification struct {
	IdentificationUserID *int    `json:"identification_user_id"`
	IdentificationUser   *string `json:"identification_user"`
	IdentifiedAt         *string `json:"identified_at"`
	SourceInfo           *string `json:"source_info"`
}


// --- OccurrenceCreate ---
type OccurrenceCreate struct {
	UserID         int                  `json:"user_id"`
	ProjectID      *int                  `json:"project_id"`
	IndividualID   *int                 `json:"individual_id"`
	Lifestage      *string               `json:"lifestage"`
	Sex            *string               `json:"sex"`
	BodyLength     *string               `json:"body_length"`
	CreatedAt      *time.Time            `json:"created_at"` 
	LanguageID     *uint                  `json:"language_id"`
	Latitude       *float64              `json:"latitude"`
	Longitude      *float64              `json:"longitude"`
	PlaceName      *string               `json:"place_name"`
	Note           *string               `json:"note"`
	Classification *ClassificationCreate `json:"classification"`
	Observation    *ObservationCreate    `json:"observation"`
	Specimen       *SpecimenCreate      `json:"specimen"`      
	Identification *IdentificationCreate `json:"identification"`
}

// ClassificationCreate は分類情報の部分構造なのだ
type ClassificationCreate struct {
	Species *string `json:"species"`
	Genus   *string `json:"genus"`
	Family  *string `json:"family"`
	Order   *string `json:"order"`
	Class   *string `json:"class"`
	Phylum  *string `json:"phylum"`
	Kingdom *string `json:"kingdom"`
	Others  *string `json:"others"`
}

// ObservationCreate は観察情報の部分構造なのだ
type ObservationCreate struct {
	ObservationUserID   *int       `json:"observation_user_id"`
	ObservationMethodID *uint       `json:"observation_method_id"`
	Behavior            *string    `json:"behavior"`
	ObservedAt          *time.Time `json:"observed_at"`
}

// SpecimenCreate は標本情報の部分構造なのだ
type SpecimenCreate struct {
	SpecimenUserID    *uint       `json:"specimen_user_id"`
	SpecimenMethodsID *uint       `json:"specimen_methods_id"`
	CreatedAt         *time.Time `json:"created_at"`
	InstitutionID     *uint       `json:"institution_id"`
	CollectionID      *string    `json:"collection_id"`
}

// IdentificationCreate は同定情報の部分構造なのだ
type IdentificationCreate struct {
	IdentificationUserID *uint       `json:"identification_user_id"`
	IdentifiedAt         *time.Time `json:"identified_at"`
	SourceInfo           *string    `json:"source_info"`
}
