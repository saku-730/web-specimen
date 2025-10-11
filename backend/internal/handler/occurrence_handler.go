// backend/internal/handler/occurence_handler.go
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saku-730/web-specimen/backend/internal/model"
	"github.com/saku-730/web-specimen/backend/internal/service"
)

type OccurrenceHandler interface {
	GetCreatePage(c *gin.Context)
	CreateOccurrence(c *gin.Context)
	AttachFiles(c *gin.Context)
}

type occurrenceHandler struct {
	service service.OccurrenceService
}

func NewOccurrenceHandler(occS service.OccurrenceService) OccurrenceHandler {
	return &occurrenceHandler{occService: occS}
}

	//register occurrence route
func (h *occurrenceHandler) RegisterOccurrenceRoutes(rg *gin.RouterGroup) {
	// /create
	rg.GET("/create", h.GetCreatePage)
	rg.POST("/create", h.CreateOccurrence)
	rg.POST("/create/:occurrence_id/attachments")

	// /occurrences
//	rg.GET("/occurrences/:occurrence_id", h.GetOccurrence)
//	rg.PUT("/occurrences/:occurrence_id", h.UpdateOccurrence)
//	rg.DELETE("/occurrences/:occurrence_id", h.DeleteOccurrence)

	// 添付ファイル関連
	// OpenAPIの /create/{id}/attachments は、/create のレスポンスを元にフロントがリクエストを投げる想定なのだ
//	rg.POST("/create/:occurrence_id/attachments", h.AttachFiles) 
	
	// 既存のデータに添付ファイルを追加・削除するルートも定義しておくのだ
	// rg.POST("/occurrences/:occurrence_id/attachments", h.AddAttachment)
	// rg.DELETE("/occurrences/:occurrence_id/attachments/:attachment_id", h.DeleteAttachment)
}



func (h *occurrenceHandler) GetCreatePage(c *gin.Context) {
	dropdowns, err := h.service.PrepareCreatePage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"prepare create page service error": err.Error()})
		return
	}

	defaultValues, err := h.service.GetDefaultValue()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"get default value service error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"dropdown_list": dropdowns,
		"default_value": defaultValues,
	})
}


func (h *occurrenceHandler) CreateOccurrence(c *gin.Context) {
	var req model.OccurrenceCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"JSON bind error": err.Error()})
		return
	}

	created, err := h.service.CreateOccurrence(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"create occurrence service error": err.Error()})
		return
	}
	
	c.Header("Location", "/occurrence/"+strconv.Itoa(created.OccurrenceID))
	c.JSON(http.StatusCreated, created)
}


func (h *occurrenceHandler) AttachFiles(c *gin.Context) {
	idStr := c.Param("occurrence_id")
	occurrenceID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid occurrence ID"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get multipart form: " + err.Error()})
		return
	}
	files := form.File["files"] // "files"front end (next.js) input name

	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No files uploaded"})
		return
	}

	attachmentInfos, err := h.service.UploadAttachments(uint(occurrenceID), files)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attachmentInfos)
}

