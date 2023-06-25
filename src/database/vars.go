package database

import (
	"sync"
	"time"

	"github.com/AnimeKaizoku/ssg/ssg"
	"gorm.io/gorm"
)

var (
	packInfoMap = func() *ssg.SafeEMap[string, StickerPackInfo] {
		m := ssg.NewSafeEMap[string, StickerPackInfo]()
		m.SetExpiration(6 * time.Hour)
		m.SetInterval(5 * time.Hour)

		return m
	}()
	dbMutex = &sync.Mutex{}
)

var (
	modelStickerPackInfo         = &StickerPackInfo{}
	stickerPackInfoValueNotFound = &StickerPackInfo{}
)

var (
	SESSION *gorm.DB
)
