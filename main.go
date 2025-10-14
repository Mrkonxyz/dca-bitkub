package main

import (
	"Mrkonxyz/github.com/config"
	"Mrkonxyz/github.com/handler"
	"Mrkonxyz/github.com/repository"
	"Mrkonxyz/github.com/router"
	"Mrkonxyz/github.com/service"
	"Mrkonxyz/github.com/utils"

	"log"
)

func main() {
	cfg := config.LoadConfig(".")

	topUp := repository.NewTopUpRepository(cfg.DB)
	// helper
	apiService := utils.NewApiService(cfg)
	bkService := service.NewBitKubService(apiService, topUp)
	dsService := service.NewDiscordService(apiService)

	// historyService
	syncTopUp := repository.NewSyncTopUpRepository(cfg.DB)

	historyService := service.NewHistoryService(syncTopUp, topUp, bkService)

	// dca
	dcaRepo := repository.NewDcaRepository(cfg.DB)
	dcaService := service.NewDcaService(dcaRepo)

	// handler
	dcaHandler := handler.NewHandler(dcaService, bkService, dsService, historyService)

	r := router.SetupRouter(cfg, dcaHandler)
	log.Printf("Server running at :%s \n", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
