package logging

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Level string

const (
	LevelInfo  Level = "info"
	LevelDebug Level = "debug"
)

var currentLevel = LevelInfo

const CorrelationIDHeader = "X-Correlation-Id"

type contextKey string

const correlationIDContextKey contextKey = "correlation_id"

func ConfigureFromEnv() {
	level := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL")))
	switch Level(level) {
	case LevelDebug:
		currentLevel = LevelDebug
	default:
		currentLevel = LevelInfo
	}
	log.Printf("Log level set to: %s", currentLevel)
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
	log.Printf(prefix(ctx)+format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	log.Printf(prefix(ctx)+format, args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	if currentLevel == LevelDebug {
		log.Printf("[DEBUG] "+prefix(ctx)+format, args...)
	}
}

func prefix(ctx context.Context) string {
	correlationID := CorrelationIDFromContext(ctx)
	if correlationID == "" {
		return ""
	}
	return "[cid=" + correlationID + "] "
}
