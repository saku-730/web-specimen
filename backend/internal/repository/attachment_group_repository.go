//internal/repository/attachment_group_repository.go
package repository

import (
	"github.com/saku-730/web-specimen/backend/internal/entity"
	"gorm.io/gorm"
)

type AttachmentGroupRepository interface {
	Create(tx *gorm.DB, group *entity.AttachmentGroup) error
}

type attachmentGroupRepository struct{}

func NewAttachmentGroupRepository() AttachmentGroupRepository {
	return &attachmentGroupRepository{}
}

func (r *attachmentGroupRepository) Create(tx *gorm.DB, group *entity.AttachmentGroup) error {
	return tx.Create(group).Error
}
