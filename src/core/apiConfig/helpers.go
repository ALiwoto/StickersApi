package apiConfig

import (
	"os"

	ws "github.com/AnimeKaizoku/ssg/ssg"
)

func LoadConfig() error {
	os.Setenv("DATABASE_URL", "ABCSF")
	return LoadConfigFromFile("config.ini:virtual")
}

func LoadConfigFromFile(fileName string) error {
	if ConfigValue != nil {
		return nil
	}

	var config = &ApiConfig{}

	err := ws.ParseConfig(config, fileName)
	if err != nil {
		return err
	}

	ConfigValue = config

	return nil
}

func GetPort() string {
	if ConfigValue != nil {
		return ConfigValue.Port
	}
	return "8080" // default port is set to 8080
}

func IsTokenValid(value string) bool {
	if ConfigValue == nil {
		return false
	}

	return ConfigValue.MasterToken == value
}

func IsDebug() bool {
	if ConfigValue == nil {
		return false
	}
	return ConfigValue.Debug
}

func UseSqlite() bool {
	if ConfigValue == nil {
		return false
	}
	return ConfigValue.UseSqlite
}

func GetDatabaseUrl() string {
	if ConfigValue == nil {
		return ""
	}
	return ConfigValue.DbUrl
}
