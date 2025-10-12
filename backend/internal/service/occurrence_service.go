// internal/service/occurrence_service.go
package service

import(
//	"mime/multipart"
	"github.com/saku-730/web-specimen/backend/internal/model"
)

type OccurrenceService interface {
	PrepareCreatePage() (*model.Dropdowns, error)
	GetDefaultValue() (*model.DefaultValues, error)
//	CreateOccurrence(req *model.OccurrenceCreate) (*model.Occurrence, error)
//	UploadAttachments(occurrenceID uint, files []*multipart.FileHeader) ([]model.UploadAttachmentInfo, error)
}
