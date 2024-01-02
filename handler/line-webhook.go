package handler

import (
	"log/slog"
	"net/http"

	"github.com/kyong0612/fitness-supporter/infra/line"
)

func PostLINEWebhook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	events, err := line.ParseWebhookRequest(ctx, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if len(events) == 0 {
		w.WriteHeader(http.StatusOK)

		return
	}

	line, err := line.NewClient()
	if err != nil {
		slog.ErrorContext(ctx, "failed to create line client", slog.Any("err", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	for _, event := range events {
		resp, err := line.ReplyMessage(ctx, event.ReplyToken, []string{event.Content})
		if err != nil {
			slog.ErrorContext(ctx, "failed to reply message",
				slog.Any("err", err),
			)
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		if err := resp.Body.Close(); err != nil {
			slog.ErrorContext(ctx, "failed to close response body", slog.Any("err", err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.WriteHeader(http.StatusOK)
}
