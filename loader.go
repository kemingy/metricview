package metricview

import (
	"fmt"
	"os"
)

type MetricView struct {
	directory string
}

func NewMetricView(directory string) *MetricView {
	info, err := os.Stat(directory)
	if err != nil {
		fmt.Errorf("directory does not exist: %s", err)
	}
}