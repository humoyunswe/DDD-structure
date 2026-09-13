package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		w := &responseBodyWriter{ResponseWriter: c.Writer, body: &bytes.Buffer{}}
		c.Writer = w

		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.NewString()
		}
		c.Set("request_id", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)

		var requestBody any
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		if len(bodyBytes) > 0 && !strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
			if err := json.Unmarshal(bodyBytes, &requestBody); err != nil {
				requestBody = string(bodyBytes)
			}
		}

		c.Next()

		latency := time.Since(start).Milliseconds()
		status := c.Writer.Status()
		api := slog.Group("api",
			slog.String("method", c.Request.Method),
			slog.String("url", c.Request.URL.String()),
			slog.String("requestId", requestID),
		)
		client := slog.Group("client",
			slog.String("ip", c.ClientIP()),
			slog.String("host", c.Request.Host),
		)

		switch {
		case status >= 200 && status < 400:
			slog.Info("request", slog.Int64("latency_ms", latency), api, client, slog.Any("requestBody", requestBody))
		case status < http.StatusInternalServerError:
			slog.Warn("client_error", slog.Int64("latency_ms", latency), api, client, slog.Any("requestBody", requestBody))
		default:
			slog.Error("server_error", slog.Int64("latency_ms", latency), api, client, slog.Any("requestBody", requestBody), slog.Any("errors", c.Errors.Errors()))
		}
	}
}
