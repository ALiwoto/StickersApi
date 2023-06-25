package main

import (
	"github.com/AnimeKaizoku/StickersApi/src/core/apiConfig"
	"github.com/AnimeKaizoku/StickersApi/src/core/utils/logging"
	"github.com/AnimeKaizoku/StickersApi/src/database"
	"github.com/AnimeKaizoku/StickersApi/src/server"
)

func main() {
	f := logging.LoadLogger(true)
	if f != nil {
		defer f()
	}

	runApp()
}

func runApp() {
	err := apiConfig.LoadConfig()
	if err != nil {
		logging.Fatal(err)
	}

	database.StartDatabase()
	server.RunStickerApi()
}
