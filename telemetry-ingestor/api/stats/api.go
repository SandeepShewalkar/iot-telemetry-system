package stats

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"telemetry-ingestor/internal/storage"

	"github.com/gin-gonic/gin"
)

// GetDeviceList godoc
// @Summary      Get device list registered
// @Description  Get list of devices/IMEI
// @Produce      json
// @Router       /v1/devices [get]
func GetDeviceList(c *gin.Context) {
	ctx := context.Background()
	rdb := storage.GetRedisClient()
	keys, err := rdb.Keys(ctx, "device:*").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Remove "device:" prefix
	var imeis []string
	for _, key := range keys {
		imeis = append(imeis, strings.TrimPrefix(key, "device:"))
	}

	c.JSON(http.StatusOK, gin.H{"devices": imeis})
}

// DeviceWiseStats godoc
// @Summary      Get device statistics
// @Description  Get telemetry statistics for a specific device by IMEI
// @Param        imei   path      string  true  "Device IMEI"
// @Produce      json
// @Router       /v1/devices/{imei} [get]
func DeviceWiseStats(c *gin.Context) {
	rdb := storage.GetRedisClient()
	imei := c.Param("imei")
	key := "device:" + imei
	ctx := context.Background()

	keyType, err := rdb.Type(ctx, key).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result interface{}
	switch keyType {
	case "string":
		result, err = rdb.Get(ctx, key).Result()
	case "hash":
		result, err = rdb.HGetAll(ctx, key).Result()
	default:
		err = fmt.Errorf("unsupported key type: %s", keyType)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"imei": imei, "data": result})
}
