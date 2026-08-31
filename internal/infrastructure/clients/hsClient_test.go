package clients

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestExecuteGetRequestPropagatesTraceContext(t *testing.T) {
	recorder := tracetest.NewSpanRecorder()
	provider := trace.NewTracerProvider(trace.WithSampler(trace.AlwaysSample()), trace.WithSpanProcessor(recorder))
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

	var outboundSpan trace.ReadOnlySpan
	for _, endedSpan := range recorder.Ended() {
		if endedSpan.Name() == "blizzard.api.request" {
			outboundSpan = endedSpan
			break
		}
	}
	require.NotNil(t, outboundSpan)
	require.Equal(t, "blizzard.api.request", outboundSpan.Name())
	require.Equal(t, span.SpanContext().TraceID(), outboundSpan.SpanContext().TraceID())
	require.NotContains(t, fmt.Sprint(outboundSpan.Attributes()), server.URL)
	require.NotContains(t, fmt.Sprint(outboundSpan.Attributes()), "test-token")

	select {
	case received := <-traceparent:
		require.NotEmpty(t, received)
		require.Contains(t, received, span.SpanContext().TraceID().String())
	case <-time.After(time.Second):
		t.Fatal("downstream service did not receive a request")
	}
}
