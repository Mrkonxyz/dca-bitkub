package handler

import (
	"Mrkonxyz/github.com/model"
	"Mrkonxyz/github.com/service"
	"Mrkonxyz/github.com/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DcaHandler struct {
	service   *service.DcaService
	BkService *service.BitKubService
	DsService *service.DiscordService
}

func NewDcaHandler(service *service.DcaService, bkService *service.BitKubService, dsService *service.DiscordService) *DcaHandler {
	return &DcaHandler{service, bkService, dsService}
}

func (h *DcaHandler) CreateDca(c *gin.Context) {
	var req model.Dca
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateDca(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create dca"})
		return
	}
	c.JSON(http.StatusCreated, req)
}

func (h *DcaHandler) GetDca(c *gin.Context) {
	dcas, err := h.service.GetDca(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch dca"})
		return
	}
	c.JSON(http.StatusOK, dcas)
}

func (h *DcaHandler) Trigger(c *gin.Context) {
	today, err := getDay()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var req model.DcaRequest
	if err := c.BindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	buyRes, err := h.BkService.BuyBitCion(req.Amount)
	if buyRes.Error != 0 || err != nil {
		errorMessage, exists := utils.ErrorMessages[buyRes.Error]
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
		# ที่ราคา %sบาท
		# จำนวนเงิน %sบาท
		# =====================
		`,
		today, utils.FormatMoney(res["THB_BTC"].Last), utils.FormatMoney(req.Amount))

	_, err = h.DsService.SentMessage(message)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "DCA BTC success ✅"})
}

func (h *DcaHandler) GetWallet(c *gin.Context) {
	res, err := h.BkService.GetWallet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, res)
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
