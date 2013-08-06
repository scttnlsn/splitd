package splitd

type Config struct {
	SourceURL string
	DestURL   string
}

func NewConfig() *Config {
	return &Config{}
}
