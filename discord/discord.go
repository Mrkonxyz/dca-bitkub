package discord

import (
	"Mrkonxyz/github.com/api"
	"bytes"
	"encoding/json"
)

type Discord struct {
	ApiService *api.ApiService
}

func NewDiscordService(ApiService *api.ApiService) *Discord {
	return &Discord{ApiService: ApiService}
}

func (d *Discord) SentMessage(message string) (response []byte, err error) {
	body, err := json.Marshal(&DiscordBody{
		Content: message,
	})
	if err != nil {
		return nil, err
	}
	return d.ApiService.Post(d.ApiService.Cfg.DiscordHook, bytes.NewBuffer(body))
}
