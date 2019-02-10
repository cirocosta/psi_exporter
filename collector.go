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

func (c *PsiCollector) Describe(ch chan<- *prometheus.Desc) {
	return
}

func (c *PsiCollector) Collect(ch chan<- prometheus.Metric) {
	return
}
