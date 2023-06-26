package database

import (
	"github.com/AnimeKaizoku/StickersApi/src/core/apiConfig"
	"github.com/AnimeKaizoku/StickersApi/src/core/utils/logging"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func StartDatabase() error {
	var err error
	var db *gorm.DB
	var loggerValue logger.Interface

	if apiConfig.IsDebug() {
		loggerValue = logger.Default.LogMode(logger.Info)
	} else {
		loggerValue = logger.Default.LogMode(logger.Silent)
	}

	conf := &gorm.Config{
		Logger: loggerValue,
	}

	if apiConfig.UseSqlite() {
		db, err = gorm.Open(sqlite.Open("stickers.db"), conf)
	} else {
		db, err = gorm.Open(postgres.Open(apiConfig.GetDatabaseUrl()), conf)
	}

	if err != nil {
		return err
	}

	SESSION = db
	logging.Info("Database connected")

	// Create tables if they don't exist
	err = SESSION.AutoMigrate(
		modelStickerPackInfo,
	)
	if err != nil {
		return err
	}

	logging.Info("Auto-migrated database schema")
	return nil
}

func GetPackInfo(packId string) *StickerPackInfo {
	if !IsPackIdAcceptable(packId) {
		return nil
	}

	info := packInfoMap.Get(packId)
	if info == stickerPackInfoValueNotFound {
		return nil
	} else if info != nil {
		return info
	}

	info = &StickerPackInfo{}
	err := SESSION.Model(modelStickerPackInfo).Where(
		"pack_id = ?", packId,
	).First(info).Error
	if err == gorm.ErrRecordNotFound {
		packInfoMap.Add(packId, stickerPackInfoValueNotFound)
		return nil
	} else if err != nil {
		logging.Error("Failed to get pack info from database: ", err.Error())
		return nil
	}

	if info.PackId != packId {
		packInfoMap.Add(packId, stickerPackInfoValueNotFound)
		return nil
	}

	packInfoMap.Add(packId, info)
	return info
}

// SearchPackInfo searches for sticker pack info in the database.
func SearchPackInfo(req *SearchPackRequest) ([]*StickerPackInfo, error) {
	if req == nil || !IsPackTitleAcceptable(req.PackTitle) {
		return nil, ErrPackTitleInvalid
	}

	var results []*StickerPackInfo

	whereStr, valuesArray := req.BuildWhereQuery()

	err := SESSION.Model(modelStickerPackInfo).Where(
		whereStr, valuesArray...,
	).Scan(&results).Error

	return results, err
}

func AddPack(packId, packTitle string) (*StickerPackInfo, error) {
	if !IsPackIdAcceptable(packId) || !IsPackTitleAcceptable(packTitle) {
		return nil, ErrPackIdInvalid
	}

	dbMutex.Lock()
	defer dbMutex.Unlock()

	pack := GetPackInfo(packId)
	if pack != nil {
		return pack, nil
	}

	info := &StickerPackInfo{
		PackId:    packId,
		PackTitle: packTitle,
	}

	tx := SESSION.Begin()
	err := tx.Save(info).Error
	tx.Commit()
	if err != nil {
		return nil, err
	}

	packInfoMap.Add(packId, info)
	return info, nil
}

// IsPackIdAcceptable returns true if pack id is acceptable.
// pack id must only contain english characters and '_'.
func IsPackIdAcceptable(packId string) bool {
	if len(packId) < 5 || len(packId) > 32 {
		return false
	}

	for _, c := range packId {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_' {
			continue
		}
		return false
	}
	return true
}

func IsPackTitleAcceptable(packTitle string) bool {
	return len(packTitle) >= 3 && len(packTitle) <= 64
}
