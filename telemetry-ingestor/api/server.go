package api

import (
	"log"
	"os"
	"telemetry-ingestor/api/gateway"
	"telemetry-ingestor/api/stats"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartAPIServer() {

	r := gin.New()

	registerRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("HTTP server running on port", port)

	r.Run(":" + port)
}

func registerRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/v1")
	v1.POST("/telemetry", gateway.TelemetryIngress)
	v1.POST("/telemetrybatch", gateway.TelemetryBatchIngress)
	v1.GET("/devices", stats.GetDeviceList)
	v1.GET("/devices/:imei", stats.DeviceWiseStats)

}
