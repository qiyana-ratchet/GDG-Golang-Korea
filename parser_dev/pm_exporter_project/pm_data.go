package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

var pmDesc = prometheus.NewDesc(
	prometheus.BuildFQName(namespace, "pm", "data"),
	"THIS IS SAMPLE METRICS",
	[]string{
		"sysname",
		"release",
		"version",
		"machine",
		"nodename",
		"domainname",
	},
	nil,
)
type pmCollector struct {
	logger   log.Logger
}

type pmdata struct {
	SysName    string
	Release    string
	Version    string
	Machine    string
	NodeName   string
	DomainName string
}

func init() {
	registerCollector("pm", defaultEnabled, newPmCollector)
}

// NewUnameCollector returns new unameCollector.
func newPmCollector(logger log.Logger) (Collector, error) {
	return &pmCollector{logger}, nil
}

func (c *pmCollector) Update(ch chan<- prometheus.Metric) error {
	pmData, err := getPmData()
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(pmDesc, prometheus.GaugeValue, 1,
		pmData.SysName,
		pmData.Release,
		pmData.Version,
		pmData.Machine,
		pmData.NodeName,
		pmData.DomainName,
	)

	return nil
}
