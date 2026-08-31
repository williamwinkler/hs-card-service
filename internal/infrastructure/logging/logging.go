package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type Level string

const (
	LevelInfo  Level = "info"
	LevelDebug Level = "debug"
	LevelError Level = "error"
)

var (
	currentLevel           = LevelInfo
	output       io.Writer = os.Stdout
	outputMu     sync.Mutex
)

const CorrelationIDHeader = "X-Correlation-Id"

type contextKey string

const correlationIDContextKey contextKey = "correlation_id"

type entry struct {
	Timestamp     string `json:"timestamp"`
	Level         Level  `json:"level"`
	Message       string `json:"message"`
	CorrelationID string `json:"correlation_id,omitempty"`
	TraceID       string `json:"trace_id,omitempty"`
	SpanID        string `json:"span_id,omitempty"`
}

func ConfigureFromEnv() {
	level := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL")))
	switch Level(level) {
	case LevelDebug:
		currentLevel = LevelDebug
	default:
		currentLevel = LevelInfo
	}

	log.SetFlags(0)
	log.SetOutput(jsonWriter{})
	Infof(context.Background(), "log level set to: %s", currentLevel)
}

func CorrelationIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := strings.TrimSpace(r.Header.Get(CorrelationIDHeader))
		if correlationID == "" {
			correlationID = uuid.NewString()
		}

		ctx := WithCorrelationID(r.Context(), correlationID)
		r = r.WithContext(ctx)
		w.Header().Set(CorrelationIDHeader, correlationID)

		next.ServeHTTP(w, r)
	})
}

func WithCorrelationID(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, correlationIDContextKey, correlationID)
}

func CorrelationIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	correlationID, _ := ctx.Value(correlationIDContextKey).(string)
	return correlationID
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	emit(ctx, LevelInfo, format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	emit(ctx, LevelError, format, args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	if currentLevel == LevelDebug {
		emit(ctx, LevelDebug, format, args...)
	}
}

func emit(ctx context.Context, level Level, format string, args ...interface{}) {
	writeEntry(buildEntry(ctx, level, fmt.Sprintf(format, args...)))
}

func buildEntry(ctx context.Context, level Level, message string) entry {
	result := entry{
		Timestamp:     time.Now().UTC().Format(time.RFC3339Nano),
		Level:         level,
		Message:       message,
		CorrelationID: CorrelationIDFromContext(ctx),
	}
	if ctx == nil {
		return result
	}
	spanContext := trace.SpanContextFromContext(ctx)
	if spanContext.IsValid() {
		result.TraceID = spanContext.TraceID().String()
		result.SpanID = spanContext.SpanID().String()
	}
	return result
}

func writeEntry(entry entry) {
	encoded, err := json.Marshal(entry)
	if err != nil {
		return
	}
	outputMu.Lock()
	defer outputMu.Unlock()
	_, _ = output.Write(append(encoded, '\n'))
}

// jsonWriter converts legacy standard-library log calls into structured stdout
// entries. Request-aware code should call Infof/Errorf/Debugf to add trace IDs.
type jsonWriter struct{}

func (jsonWriter) Write(message []byte) (int, error) {
	writeEntry(buildEntry(context.Background(), LevelInfo, strings.TrimSpace(string(message))))
	return len(message), nil
}
