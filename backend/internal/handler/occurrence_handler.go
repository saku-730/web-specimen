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
	GetSearchPage(c *gin.Context)
	Search(c *gin.Context)
}

type occurrenceHandler struct {
	service service.OccurrenceService
}

func NewOccurrenceHandler(occS service.OccurrenceService) OccurrenceHandler {
	return &occurrenceHandler{service: occS}
}


func (h *occurrenceHandler) GetCreatePage(c *gin.Context) {
	dropdowns, err := h.service.PrepareCreatePage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"prepare create page service error": err.Error()})
		return
	}

	userIDInterface, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Not fount userID context"})
		return
	}

	// interface{} 型を、元の型 (ここではuintだと仮定するのだ) に変換する
	userID, ok := userIDInterface.(int)
	if !ok {
		// 型の変換に失敗した場合
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid context userID type"})
		return
	}

	defaultValues, err := h.service.GetDefaultValues(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"get default value service error": err.Error()})
		return
	}

	pageData := model.CreatePageData{
		DropdownList: *dropdowns,
		DefaultValue: *defaultValues,
	}

	c.JSON(http.StatusOK, pageData)
}

func(h * occurrenceHandler) CreateOccurrence(c *gin.Context) {
	var req model.OccurrenceCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid request body": err.Error()})
		return
	}

	created, err := h.service.CreateOccurrence(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"create occurrence service error": err.Error()})
		return
	}
	
	c.Header("Location", "/occurrence/"+strconv.Itoa(int(created.OccurrenceID)))
	c.JSON(http.StatusCreated, created)
}


func (h *occurrenceHandler) AttachFiles(c *gin.Context) {
	occurrenceIDStr := c.Param("occurrence_id")
	occurrenceID, err := strconv.ParseUint(occurrenceIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid occurrence_id"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed analyze form data: " + err.Error()})
		return
	}
	files := form.File["upload_files"]

	userID := c.MustGet("userID").(int)
	savedFileNames, err := h.service.AttachFiles(uint(occurrenceID), uint(userID), files)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed upload file: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "success upload file",
		"fileNames": savedFileNames,
	})
}


func (h *occurrenceHandler) GetSearchPage(c *gin.Context) {
	dropdowns, err := h.service.PrepareCreatePage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"prepare create page service error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dropdowns)
}


func (h *occurrenceHandler) Search(c *gin.Context) {
	var query model.SearchQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query paramate: " + err.Error()})
		return
	}

	response, err := h.service.Search(&query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed search process: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
