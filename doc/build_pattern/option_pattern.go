package build_pattern

type Config struct {
	host   string
	port   int
	maxNum int
}

type Option func(cfg *Config)

func WithHost(host string) Option {
	return func(cfg *Config) {
		cfg.host = host
	}
}

func WithPort(port int) Option {
	return func(cfg *Config) {
		cfg.port = port
	}
}

func WithMaxNum(maxNum int) Option {
	return func(cfg *Config) {
		cfg.maxNum = maxNum
	}
}

func NewConfig(options ...Option) Config {
	cfg := Config{}
	for _, option := range options {
		option(&cfg)
	}
	return cfg
}
