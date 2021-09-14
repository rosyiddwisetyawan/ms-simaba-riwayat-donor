package main

import (
	"ms-simaba-riwayat-donor/database"
	"ms-simaba-riwayat-donor/router"
	"os"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	database.ConnectPostgres()
	router := router.InitRouter()
	router.Run(":" + os.Getenv("PORT"))
}
