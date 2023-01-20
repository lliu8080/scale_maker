package config

var (
	//AppConfig - app config object
	AppConfig Config
)

// Config - struct config
type Config struct {
	Port string `json:"port" yaml:"port"`
	Prod bool   `json:"prod" yaml:"prod"`
}

// NewConfig - config
func NewConfig() {
	AppConfig = Config{
		Port: ":3000",
		Prod: false,
	}
}
