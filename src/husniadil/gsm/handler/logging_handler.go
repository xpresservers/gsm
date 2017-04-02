package handler

import (
	"husniadil/gsm/util/logger"
	"net/http"
	"net/http/httptest"
	"time"
)

func LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		logger.Info("Request: %s %q", req.Method, req.URL.String())

		res := httptest.NewRecorder()
		next.ServeHTTP(res, req)
		for k, v := range res.HeaderMap {
			w.Header()[k] = v
		}
		w.WriteHeader(res.Code)

		body := res.Body.String()
		code := res.Code / 100

		elapsed := time.Since(start).Seconds() * 1000

		switch code {
		case 2:
			logger.Info("Response: %s %q returned %v and took %.0fms", req.Method, req.URL.String(), res.Code, elapsed)
		default:
			logger.Warn("Response: %s %q returned %v and took %.0fms with response body: %s", req.Method, req.URL.String(), res.Code, elapsed, body)
		}

		res.Body.WriteTo(w)
	}
	return http.HandlerFunc(fn)
}
