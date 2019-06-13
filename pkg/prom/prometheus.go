package prom

import (
	"github.com/yitume/shop-api/pkg/bootstrap"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// HTTPServer for http server
	HTTPServerTimer = NewPromTimer("yitu_http_server", []string{"app", "env", "user", "method"})

	HTTPServerCounter = NewPromCounter("yitu_http_server_code", []string{"app", "env", "user", "method", "code"})

	AppBuildInfo = NewPromCounter("yitu_app_build_info", []string{"app", "env", "version"})
)

// Prom struct info
type PromTimer struct {
	histogram *prometheus.HistogramVec
	summary   *prometheus.SummaryVec
}

type PromCounter struct {
	counter *prometheus.GaugeVec
}

// New creates a Prom instance.
func NewPromTimer(name string, labels []string) *PromTimer {
	obj := &PromTimer{
		histogram: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: name,
				Help: name,
			}, labels),
	}
	prometheus.MustRegister(obj.histogram)
	return obj
}

func NewPromCounter(name string, labels []string) *PromCounter {
	obj := &PromCounter{
		counter: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: name,
				Help: name,
			}, labels),
	}
	prometheus.MustRegister(obj.counter)
	return obj

}

// Timing log timing information (in milliseconds) without sampling
func (p *PromTimer) Timing(name string, time int64, extra ...string) {
	label := append([]string{bootstrap.Conf.Build.App, bootstrap.Conf.Build.Env, name}, extra...)
	p.histogram.WithLabelValues(label...).Observe(float64(time))
}

// Incr increments one stat counter without sampling
func (p *PromCounter) Incr(name string, extra ...string) {
	label := append([]string{bootstrap.Conf.Build.App, bootstrap.Conf.Build.Env, name}, extra...)
	p.counter.WithLabelValues(label...).Inc()
}

// Decr decrements one stat counter without sampling
func (p *PromCounter) Decr(name string, extra ...string) {
	label := append([]string{bootstrap.Conf.Build.App, bootstrap.Conf.Build.Env, name}, extra...)
	p.counter.WithLabelValues(label...).Dec()
}

// Add add count    v must > 0
func (p *PromCounter) Add(name string, v int64, extra ...string) {
	label := append([]string{bootstrap.Conf.Build.App, bootstrap.Conf.Build.Env, name}, extra...)
	p.counter.WithLabelValues(label...).Add(float64(v))
}

func (p *PromCounter) Set(name string, extra ...string) {
	label := append([]string{bootstrap.Conf.Build.App, bootstrap.Conf.Build.Env, name}, extra...)
	p.counter.WithLabelValues(label...).Set(1)
}
