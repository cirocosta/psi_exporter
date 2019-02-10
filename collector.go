package main

import (
	"os"

	"github.com/cirocosta/psi_exporter/parser"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

type PsiCollector struct{}

var (
	_          prometheus.Collector = (*PsiCollector)(nil)
	subsystems                      = []string{"cpu", "io", "memory"}
	metrics                         = []string{"avg10", "avg60", "avg300", "total"}
)

func init() {
	prometheus.Register(&PsiCollector{})
}

func (c *PsiCollector) Describe(ch chan<- *prometheus.Desc) {
	return
}

func collectRowsFromFilepath(filepath string) (rows []*parser.Row, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to open file %s for parsing", filepath)
		return
	}
	defer file.Close()

	rows, err = parser.Parse(file)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to parse memory stats from %s",
			file.Name())
		return
	}

	return
}

func psiStatDescription(subsystem, metric string) *prometheus.Desc {
	return prometheus.NewDesc(
		"psi_"+subsystem+"_"+metric,
		"Values for "+metric+" under /proc/pressure/"+subsystem,
		[]string{"type"},
		nil)
}

func (c *PsiCollector) Collect(ch chan<- prometheus.Metric) {
	var structVals [4]*float64

	for _, subsystem := range subsystems {
		rows, err := collectRowsFromFilepath("/proc/pressure/" + subsystem)
		if err != nil {
			panic(err)
		}

		for _, row := range rows {
			structVals = [4]*float64{&row.Avg10, &row.Avg60, &row.Avg300, &row.Total}

			for idx, metric := range metrics {
				ch <- prometheus.MustNewConstMetric(
					psiStatDescription(subsystem, metric),
					prometheus.UntypedValue,
					*structVals[idx],
					row.Type,
				)
			}
		}
	}

	return
}
