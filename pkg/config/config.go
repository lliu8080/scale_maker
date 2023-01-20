package config

var (
	AppConfig Config
)

type Config struct {
	OpenWeatherURL       string `json:"openWeatherURL" yaml:"openWeatherURL"`
	OpenWeatherOneAPIURI string `json:"openWeatherOneAPIURI" yaml:"openWeatherOneAPIURI"`
	OpenWeatherAppID     string `json:"openWeatherAppID" yaml:"openWeatherAppID"`
	Port                 string `json:"port" yaml:"port"`
	Prod                 bool   `json:"prod" yaml:"prod"`
}

func NewConfig() {
	AppConfig = Config{
		OpenWeatherURL:       "api.openweathermap.org",
		OpenWeatherOneAPIURI: "/data/2.5/onecall",
		OpenWeatherAppID:     "19477321e9f6c4d85f068c80dd5fa784",
		Port:                 ":3000",
		Prod:                 false,
	}
}
