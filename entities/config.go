package entities

type Config struct {
	AppVersionName string `toml:"AppVersionName"`

	AppUpdateVersionRequired bool `toml:"AppUpdateVersionRequired"`

	Database Database `toml:"Database"`

	Server Server `toml:"Server"`

	Paseto Paseto `toml:"Paseto"`

	LogDir string `toml:"LogDir"`
}

type Server struct {
	Port int `toml:"Port"`

	AllowedOrigins []string `toml:"AllowedOrigins"`

	ReadTimeout int `toml:"ReadTimeout"`

	IdleTimeout int `toml:"IdleTimeout"`

	WriteTimeout int `toml:"WriteTimeout"`
}

type Database struct {
	Host     string `toml:"Host"`
	Port     string `toml:"Port"`
	User     string `toml:"User"`
	Password string `toml:"Password"`
	Database string `toml:"Database"`
}

type Paseto struct {
	PasetoSecurityKey  string `toml:"PasetoSecurityKey"`
	UserPassSaltSecret string `toml:"UserPassSaltSecret"`
}
