package riwayat

import (
	jdd_handler "ms-simaba-riwayat-donor/handler/riwayat"
	mdwr "ms-simaba-riwayat-donor/middleware"

	"github.com/gin-gonic/gin"
)

func RiwayatRouter(router *gin.Engine) {
	router.GET("/", jdd_handler.GetMain)
	router.Use(mdwr.JWTMiddleware())
	router.POST("simaba", jdd_handler.RiwayatHandler)
	router.POST("simaba/detail", jdd_handler.RiwayatDetailHandler)
	router.POST("simaba/add", jdd_handler.CreateRiwayatHandler)

}
