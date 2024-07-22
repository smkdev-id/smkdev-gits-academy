package middleware

import (
	"EkoEdyPurwanto/task-2/config"
	"EkoEdyPurwanto/task-2/model/dto/req"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func Logging(log *logrus.Logger) func(http.Handler) http.Handler {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	log.Out = file

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			// Create a custom response writer to capture the status code
			crw := &customResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(crw, r)

			endTime := time.Since(startTime)

			requestLog := req.LoggingRequest{
				StartTime:  startTime,
				EndTime:    endTime,
				StatusCode: crw.statusCode,
				ClientIP:   r.RemoteAddr,
				Method:     r.Method,
				Path:       r.URL.Path,
				UserAgent:  r.UserAgent(),
			}

			switch {
			case crw.statusCode >= 500:
				log.Error(requestLog)
			case crw.statusCode >= 400:
				log.Warn(requestLog)
			case crw.statusCode >= 300:
				log.Info(requestLog)
			case crw.statusCode >= 200:
				log.Debug(requestLog)
			case crw.statusCode >= 100:
				log.Trace(requestLog)
			default:
				log.Warnf("Unexpected status code: %d", crw.statusCode)
			}
		})
	}

}

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (crw *customResponseWriter) WriteHeader(code int) {
	crw.statusCode = code
	crw.ResponseWriter.WriteHeader(code)
}
