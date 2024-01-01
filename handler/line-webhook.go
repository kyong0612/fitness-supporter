package handler

import (
	"net/http"

	"github.com/kyong0612/fitness-saporter/infra/line"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, event := range events {
		if _, err := line.ReplyMessage(ctx, event.ReplyToken, []string{event.Content}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	return
}
