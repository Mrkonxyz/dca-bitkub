package handler

import (
	"Mrkonxyz/github.com/bitkub"
	"Mrkonxyz/github.com/discord"
	"Mrkonxyz/github.com/util"
	"fmt"
	"log"
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

type DcaRequest struct {
	Amount float64 `json:"amount"`
}

func getDay() (string, error) {
	// Load the Bangkok timezone
	bangkokTimeZone := time.FixedZone("UTC+7", 7*60*60)

	// Get the current time in the Bangkok timezone
	currentTime := time.Now().In(bangkokTimeZone)

	year := currentTime.Year() + 543 // Convert AD to BE
	month := currentTime.Month()
	day := currentTime.Day()

	// Log the date in Thai format
	return fmt.Sprintf("วันที่ %d เดือน %d พ.ศ. %d เวลา %02d:%02d\n", day, month, year, currentTime.Hour(), currentTime.Minute()), nil
}

func (h *Handler) DcaBTC(c *gin.Context) {
	today, err := getDay()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var req DcaRequest
	if err := c.BindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	buyRes, err := h.BkService.BuyBitCion(req.Amount)
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
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		log.Println(errorMessage)
		c.JSON(http.StatusInternalServerError, gin.H{"message": errorMessage})
		return
	}
	res, err := h.BkService.GetPrice("THB_BTC")
	if err != nil {
		log.Println(err.Error())
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
		today, res["THB_BTC"].Last, req.Amount)

	_, err = h.DsService.SentMessage(message)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "DCA BTC success ✅"})
}

func (h *Handler) GetWallet(c *gin.Context) {
	res, err := h.BkService.GetWallet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, res)
}
