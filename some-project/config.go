package someProject

type Config struct {
	MetricsAddress string
}

func (cfg *Config) MetricServerAddress() string {
	return cfg.MetricsAddress
}
