package handler

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/kyong0612/fitness-supporter/infra/gemini"
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
		replyMsg, err := generateReply(ctx, line, event)
		if err != nil {
			slog.ErrorContext(ctx, "failed to generate reply message",
				slog.Any("err", err),
			)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		resp, err := line.ReplyMessage(ctx, event.ReplyToken, []string{replyMsg})
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

func generateReply(ctx context.Context, lineClient line.Client, event line.MessageEvent) (string, error) {
	gemini, err := gemini.NewClient(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to create gemini client")
	}
	defer gemini.Close()

	var replyMsg string

	switch event.Type {
	case line.MessageTypeText:
		replyMsg, err = gemini.GenerateContentByText(ctx, event.Content)
		if err != nil {
			return "", errors.Wrap(err, "failed to generate content by text")
		}

	case line.MessageTypeImage:
		slog.InfoContext(ctx, "image message", slog.String("messageId", event.Content))

		resp, err := lineClient.GetContent(ctx, event.Content)
		if err != nil {
			return "", errors.Wrap(err, "failed to get image")
		}
		defer resp.Body.Close()

		file := make([]byte, resp.ContentLength)
		if _, err := resp.Body.Read(file); err != nil {
			return "", errors.Wrap(err, "failed to read response body")
		}

		minetype := resp.Header.Get("Content-Type")

		replyMsg, err = gemini.GenerateContentByImage(ctx, minetype, file)
		if err != nil {
			return "", errors.Wrap(err, "failed to generate content by image")
		}

	default:
		replyMsg = "ごめんなさい、わかりません"
	}

	slog.InfoContext(ctx,
		fmt.Sprintf("generated content by %s", event.Type),
		slog.String("replyMsg", replyMsg),
	)

	return replyMsg, nil
}
