package handler

import (
	"log/slog"
	"net/http"
	"net/http/httputil"

	"go.opentelemetry.io/otel"
)

func (h handler) SyncHealthcareApple(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer("").Start(r.Context(), "SyncHealthcareApple")
	defer span.End()

	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		span.RecordError(err)
		slog.ErrorContext(ctx, "failed to dump request", slog.Any("err", err))

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	slog.InfoContext(ctx, "request dump", slog.String("dump", string(dump)))

	w.WriteHeader(http.StatusOK)
}
