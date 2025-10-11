// backend/internal/handler/occurence_handler_test.go
package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/saku-730/web-specimen/backend/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- ステップ1: サービスのモックを作るのだ ---

// service.OccurrenceService インターフェースを満たす偽物の構造体を定義する
type mockOccurrenceService struct {
	mock.Mock // testify/mock のおまじない
}

// インターフェースのメソッドを全部実装するのだ
func (m *mockOccurrenceService) PrepareCreatePage() (*model.Dropdowns, error) {
	// m.Called() で「このメソッドが呼ばれたよ」と記録する
	// On() で設定した返り値を ret.Get() で取り出すのだ
	ret := m.Called()
	
	// 1つ目の返り値 (Dropdowns) を取り出す
	// nil チェックをして安全に取り出すのだ
	var r0 *model.Dropdowns
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.Dropdowns)
	}

	// 2つ目の返り値 (error) を取り出す
	return r0, ret.Error(1)
}

func (m *mockOccurrenceService) GetDefaultValue() (*model.DefaultValues, error) {
	ret := m.Called()
	var r0 *model.DefaultValues
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.DefaultValues)
	}
	return r0, ret.Error(1)
}

func (m *mockOccurrenceService) CreateOccurrence(req *model.OccurrenceCreate) (*model.Occurrence, error) {
	// 引数も記録するのだ
	ret := m.Called(req)
	var r0 *model.Occurrence
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.Occurrence)
	}
	return r0, ret.Error(1)
}

// AttachFiles のモックも同様に作る（今回はテストしないので中身は省略）
// func (m *mockOccurrenceService) UploadAttachments(...) ...

// --- ステップ2: テスト関数を書くのだ ---

func TestGetCreatePage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("成功ケース", func(t *testing.T) {
		// --- Arrange (準備) ---
		// モックのインスタンスを作成
		mockService := new(mockOccurrenceService)
		
		// 期待する返り値を準備
		expectedDropdowns := &model.Dropdowns{ /* ... ダミーデータ ... */ }
		expectedDefaults := &model.DefaultValues{ /* ... ダミーデータ ... */ }

		// モックの振る舞いを定義：「このメソッドが呼ばれたら、この値を返す」と設定
		mockService.On("PrepareCreatePage").Return(expectedDropdowns, nil)
		mockService.On("GetDefaultValue").Return(expectedDefaults, nil)
		
		// テスト用のHTTPレコーダーとGinコンテキストを作成
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// ハンドラに本物ではなく、モックのサービスを注入する！
		handler := &occurrenceHandler{service: mockService}

		// --- Act (実行) ---
		handler.GetCreatePage(c)

		// --- Assert (検証) ---
		// ステータスコードが200 OKか？
		assert.Equal(t, http.StatusOK, w.Code)
		
		// レスポンスボディは期待通りか？
		var resBody gin.H
		json.Unmarshal(w.Body.Bytes(), &resBody)
		assert.Equal(t, expectedDropdowns, resBody["dropdown_list"])
		assert.Equal(t, expectedDefaults, resBody["default_value"])

		// モックがちゃんと設定通りに呼ばれたか検証
		mockService.AssertExpectations(t)
	})

	t.Run("サービスでエラーが発生するケース", func(t *testing.T) {
		// --- Arrange (準備) ---
		mockService := new(mockOccurrenceService)
		
		// PrepareCreatePage でエラーが返るように設定
		mockService.On("PrepareCreatePage").Return(nil, errors.New("DB connection error"))
		
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		handler := &occurrenceHandler{service: mockService}

		// --- Act (実行) ---
		handler.GetCreatePage(c)

		// --- Assert (検証) ---
		// ステータスコードが 500 Internal Server Error か？
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		
		// GetDefaultValue は呼ばれていないはずなのだ
		mockService.AssertNotCalled(t, "GetDefaultValue")
		mockService.AssertExpectations(t)
	})
}

// CreateOccurrence のテストも同様に書けるのだ！
// t.Run() を使うと、1つのテスト関数の中に複数のテストケースを書けて便利なのだ。
