package line

import (
	"context"
	"net/http"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/jarcoal/httpmock"
)

type mockClient struct{}

func NewMockClient() (Client, error) {
	return mockClient{}, nil
}

func (c mockClient) ReplyMessage(ctx context.Context, replyToken string, messages []string) (*http.Response, error) {
	resp, err := httpmock.NewJsonResponse(200, httpmock.File("infra/line/mock-data/reply-message-resp.json"))
	if err != nil {
		return nil, errors.New("failed to create mock response in ReplyMessage")
	}

	return resp, nil
}

func (c mockClient) GetContent(ctx context.Context, messageID string) (*http.Response, error) {
	image, err := os.ReadFile("infra/line/mock-data/get-content-image.JPG")
	if err != nil {
		return nil, errors.New("failed to read image file in GetContent")
	}

	return httpmock.NewBytesResponse(200, image), nil
}
