package splitd

type Config struct {
	SourceURL    string
	SourceAuth   string
	SourceMethod string
	DestURL      string
	DestAuth     string
	DestMethod   string
}

func NewConfig() *Config {
	return &Config{}
}
