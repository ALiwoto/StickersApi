package apiConfig

type ApiConfig struct {
	// MasterToken is the token that is used to add packs.
	MasterToken string `section:"main" key:"master_token"`

	// Port is the http port that we should listen to
	Port string `section:"main" key:"port" default:"8080"`

	// Debug is the debug mode.
	Debug bool `section:"main" key:"debug" default:"false"`

	// DbUrl is the of your postgresql database.
	// if `use_sqlite` is set to `true`, this variable will be ignored.
	DbUrl string `section:"database" key:"url"`

	// UseSqlite
	UseSqlite bool `section:"database" key:"use_sqlite"`
}
