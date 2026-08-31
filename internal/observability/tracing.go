// Package observability contains dependency-neutral tracing helpers shared by
// application, endpoint, and infrastructure code.
package observability

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const instrumentationName = "github.com/williamwinkler/hs-card-service"

// Start creates a span with caller-supplied, low-cardinality attributes.
func Start(ctx context.Context, name string, attributes ...attribute.KeyValue) (context.Context, trace.Span) {
	return otel.Tracer(instrumentationName).Start(ctx, name, trace.WithAttributes(attributes...))
}

// Fail marks a span as failed without recording potentially sensitive error
// text, query values, credentials, or payloads.
func Fail(span trace.Span, reason string) {
	span.SetStatus(codes.Error, reason)
	span.AddEvent(reason)
}
