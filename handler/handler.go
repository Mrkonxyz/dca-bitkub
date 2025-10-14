package handler

import "Mrkonxyz/github.com/service"

type Handler struct {
	Dca     *DcaHandler
	History *HistoryHandler
}

func NewHandler(dcaService *service.DcaService, bkService *service.BitKubService, dsService *service.DiscordService, historyService *service.HistoryService) *Handler {
	return &Handler{
		Dca:     NewDcaHandler(dcaService, bkService, dsService, historyService),
		History: NewHistoryHandler(historyService),
	}
}
