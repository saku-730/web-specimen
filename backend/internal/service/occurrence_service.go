// internal/service/occurrence_service.go
package service

import (
//	"mime/multipart"

	"github.com/saku-730/web-specimen/backend/internal/model"
	"github.com/saku-730/web-specimen/backend/internal/repository"
)

// OccurrenceService のインターフェース定義なのだ
// ハンドラが期待するメソッドを全て書いておくのだ
type OccurrenceService interface {
	PrepareCreatePage() (*model.Dropdowns, error)
	GetDefaultValue()(*model.DefaultValues,error)
}

// occurrenceService 構造体。リポジトリを持つ形だけ作っておくのだ
type occurrenceService struct {
	occRepo repository.OccurrenceRepository
}

// NewOccurrenceService は空のサービスを返すコンストラクタなのだ
func NewOccurrenceService(occRepo repository.OccurrenceRepository) OccurrenceService {
	return &occurrenceService{occRepo: occRepo}
}

func (s *occurrenceService) PrepareCreatePage() (*model.Dropdowns, error) {
	// TODO: あとで実装するのだ
	return nil, nil
}

func (s *occurrenceService) GetDefaultValue() (*model.DefaultValues, error) {
	// TODO: あとで実装するのだ
	return nil, nil
}

