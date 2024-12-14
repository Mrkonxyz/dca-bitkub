package handler

import (
	"Mrkonxyz/github.com/bitkub"
	"Mrkonxyz/github.com/discord"
	"Mrkonxyz/github.com/util"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	BkService *bitkub.Bitkub
	DsService *discord.Discord
}

func NewHandler(bkService *bitkub.Bitkub, dsService *discord.Discord) *Handler {
	return &Handler{BkService: bkService, DsService: dsService}
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "service up."})
}
func (h *Handler) GetBTC(c *gin.Context) {
	res, err := h.BkService.BtcPrice()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": res})
}

type BuyBitCionRequest struct {
	Amount float64 `json:"amount"`
}

func getDay() (string, error) {
	// Load the Bangkok timezone
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return "", err
	}

	// Get the current time in the Bangkok timezone
	currentTime := time.Now().In(location)

	// Format the date in Thai language
	// Note: The `January` and `Monday` values will remain in English without additional localization libraries.
	// We'll use numeric formatting here.
	year := currentTime.Year() + 543 // Convert AD to BE
	month := currentTime.Month()
	day := currentTime.Day()

	// Log the date in Thai format
	return fmt.Sprintf("วันที่ %d เดือน %d พ.ศ. %d เวลา %d:%d\n", day, month, year, currentTime.Hour(), currentTime.Minute()), nil
}

func (h *Handler) BuyBitCion(c *gin.Context) {
	today, err := getDay()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var req BuyBitCionRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	buyRes, err := h.BkService.BuyBitCion(req.Amount)
	fmt.Printf("buyRes: %v\n", buyRes)
	if buyRes.Error != 0 || err != nil {
		errorMessage, exists := util.ErrorMessages[buyRes.Error]
		if !exists {
			errorMessage = err.Error()
		}
		message := fmt.Sprintf(
			`
		# =====================
		# %s
		# 🚀 **แจ้งเตือนการซื้อ BTC ไม่สำเร็จ**
		# Error code: %d
		# Reason: %s
		# =====================
		`,
			today, buyRes.Error, errorMessage)

		if _, err = h.DsService.SentMessage(message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": errorMessage})
		return
	}
	res, err := h.BkService.BtcPrice()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	message := fmt.Sprintf(
		`
		# =====================
		# %s
		# 🚀 **แจ้งเตือนการซื้อ BTC**
		# ที่ราคา %.2fบาท
		# จำนวนเงิน %.2fบาท
		# =====================
		`,
		today, res.Btc.Last, req.Amount)

	_, err = h.DsService.SentMessage(message)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "DCA BTC success ✅"})
}
