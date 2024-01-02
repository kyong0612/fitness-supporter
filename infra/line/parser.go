package line

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httputil"

	"github.com/cockroachdb/errors"
	"github.com/kyong0612/fitness-supporter/infra/config"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type MessageType string

const (
	MessageTypeText  MessageType = "text"
	MessageTypeImage MessageType = "image"
)

type MessageEvent struct {
	ReplyToken string
	Type       MessageType
	Content    string
}

func ParseWebhookRequest(ctx context.Context, req *http.Request) ([]MessageEvent, error) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dump request")
	}

	slog.InfoContext(
		ctx,
		"webhook request parsed",
		slog.String("dump", string(dump)),
	)

	cb, err := webhook.ParseRequest(
		config.Get().LINEChannelToken,
		req,
	)
	if err != nil {
		slog.ErrorContext(ctx, "failed to parse webhook request", err)
		// NOTE: healthcheckを通過させるために、エラー時はnilを返す
		return nil, nil
	}

	result := make([]MessageEvent, 0, len(cb.Events))

	for _, event := range cb.Events {
		switch e := event.(type) {
		case webhook.MessageEvent:
			replyToken := e.ReplyToken
			switch message := e.Message.(type) {
			case webhook.TextMessageContent:
				result = append(result, MessageEvent{
					ReplyToken: replyToken,
					Type:       "text",
					Content:    message.Text,
				})
			case webhook.ImageMessageContent:
				result = append(result, MessageEvent{
					ReplyToken: replyToken,
					Type:       "image",
					Content:    message.ContentProvider.OriginalContentUrl,
				})
			default:
				slog.WarnContext(
					ctx,
					"unsupported message type",
					slog.String("type", message.GetType()),
				)
			}
		default:
			slog.WarnContext(
				ctx,
				"unsupported event type",
				slog.String("type", e.GetType()),
			)
		}
	}

	return result, nil
}
