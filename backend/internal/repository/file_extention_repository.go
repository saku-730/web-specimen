//internal/repository/file_extension_repository.go 
package repository

import (
	"errors"
	"github.com/saku-730/web-specimen/backend/internal/entity"
	"gorm.io/gorm"
)

type FileExtensionRepository interface {
	FindByText(tx *gorm.DB, extText string) (*entity.FileExtension, error)
}

type fileExtensionRepository struct{}

func NewFileExtensionRepository() FileExtensionRepository {
	return &fileExtensionRepository{}
}

func (r *fileExtensionRepository) FindByText(tx *gorm.DB, extText string) (*entity.FileExtension, error) {
	var fileExtension entity.FileExtension
	err := tx.Where("extension_text = ?", extText).First(&fileExtension).Error
	
	// エラーがあった場合だけ、中身をチェックする
	if err != nil {
		// エラーが「レコードが見つからない」エラーかどうかを判定するのだ
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 見つからなかった場合は、エラーではないので nil, nil を返す
			return nil, nil
		}
		// それ以外の、本当に問題があるエラーの場合は、エラーを上に伝える
		return nil, err
	}
	
	// エラーがなければ、見つかったentityを返す
	return &fileExtension, nil
}
