package service

import (
	"Mrkonxyz/github.com/model/entity"
	"Mrkonxyz/github.com/repository"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type HistoryService struct {
	syncTopUpRepository *repository.SyncTopUpRepository
	topUpRepository     *repository.TopUpRepository
	bitKubService       *BitKubService
}

func NewHistoryService(repoSync *repository.SyncTopUpRepository, repoTopUp *repository.TopUpRepository, BitKubService *BitKubService) *HistoryService {
	return &HistoryService{repoSync, repoTopUp, BitKubService}
}

func (s *HistoryService) SyncDepositHistory(ctx context.Context) (interface{}, error) {

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	res, err := s.bitKubService.DepositHistory()
	if err != nil {
		return nil, err
	}

	// check last sync time
	var lastSync entity.SyncTopUpOffset
	lastSync, err = s.syncTopUpRepository.FindLastSync(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			lastSync = entity.SyncTopUpOffset{
				LastSyncDate: time.Date(
					2023,         // Year
					time.January, // Month (ใช้ constant time.January แทน 1)
					1,            // Day
					0,            // Hour
					0,            // Minute
					0,            // Second
					0,            // Nanosecond
					time.UTC,     // Location/Timezone
				),
			}
			lastSync, _ = s.syncTopUpRepository.CreateLastSync(ctx, lastSync)
		}
	}

	var histories []entity.TopUpHistory
	for _, item := range res.Result {
		if time.Time(item.Time).Before(lastSync.LastSyncDate) {
			continue
		}
		history := entity.TopUpHistory{
			TxnID:    item.TxnID,
			Amount:   item.Amount,
			Currency: item.Currency,
			Status:   item.Status,
			Time:     time.Time(item.Time),
		}
		histories = append(histories, history)
	}
	if len(histories) == 0 {
		return gin.H{"message": "No new deposit history found"}, nil
	}
	err = s.topUpRepository.SaveAll(ctx, histories)
	if err != nil {
		return nil, err
	}

	lastSync.LastSyncDate = time.Now()
	err = s.syncTopUpRepository.UpdateLastSync(ctx, lastSync)
	return histories, err
}
