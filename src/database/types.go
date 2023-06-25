package database

type StickerPackInfo struct {
	PackId    string `gorm:"primaryKey" json:"pack_id"`
	PackTitle string `json:"pack_title"`
}

type SearchPackRequest struct {
	PackTitle string `json:"pack_title"`
}
