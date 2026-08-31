package observability

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestFailUsesOnlyTheStableReason(t *testing.T) {
	recorder := tracetest.NewSpanRecorder()
	provider := trace.NewTracerProvider(trace.WithSpanProcessor(recorder))
	otel.SetTracerProvider(provider)
	t.Cleanup(func() { _ = provider.Shutdown(context.Background()) })

	_, span := Start(context.Background(), "test.operation")
	Fail(span, "database_failed")
	span.End()

	ended := recorder.Ended()
	require.Len(t, ended, 1)
	require.Equal(t, codes.Error, ended[0].Status().Code)
	require.Equal(t, "database_failed", ended[0].Status().Description)
	require.Len(t, ended[0].Events(), 1)
	require.Equal(t, "database_failed", ended[0].Events()[0].Name)
}
