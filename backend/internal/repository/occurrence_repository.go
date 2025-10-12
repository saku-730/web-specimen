// internal/repository/occurrence_repository.go
package repository

import (
	"gorm.io/gorm"
)


type OccurrenceRepository interface {
}

type occurrenceRepository struct {
	db *gorm.DB
}

func NewOccurrenceRepository(db *gorm.DB) OccurrenceRepository {
	return &occurrenceRepository{db: db}
}
