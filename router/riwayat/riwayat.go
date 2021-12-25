package riwayat

import (
	jdd_handler "ms-simaba-riwayat-donor/handler/riwayat"
	mdwr "ms-simaba-riwayat-donor/middleware"

	"github.com/gin-gonic/gin"
)

func RiwayatRouter(router *gin.Engine) {
	router.GET("/", jdd_handler.GetMain)
	router.Use(mdwr.JWTMiddleware())
	router.POST("api/simaba/riwayat-donor", jdd_handler.RiwayatHandler)
	router.POST("api/simaba/riwayat-donor/detail", jdd_handler.RiwayatDetailHandler)
	router.POST("api/simaba/riwayat-donor/add", jdd_handler.CreateRiwayatHandler)

}
