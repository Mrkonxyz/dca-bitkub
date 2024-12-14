package bitkub

import "encoding/json"

type BTC struct {
	Btc MarketData `json:"THB_BTC"`
}
type MarketData struct {
	ID            int     `json:"id"`
	Last          float64 `json:"last"`
	LowestAsk     float64 `json:"lowestAsk"`
	HighestBid    float64 `json:"highestBid"`
	PercentChange float64 `json:"percentChange"`
	BaseVolume    float64 `json:"baseVolume"`
	QuoteVolume   float64 `json:"quoteVolume"`
	IsFrozen      int     `json:"isFrozen"`
	High24Hr      float64 `json:"high24hr"`
	Low24Hr       float64 `json:"low24hr"`
}

func (bk *Bitkub) BtcPrice() (response *BTC, err error) {
	res, err := bk.ApiService.Get("https://api.bitkub.com/api/market/ticker?sym=THB_BTC")
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res, &response); err != nil {
		return nil, err
	}

	return response, nil
}
