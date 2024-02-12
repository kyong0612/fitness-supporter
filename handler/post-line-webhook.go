package handler

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/kyong0612/fitness-supporter/infra/config"
	"github.com/kyong0612/fitness-supporter/infra/gcs"
	"github.com/kyong0612/fitness-supporter/infra/gemini"
	"github.com/kyong0612/fitness-supporter/infra/line"
	"github.com/kyong0612/fitness-supporter/infra/pubsub"
	"go.opentelemetry.io/otel"
)

func (h handler) PostLINEWebhook(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer("").Start(r.Context(), "PostLINEWebhook")
	defer span.End()

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
	ctx, span := otel.Tracer("").Start(ctx, "generateReply")
	defer span.End()

	gemini, err := gemini.NewClient(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to create gemini client")
	}
	defer gemini.Close()

	replyMsg := "ごめんなさい、わかりません" // default reply message

	switch event.Type {
	case line.MessageTypeText:
		replyMsg, err = gemini.GenerateContentByText(ctx, event.Content)
		if err != nil {
			return "", errors.Wrap(err, "failed to generate content by text")
		}

	case line.MessageTypeImage:
		slog.InfoContext(ctx, "image message", slog.String("messageId", event.Content))

		replyMsg, err = generateReplyByImage(ctx, lineClient, gemini, event.Content)
		if err != nil {
			return "", errors.Wrap(err, "failed to generate reply by image")
		}
	}

	slog.InfoContext(ctx,
		fmt.Sprintf("generated content by %s", event.Type),
		slog.String("replyMsg", replyMsg),
	)

	return replyMsg, nil
}

// TODO: too long
//
//nolint:funlen
func generateReplyByImage(
	ctx context.Context,
	lineClient line.Client,
	geminiClient gemini.Client,
	messageID string,
) (string, error) {
	resp, err := lineClient.GetContent(ctx, messageID)
	if err != nil {
		return "", errors.Wrap(err, "failed to get image")
	}
	defer resp.Body.Close()

	file, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read response body")
	}

	minetype := resp.Header.Get("Content-Type")

	// upload image to gcs
	gcsClient, err := gcs.NewClient(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to create gcs client")
	}

	bucket := config.Get().GCSBucketFitnessSupporter
	fileName := fmt.Sprintf("%s.%s", uuid.New(), strings.Split(minetype, "/")[1])

	if err := gcsClient.Upload(ctx, bucket, fileName, file); err != nil {
		return "", errors.Wrap(err, "failed to upload image to gcs")
	}

	// generate content by AI
	msg, err := geminiClient.GenerateContentByImage(ctx, minetype, file)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate content by image")
	}

	filePath, err := gcsClient.GetContentURL(ctx, bucket, fileName)
	if err != nil {
		return "", errors.Wrap(err, "failed to get content url")
	}

	// publish analyze image topic
	pubsubClient, err := pubsub.NewClient(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to create pubsub client")
	}

	topicMsg, err := pubsub.NewAnalyzeImageTopic("TODO:", filePath)
	if err != nil {
		return "", errors.Wrap(err, "failed to create topic message")
	}

	if err := pubsubClient.PublishTopic(
		ctx,
		config.Get().PubSubTopicAnalyzeImage,
		topicMsg,
	); err != nil {
		return "", errors.Wrap(err, "failed to publish topic")
	}

	return fmt.Sprintf("画像は以下のURLから取得できます。\n%s\n\n%s", filePath, msg), nil
}
