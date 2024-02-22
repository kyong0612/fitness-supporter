package handler

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/kyong0612/fitness-supporter/infra/config"
	"github.com/kyong0612/fitness-supporter/infra/gcs"
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

	// upload sync data to gcs
	gcsClient, err := gcs.NewClient(ctx)
	if err != nil {
		span.RecordError(err)
		slog.ErrorContext(ctx, "failed to create gcs client", slog.Any("err", err))

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	bucket := config.Get().GCSBucketFitnessSupporter
	fileName := fmt.Sprintf("%s/%s/%d.%s", "sync", "apple-healthcare", time.Now().UnixNano(), "json")

	if err := gcsClient.Upload(ctx, bucket, fileName, body); err != nil {
		span.RecordError(err)
		slog.ErrorContext(ctx, "failed to upload to gcs", slog.Any("err", err))

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
