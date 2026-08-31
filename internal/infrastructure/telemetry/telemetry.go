// Package telemetry configures OpenTelemetry without making telemetry availability
// a dependency of the API's request path.
package telemetry

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

const (
	defaultServiceName = "hscards-api"
	exportTimeout      = 5 * time.Second
)

// Providers owns the process-wide providers installed by Setup.
type Providers struct {
	tracerProvider *trace.TracerProvider
	meterProvider  *metric.MeterProvider
}

// Setup installs OTLP/HTTP trace and metric providers. Export failures happen in
// the SDK's background workers and never make this function fail after providers
// have been constructed.
func Setup(ctx context.Context) (*Providers, error) {
	resource := newResource(ctx)

	traceExporter, err := otlptracehttp.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("create OTLP trace exporter: %w", err)
	}
	metricExporter, err := otlpmetrichttp.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("create OTLP metric exporter: %w", err)
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithResource(resource),
		trace.WithBatcher(traceExporter,
			trace.WithMaxQueueSize(2048),
			trace.WithMaxExportBatchSize(512),
			trace.WithBatchTimeout(5*time.Second),
			trace.WithExportTimeout(exportTimeout),
		),
	)
	meterProvider := metric.NewMeterProvider(
		metric.WithResource(resource),
		metric.WithReader(metric.NewPeriodicReader(metricExporter,
			metric.WithInterval(30*time.Second),
			metric.WithTimeout(exportTimeout),
		)),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetMeterProvider(meterProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second))

	return &Providers{tracerProvider: tracerProvider, meterProvider: meterProvider}, nil
}

func newResource(ctx context.Context) *resource.Resource {
	serviceName := strings.TrimSpace(os.Getenv("OTEL_SERVICE_NAME"))
	if serviceName == "" {
		serviceName = defaultServiceName
	}

	result, _ := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithTelemetrySDK(),
		resource.WithProcess(),
		resource.WithOS(),
		resource.WithHost(),
		resource.WithAttributes(attribute.String("service.name", serviceName)),
	)
	return result
}

// Shutdown flushes providers within the caller's deadline. Errors are returned
// for lifecycle logging only; callers must not use them to fail in-flight work.
func (p *Providers) Shutdown(ctx context.Context) error {
	if p == nil {
		return nil
	}

	var errs []error
	if p.meterProvider != nil {
		errs = append(errs, p.meterProvider.Shutdown(ctx))
	}
	if p.tracerProvider != nil {
		errs = append(errs, p.tracerProvider.Shutdown(ctx))
	}
	return errors.Join(errs...)
}
