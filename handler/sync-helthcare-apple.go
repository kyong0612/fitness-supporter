package handler

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"

	"go.opentelemetry.io/otel"
)

func (h handler) SyncHealthcareApple(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer("").Start(r.Context(), "SyncHealthcareApple")
	defer span.End()

	dump, err := httputil.DumpRequest(r, false)
	if err != nil {
		span.RecordError(err)
		slog.ErrorContext(ctx, "failed to dump request", slog.Any("err", err))

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	slog.InfoContext(ctx, "request dump", slog.String("dump", string(dump)))

	var body []byte
	if r.Body != nil {
		body, err = io.ReadAll(r.Body)
		if err != nil {
			span.RecordError(err)
			slog.ErrorContext(ctx, "failed to read request body", slog.Any("err", err))

			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	slog.InfoContext(ctx, "request body", slog.String("body", string(body)))

	w.WriteHeader(http.StatusOK)
}
