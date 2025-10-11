package main

import (
	"log"

	"github.com/saku-730/web-specimen/backend/config"
	"github.com/saku-730/web-specimen/backend/internal/handler"
	"github.com/saku-730/web-specimen/backend/internal/infrastructure"
	"github.com/saku-730/web-specimen/backend/internal/repository"
	"github.com/saku-730/web-specimen/backend/internal/service"
	"github.com/saku-730/web-specimen/backend/internal/router"
)

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

	// Service層を初期化
	occService := service.NewOccurrenceService(occRepo)

	// Handler層を初期化
	occHandler := handler.NewOccurrenceHandler(occService)

	//setup router

	appRouter := router.SetupRouter(occHandler)

	// start server
	log.Printf("Start server port:%s", cfg.ServerPort)
	if err := appRouter.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed start server: %v", err)
	}
}
