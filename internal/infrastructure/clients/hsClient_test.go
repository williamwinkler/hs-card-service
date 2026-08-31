package clients

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
)

func TestExecuteGetRequestPropagatesTraceContext(t *testing.T) {
	provider := trace.NewTracerProvider(trace.WithSampler(trace.AlwaysSample()))
	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	t.Cleanup(func() {
		_ = provider.Shutdown(context.Background())
	})

	traceparent := make(chan string, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceparent <- r.Header.Get("traceparent")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &HsClient{
		client: newHTTPClient(),
		token:  token{accessToken: "test-token", expireDate: time.Now().Add(time.Hour)},
	}
	ctx, span := otel.Tracer("test").Start(context.Background(), "parent")
	defer span.End()

	response, err := client.executeGetRequest(ctx, server.URL)
	require.NoError(t, err)
	response.Body.Close()

	select {
	case received := <-traceparent:
		require.NotEmpty(t, received)
		require.Contains(t, received, span.SpanContext().TraceID().String())
	case <-time.After(time.Second):
		t.Fatal("downstream service did not receive a request")
	}
}
