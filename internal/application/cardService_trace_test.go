package application_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/williamwinkler/hs-card-service/internal/application"
	"github.com/williamwinkler/hs-card-service/internal/domain"
	"github.com/williamwinkler/hs-card-service/internal/tests/mocks"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestGetCardsCreatesChildSpan(t *testing.T) {
	recorder := tracetest.NewSpanRecorder()
	provider := trace.NewTracerProvider(trace.WithSpanProcessor(recorder))
	otel.SetTracerProvider(provider)
	t.Cleanup(func() { _ = provider.Shutdown(context.Background()) })

	repository := mocks.NewCardRepository()
	repository.Cards[1] = domain.Card{ID: 1, Name: "test"}
	service := application.NewCardService(nil, &repository, nil)

	ctx, requestSpan := otel.Tracer("test").Start(context.Background(), "request")
	_, _, err := service.GetCards(ctx, domain.CardFilter{}, 1, 10)
	requestSpan.End()
	require.NoError(t, err)

	var operationSpan trace.ReadOnlySpan
	for _, span := range recorder.Ended() {
		if span.Name() == "cards.get" {
			operationSpan = span
			break
		}
	}
	require.NotNil(t, operationSpan)
	require.Equal(t, "cards.get", operationSpan.Name())
	require.Equal(t, requestSpan.SpanContext().TraceID(), operationSpan.SpanContext().TraceID())
}
