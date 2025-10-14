package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"


	"github.com/saku-730/web-specimen/backend/config"
	"github.com/saku-730/web-specimen/backend/internal/handler"
	"github.com/saku-730/web-specimen/backend/internal/infrastructure"
	"github.com/saku-730/web-specimen/backend/internal/repository"
	"github.com/saku-730/web-specimen/backend/internal/service"
	"github.com/saku-730/web-specimen/backend/internal/router"
	"github.com/saku-730/web-specimen/backend/internal/middleware"
)

func ProfileHandler(c *gin.Context) {
	// ミドルウェアで保存したユーザー情報を取得する
	userID, _ := c.Get("userID")
	userName, _ := c.Get("userName")

	c.JSON(http.StatusOK, gin.H{
		"message":   userName.(string),
		"user_id":   userID,
	})
}

func main() {
	// load config
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed load config: %v", err)
	}

	// connect database
	db, err := database.NewDatabaseConnection(cfg)
	if err != nil {
		log.Fatalf("Failed connect database: %v", err)
	}


	// Repository層を初期化
	occRepo := repository.NewOccurrenceRepository(db)
	userRepo := repository.NewUserRepository(db)
	userDefaultsRepo := repository.NewUserDefaultsRepository(db)
	attachmentRepo := repository.NewAttachmentRepository()
	attachmentGroupRepo := repository.NewAttachmentGroupRepository()
	fileExtensionRepo := repository.NewFileExtensionRepository()

	// Service層を初期化
	authService := service.NewAuthService(userRepo,cfg)
	occService := service.NewOccurrenceService(db,occRepo,userDefaultsRepo,attachmentRepo,attachmentGroupRepo,fileExtensionRepo)

	// Handler層を初期化
	authHandler := handler.NewAuthHandler(authService)
	occHandler := handler.NewOccurrenceHandler(occService)

	// Middlreware
	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret)

	//setup router

	appRouter := router.SetupRouter(
		authHandler,
		occHandler,
		authMiddleware,
	)

	// start server
	log.Printf("Start server port:%s", cfg.ServerPort)
	if err := appRouter.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed start server: %v", err)
	}
}
