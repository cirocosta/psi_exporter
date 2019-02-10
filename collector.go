package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type PsiCollector struct{}

var (
	_ prometheus.Collector = (*PsiCollector)(nil)
)

func init() {
	prometheus.Register(&PsiCollector{})
}

// some avg10=0.00 avg60=0.00 avg300=0.00 total=3435136
func (c *PsiCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- prometheus.NewDesc(
		"psi_cpu_some_avg10",
		"todo",
		nil,
		nil)

	ch <- prometheus.NewDesc(
		"psi_cpu_some_avg60",
		"todo",
		nil,
		nil)

	ch <- prometheus.NewDesc(
		"psi_cpu_some_avg300",
		"todo",
		nil,
		nil)

	ch <- prometheus.NewDesc(
		"psi_cpu_some_total",
		"todo",
		nil,
		nil)
	return
}

func (c *PsiCollector) Collect(ch chan<- prometheus.Metric) {
	return
}
