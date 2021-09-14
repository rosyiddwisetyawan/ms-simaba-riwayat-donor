package router

import (
	"time"

	jdd_router "ms-simaba-riwayat-donor/router/riwayat"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

//InitRouter is used for initiate router
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Headers",
			"Accept-Encoding", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.Use(apmgin.Middleware(router))
	jdd_router.RiwayatRouter(router)
	return router
}
