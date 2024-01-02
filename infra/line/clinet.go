package line

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/kyong0612/fitness-supporter/infra/config"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

var _ Client = client{}

type Client interface {
	ReplyMessage(ctx context.Context, replyToken string, messages []string) (*http.Response, error)
}

type client struct {
	bot *messaging_api.MessagingApiAPI
}

func NewClient() (Client, error) {
	bot, err := messaging_api.NewMessagingApiAPI(
		config.Get().LINEChannelAccessToken,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create line client")
	}

	return client{
		bot,
	}, nil
}

func (c client) ReplyMessage(ctx context.Context, replyToken string, messages []string) (*http.Response, error) {
	textMessages := make([]messaging_api.MessageInterface, 0, len(messages))
	for _, message := range messages {
		textMessages = append(textMessages, messaging_api.TextMessage{
			Text: message,
		})
	}

	resp, _, err := c.bot.ReplyMessageWithHttpInfo(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
			Messages:   textMessages,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to reply message")
	}

	slog.InfoContext(ctx, fmt.Sprintf("status code: (%v), x-line-request-id: (%v)", resp.StatusCode, resp.Header.Get("x-line-request-id")))

	return resp, nil
}
