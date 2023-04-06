package module_library

import (
	"fmt"
)

type ConfigWithMetricServer interface {
	MetricServerAddress() string
}

type ModMetrics[TConfig ConfigWithMetricServer, TState any] struct{}

func (m *ModMetrics[TConfig, TState]) Start(cfg TConfig) error {
	fmt.Println(cfg.MetricServerAddress())
	return nil
}
