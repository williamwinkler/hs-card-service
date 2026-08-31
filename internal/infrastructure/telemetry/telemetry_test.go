package telemetry

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTelemetryUnavailableDoesNotBlockRequestHandling(t *testing.T) {
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://127.0.0.1:1")
	t.Setenv("OTEL_EXPORTER_OTLP_PROTOCOL", "http/protobuf")

	providers, err := Setup(context.Background())
	require.NoError(t, err)
	t.Cleanup(func() {
		// A failed flush is expected because the test endpoint is unreachable.
		_ = providers.Shutdown(context.Background())
	})

	middleware, err := NewHTTPMiddleware()
	require.NoError(t, err)
	handler := middleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))

	request := httptest.NewRequest(http.MethodGet, "/api/v1/cards?name=not-a-metric-label", nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	require.Equal(t, http.StatusNoContent, response.Code)
}

func TestStableRouteDoesNotExposeUnknownPaths(t *testing.T) {
	require.Equal(t, "/api/v1/cards", stableRoute("/api/v1/cards"))
	require.Equal(t, "unmatched", stableRoute("/api/v1/cards/customer@example.com"))
}
