package gapi

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcLogger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	startTime := time.Now()

	result, err := handler(ctx, req)

	duration := time.Since(startTime)

	statusCode := codes.Unknown
	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}

	logger := log.Info()
	if err != nil {
		logger = log.Error().Err(err)
	}

	logger.Str("protocol", "grpc").
		Str("method", info.FullMethod).
		Int("status_code", int(statusCode)).
		Str("status_message", statusCode.String()).
		Dur("duration", duration).
		Msg("received a GRPC request")

	return result, err
}

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
}

func (rec *ResponseRecorder) WriteHeader(code int) {
	rec.StatusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *ResponseRecorder) Write(body []byte) (int, error) {
	rec.Body = body

	return rec.ResponseWriter.Write(body)
}

func HttpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		rec := &ResponseRecorder{
			ResponseWriter: res,
			StatusCode:     http.StatusOK,
		}

		handler.ServeHTTP(rec, req)

		logger := log.Info()
		if rec.StatusCode != http.StatusOK {
			logger = log.Error().Bytes("response_body", rec.Body)
		}

		logger.Str("protocol", "http").
			Str("method", req.Method).
			Str("path", req.RequestURI).
			Int("status", rec.StatusCode).
			Str("status_text", http.StatusText(rec.StatusCode)).
			Str("ip", req.RemoteAddr).
			Dur("duration", time.Since(startTime)).
			Msg("received a HTTP request")
	})
}
