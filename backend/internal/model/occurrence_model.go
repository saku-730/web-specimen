// backend/internal/model/occurrence_model.go
package model

import (
	"mime/multipart"
	"time"
)


// --- Dropdowns for create paga ---
type Dropdowns struct {
	Users		  []User	      `json:"users"`
	Projects          []Project           `json:"projects"`
	Languages         []Language          `json:"languages"`
	ObservationMethods []ObservationMethod `json:"observation_methods"`
	SpecimenMethods   []SpecimenMethod    `json:"specimen_methods"`
	Institutions      []Institution       `json:"institutions"`
}

// User choice
type User struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

// Project choice
type Project struct {
	ProjectID   int    `json:"project_id"`
	ProjectName string `json:"project_name"`
}

// Language choice
type Language struct {
	LanguageID     int    `json:"language_id"`
	LanguageCommon string `json:"language_common"`
}

// ObservationMethod choice
type ObservationMethod struct {
	ObservationMethodID   int    `json:"observation_method_id"`
	ObservationMethodName string `json:"observation_method_name"`
}

// SpecimenMethod choice
type SpecimenMethod struct {
	SpecimenMethodsID     int    `json:"specimen_methods_id"`
	SpecimenMethodsCommon string `json:"specimen_methods_common"`
}

// Institution choice
type Institution struct {
	InstitutionID   int    `json:"institution_id"`
	InstitutionCode string `json:"institution_code"`
}



// --- Default values for create paga ---
// DefaultValues 
type DefaultValues struct {
	UserID         int            `json:"user_id"`
	ProjectID      int            `json:"project_id"`
	IndividualID   int            `json:"individual_id"`
	Lifestage      string         `json:"lifestage"`
	Sex            string         `json:"sex"`
	BodyLength     string         `json:"body_length"`
	CreatedAt      string         `json:"created_at"` // サービス層で time.Time に変換する
	LanguageID     int            `json:"language_id"`
	Latitude       float64        `json:"latitude"`
	Longitude      float64        `json:"longitude"`
	PlaceName      string         `json:"place_name"`
	Note           string         `json:"note"`
	Classification Classification `json:"classification"`
	Observation    Observation    `json:"observation"`
	Specimen       Specimen       `json:"specimen"`
	Identification Identification `json:"identification"`
}

type Classification struct {
	ID        int    `json:"classification_id"`
	Species   string `json:"species"`
	Genus     string `json:"genus"`
	Family    string `json:"family"`
	Order     string `json:"order"`
	Class     string `json:"class"`
	Phylum    string `json:"phylum"`
	Kingdom   string `json:"kingdom"`
	Others    string `json:"others"`
}

type Observation struct {
	ID         int    `json:"observation_id"`
	UserID     int    `json:"observation_user_id"`
	MethodID   int    `json:"observation_method_id"`
	MethodName string `json:"observation_method_name"`
	PageID     int    `json:"page_id"`
	Behavior   string `json:"behavior"`
	ObservedAt string `json:"observed_at"`
}

type Specimen struct {
	MethodID   int    `json:"specimen_methods_id"`
	MethodName string `json:"specimen_methods_common"`
	PageID     int    `json:"page_id"`
}

type Identification struct {
	ID         int    `json:"identification_id"`
	UserID     int    `json:"identification_user_id"`
	IdentifiedAt string `json:"identified_at"`
	SourceInfo string `json:"source_info"`
}

