package logging

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace"
)

func TestBuildEntryIncludesLowercaseTraceAndSpanIDs(t *testing.T) {
	traceID, err := trace.TraceIDFromHex("4bf92f3577b34da6a3ce929d0e0e4736")
	require.NoError(t, err)
	spanID, err := trace.SpanIDFromHex("00f067aa0ba902b7")
	require.NoError(t, err)

	ctx := trace.ContextWithSpanContext(context.Background(), trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    traceID,
		SpanID:     spanID,
		TraceFlags: trace.FlagsSampled,
	}))
	ctx = WithCorrelationID(ctx, "correlation-123")

	value := buildEntry(ctx, LevelInfo, "request complete")
	encoded, err := json.Marshal(value)
	require.NoError(t, err)

	var result map[string]string
	require.NoError(t, json.Unmarshal(encoded, &result))
	require.Equal(t, traceID.String(), result["trace_id"])
	require.Equal(t, spanID.String(), result["span_id"])
	require.Equal(t, "correlation-123", result["correlation_id"])
	require.Equal(t, "request complete", result["message"])
}
