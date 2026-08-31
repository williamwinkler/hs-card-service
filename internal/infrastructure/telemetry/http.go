package telemetry

import (
	"net/http"
	"strconv"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const instrumentationName = "github.com/williamwinkler/hs-card-service"

// HTTPMiddleware adds low-cardinality request metrics and server spans. It does
// not record request URLs or query strings.
type HTTPMiddleware struct {
	requestCount metric.Int64Counter
	errorCount   metric.Int64Counter
	duration     metric.Float64Histogram
}

func NewHTTPMiddleware() (*HTTPMiddleware, error) {
	meter := otel.Meter(instrumentationName)
	requestCount, err := meter.Int64Counter("api.http.server.request.count")
	if err != nil {
		return nil, err
	}
	errorCount, err := meter.Int64Counter("api.http.server.error.count")
	if err != nil {
		return nil, err
	}
	duration, err := meter.Float64Histogram("api.http.server.request.duration", metric.WithUnit("s"))
	if err != nil {
		return nil, err
	}
	return &HTTPMiddleware{requestCount: requestCount, errorCount: errorCount, duration: duration}, nil
}

func (m *HTTPMiddleware) Handler(next http.Handler) http.Handler {
	return otelhttp.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startedAt := time.Now()
		writer := &statusWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(writer, r)

		attributes := metricAttributes(r.Method, stableRoute(r.URL.Path), writer.statusCode)
		m.requestCount.Add(r.Context(), 1, metric.WithAttributes(attributes...))
		m.duration.Record(r.Context(), time.Since(startedAt).Seconds(), metric.WithAttributes(attributes...))
		if writer.statusCode >= http.StatusInternalServerError {
			m.errorCount.Add(r.Context(), 1, metric.WithAttributes(attributes...))
		}
	}), "hscards-api.request")
}

type statusWriter struct {
	http.ResponseWriter
	statusCode  int
	wroteHeader bool
}

func (w *statusWriter) WriteHeader(statusCode int) {
	if w.wroteHeader {
		return
	}
	w.statusCode = statusCode
	w.wroteHeader = true
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *statusWriter) Write(body []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(body)
}

func (w *statusWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}

func metricAttributes(method, route string, statusCode int) []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("http.request.method", method),
		attribute.String("http.route", route),
		attribute.String("http.response.status_code", strconv.Itoa(statusCode)),
	}
}

// stableRoute maps every known API endpoint to a static route label. Unknown
// paths are deliberately collapsed to avoid unbounded metric cardinality.
func stableRoute(path string) string {
	switch path {
	case "/api/v1/cards":
		return "/api/v1/cards"
	case "/api/v1/richcards":
		return "/api/v1/richcards"
	case "/api/v1/sets":
		return "/api/v1/sets"
	case "/api/v1/types":
		return "/api/v1/types"
	case "/api/v1/classes":
		return "/api/v1/classes"
	case "/api/v1/rarities":
		return "/api/v1/rarities"
	case "/api/v1/keywords":
		return "/api/v1/keywords"
	case "/api/v1/update":
		return "/api/v1/update"
	case "/swagger.json":
		return "/swagger.json"
	default:
		return "unmatched"
	}
}
