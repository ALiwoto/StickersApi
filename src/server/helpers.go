package server

import (
	"github.com/AnimeKaizoku/StickersApi/src/core/apiConfig"
	"github.com/AnimeKaizoku/StickersApi/src/core/utils/logging"
	"github.com/gin-gonic/gin"
)

func RunStickerApi() {
	port := apiConfig.GetPort()

	if !apiConfig.IsDebug() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Starts a new Gin instance with no middle-ware
	ServerEngine = gin.New()
	ServerEngine.SetTrustedProxies(nil)
	LoadHandlers()
	// Listen and serve on defined port
	logging.Info("Listening on port ", port)

	err := ServerEngine.Run(":" + port)
	if err != nil {
		logging.Error(err)
	}
}
