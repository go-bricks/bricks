package monitoring

import (
	"context"

	"github.com/go-bricks/bricks/v2/interfaces/monitor"
)

type bricksReporter struct {
	externalMetrics monitor.BricksMetrics
	cfg             *monitorConfig
	registry        *externalRegistry
}

// NewBricksReporter creates a new bricks monitoring reporter which is a wrapper to support
//   - ContextExtractors
//   - Default Tags, for example: {"version":"v1.0.1", "service":"awesome"}
//
// Meaning, it is possible to also extract tag values from the context, this is useful when the value is set per request/call within the context.Context:
//   - Canary release https://martinfowler.com/bliki/CanaryRelease.html identifier
//   - Authentication Token values, but avoid using high cardinality values such as UserID
func newBricksReporter(cfg *monitorConfig) monitor.Reporter {
	externalMetrics := cfg.reporter.Metrics()
	return &bricksReporter{
		externalMetrics: externalMetrics,
		cfg:             cfg,
		registry:        newRegistry(externalMetrics),
	}
}

func (r *bricksReporter) Connect(ctx context.Context) error {
	return r.cfg.reporter.Connect(ctx)
}

func (r *bricksReporter) Close(ctx context.Context) error {
	return r.cfg.reporter.Close(ctx)
}

func (r *bricksReporter) Metrics() monitor.Metrics {
	return r
}

// Counter creates a counter with possible predefined tags
func (r *bricksReporter) Counter(name string, desc string) monitor.TagsAwareCounter {
	return newMetric(r.registry, r.cfg).WithTags(r.cfg.tags).Counter(name, desc)
}

// Gauge creates a gauge with possible predefined tags
func (r *bricksReporter) Gauge(name string, desc string) monitor.TagsAwareGauge {
	return newMetric(r.registry, r.cfg).WithTags(r.cfg.tags).Gauge(name, desc)
}

// Histogram creates a histogram with possible predefined tags
func (r *bricksReporter) Histogram(name string, desc string, buckets monitor.Buckets) monitor.TagsAwareHistogram {
	return newMetric(r.registry, r.cfg).WithTags(r.cfg.tags).Histogram(name, desc, buckets)
}

// Timer creates a timer with possible predefined tags
func (r *bricksReporter) Timer(name string, desc string) monitor.TagsAwareTimer {
	return newMetric(r.registry, r.cfg).WithTags(r.cfg.tags).Timer(name, desc)
}

// WithTags sets custom tags to be included if possible in every Metric
func (r *bricksReporter) WithTags(tags monitor.Tags) monitor.Metrics {
	return newMetric(r.registry, r.cfg).
		WithTags(r.cfg.tags). // first apply default tags
		WithTags(tags)        // then apply custom ones
}
