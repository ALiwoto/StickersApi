package stickerHandlers

import (
	"github.com/AnimeKaizoku/StickersApi/src/core/apiConfig"
	"github.com/AnimeKaizoku/StickersApi/src/core/utils"
	"github.com/AnimeKaizoku/StickersApi/src/core/utils/logging"
	"github.com/AnimeKaizoku/StickersApi/src/database"
	entry "github.com/AnimeKaizoku/StickersApi/src/entryPoints"
	"github.com/gin-gonic/gin"
)

func GetPackInfoHandler(c *gin.Context) {
	packId := utils.GetParam(c, "pack-id", "packId")
	if len(packId) < 3 {
		entry.SendNoDataError(c, OriginGetPackInfo)
		return
	}

	info := database.GetPackInfo(packId)
	if info == nil {
		entry.SendPackNotFoundError(c, OriginGetPackInfo)
		return
	}

	entry.SendResult(c, info)
}

func SearchPackHandler(c *gin.Context) {
	packTitle := utils.GetParam(c, "title", "pack-title", "packTitle")
	if len(packTitle) < 3 {
		entry.SendNoDataError(c, OriginSearchPack)
	}

	info, err := database.SearchPackInfo(&database.SearchPackRequest{
		PackTitle: packTitle,
	})
	if err != nil {
		logging.Error("Failed to search pack info:", err)
		entry.SendInternalServerError(c, OriginSearchPack)
		return
	} else if info == nil {
		entry.SendPackNotFoundError(c, OriginSearchPack)
		return
	}

	entry.SendResult(c, info)
}

func AddPackHandler(c *gin.Context) {
	packId := utils.GetParam(c, "pack-id", "packId")
	packTitle := utils.GetParam(c, "title", "pack-title", "packTitle")
	token := utils.GetParam(c, "token")
	if !apiConfig.IsTokenValid(token) {
		entry.SendPermissionDenied(c, OriginAddPack)
	}

	if len(packId) < 3 || len(packTitle) < 3 {
		entry.SendNoDataError(c, OriginAddPack)
		return
	}

	info, err := database.AddPack(packId, packTitle)
	if err != nil {
		logging.Error("Failed to add pack info:", err)
		entry.SendInternalServerError(c, OriginAddPack)
		return
	}

	entry.SendResult(c, info)
}
