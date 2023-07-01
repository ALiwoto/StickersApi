package tests_test

import (
	"testing"

	"github.com/AnimeKaizoku/StickersApi/src/database"
)

func TestValidation(t *testing.T) {
	if !database.IsPackIdAcceptable("K340828c0e6327567c457_by_StickerStealRobot") {
		t.Error("Pack ID should be acceptable")
		return
	}
}
