package helpers

import (
	"net/http"
	"os"
	"runtime/debug"
	"time"

	logKit "github.com/go-kit/kit/log"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

// LoggingMiddleware logs the incoming HTTP request & its duration.
func loggingMiddleware(logger logKit.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Log(
						"err", err,
						"trace", debug.Stack(),
					)
				}
			}()

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			logger.Log(
				"method", r.Method,
				"status", wrapped.status,
				"path", r.URL.EscapedPath(),
				"duration", time.Since(start),
			)
		}

		return http.HandlerFunc(fn)
	}
}

func Logger(r http.Handler) http.Handler {
	var logger logKit.Logger
	// Logfmt is a structured, key=val logging format that is easy to read and parse
	logger = logKit.NewLogfmtLogger(logKit.NewSyncWriter(os.Stderr))
	// Direct any attempts to use Go's log package to our structured logger
	// log.SetOutput(logKit.NewStdlibAdapter(logger))
	// Log the timestamp (in UTC) and the callsite (file + line number) of the logging
	// call for debugging in the future.
	logger = logKit.With(logger, "t", logKit.DefaultTimestampUTC)

	// Create an instance of our LoggingMiddleware with our configured logger
	loggingMiddleware := loggingMiddleware(logger)
	return loggingMiddleware(r)
}
