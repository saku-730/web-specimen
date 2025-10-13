//internal/repository/attachment_repository.go
package repository

import (
	"github.com/saku-730/web-specimen/backend/internal/entity"
	"gorm.io/gorm"
)

type AttachmentRepository interface {
	Create(tx *gorm.DB, attachment *entity.Attachment) error
}

type attachmentRepository struct{}

func NewAttachmentRepository() AttachmentRepository {
	return &attachmentRepository{}
}

func (r *attachmentRepository) Create(tx *gorm.DB, attachment *entity.Attachment) error {
	return tx.Create(attachment).Error
}
