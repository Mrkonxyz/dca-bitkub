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

	// dependencies for dca
	dcaRepo := repository.NewDcaRepository(cfg.DB)
	dcaService := service.NewDcaService(dcaRepo)
	apiService := utils.NewApiService(cfg)
	bkService := service.NewBitKubService(apiService)
	dsService := service.NewDiscordService(apiService)
	dcaHandler := handler.NewDcaHandler(dcaService, bkService, dsService)

	r := router.SetupRouter(cfg, dcaHandler)
	log.Printf("Server running at :%s \n", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
