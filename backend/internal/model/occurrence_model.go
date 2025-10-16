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
	UserID         uint                  `json:"user_id"`
	ProjectID      *uint                  `json:"project_id"`
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

type ObservationCreate struct {
	ObservationUserID   *uint       `json:"observation_user_id"`
	ObservationMethodID *uint       `json:"observation_method_id"`
	Behavior            *string    `json:"behavior"`
	ObservedAt          *time.Time `json:"observed_at"`
}

type SpecimenCreate struct {
	SpecimenUserID    *uint       `json:"specimen_user_id"`
	SpecimenMethodsID *uint       `json:"specimen_methods_id"`
	CreatedAt         *time.Time `json:"created_at"`
	InstitutionID     *uint       `json:"institution_id"`
	CollectionID      *string    `json:"collection_id"`
}

type IdentificationCreate struct {
	IdentificationUserID *uint       `json:"identification_user_id"`
	IdentifiedAt         *time.Time `json:"identified_at"`
	SourceInfo           *string    `json:"source_info"`
}


// --- Occurrence Detail for /occurrence/{occurrence_id}
type OccurrenceDetailResponse struct {
	UserID         uint                    `json:"user_id"`
	UserName       string                 `json:"user_name"`
	ProjectID      *uint                    `json:"project_id"`
	ProjectName    *string                 `json:"project_name"`
	IndividualID   *int                   `json:"individual_id,omitempty"`
	Lifestage      *string                `json:"lifestage,omitempty"`
	Sex            *string                `json:"sex,omitempty"`
	BodyLength     *string                `json:"body_length,omitempty"`
	CreatedAt      time.Time              `json:"created_at"`
	LanguageID     *uint                   `json:"language_id,omitempty"`
	Latitude       *float64               `json:"latitude,omitempty"`
	Longitude      *float64               `json:"longitude,omitempty"`
	PlaceName      *string                 `json:"place_name,omitempty"`
	Note           *string                `json:"note,omitempty"`
	Classification *ClassificationDetail  `json:"classification,omitempty"`
	Observations   []ObservationDetail    `json:"observation"`   // ⬅️ リスト形式
	Specimens      []SpecimenDetail       `json:"specimen"`      // ⬅️ リスト形式
	Identifications []IdentificationDetail `json:"identification"` // ⬅️ リスト形式
	Attachments    []AttachmentDetail     `json:"attachments"`   // ⬅️ リスト形式
}

type ClassificationDetail struct {
	ClassificationID *uint   `json:"classification_id"`
	Species          *string `json:"species,omitempty"`
	Genus            *string `json:"genus,omitempty"`
	Family           *string `json:"family,omitempty"`
	Order            *string `json:"order,omitempty"`
	Class            *string `json:"class,omitempty"`
	Phylum           *string `json:"phylum,omitempty"`
	Kingdom          *string `json:"kingdom,omitempty"`
	Others           *string `json:"others,omitempty"`
}

type ObservationDetail struct {
	ObservationID         *uint      `json:"observation_id"`
	ObservationUserID     *uint       `json:"observation_user_id"`
	ObservationUser       *string    `json:"observation_user"`
	ObservationMethodID   *uint       `json:"observation_method_id"`
	ObservationMethodName *string    `json:"observation_method_name"`
	PageID                *uint      `json:"page_id,omitempty"`
	Behavior              *string   `json:"behavior,omitempty"`
	ObservedAt            *time.Time `json:"observed_at"`
}

type SpecimenDetail struct {
	SpecimenID            *uint      `json:"specimen_id"`
	SpecimenUserID        *uint       `json:"specimen_user_id"`
	SpecimenUser          *string    `json:"specimen_user"`
	SpecimenMethodsID     *uint       `json:"specimen_methods_id"`
	SpecimenMethodsCommon *string    `json:"specimen_methods_common"`
	CreatedAt             *time.Time `json:"created_at"`
	PageID                *uint      `json:"page_id,omitempty"`
	InstitutionID         *uint       `json:"institution_id"`
	InstitutionCode       *string    `json:"institution_code"`
	CollectionID          *string   `json:"collection_id,omitempty"`
}

type IdentificationDetail struct {
	IdentificationID     *uint      `json:"identification_id"`
	IdentificationUserID *uint       `json:"identification_user_id"`
	IdentificationUser   *string    `json:"identification_user"`
	IdentifiedAt         *time.Time `json:"identified_at"`
	SourceInfo           *string   `json:"source_info,omitempty"`
}

type AttachmentDetail struct {
	AttachmentID *uint   `json:"attachment_id"`
	FilePath     *string `json:"file_path"`
	FileName     *string `json:"file_name"`
	Note	     *string `json:"note"`
}
