package service

import (
	"Mrkonxyz/github.com/model"
	"Mrkonxyz/github.com/utils"

	"bytes"
	"encoding/json"
)

type DiscordService struct {
	ApiService *utils.ApiService
}

func NewDiscordService(ApiService *utils.ApiService) *DiscordService {
	return &DiscordService{ApiService: ApiService}
}

func (d *DiscordService) SentMessage(message string) (response []byte, err error) {
	body, err := json.Marshal(&model.DiscordBody{
		Content: message,
	})
	if err != nil {
		return nil, err
	}
	return d.ApiService.Post(d.ApiService.Cfg.DiscordHook, bytes.NewBuffer(body))
}
