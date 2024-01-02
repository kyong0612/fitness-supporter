package gemini

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"github.com/kyong0612/fitness-supporter/infra/config"
	"google.golang.org/api/option"
)

type client struct {
	*genai.Client
}

type Client interface {
	GenerateContentByText(ctx context.Context, input string) (string, error)
	Close()
}

func NewClient(ctx context.Context) (Client, error) {
	c, err := genai.NewClient(ctx,
		option.WithAPIKey(
			config.Get().GeminiAPIKey,
		))
	if err != nil {
		return nil, err
	}

	return &client{c}, nil
}

func (c client) GenerateContentByText(ctx context.Context, input string) (string, error) {
	modal := c.GenerativeModel("gemini-pro")

	resp, err := modal.GenerateContent(ctx, genai.Text(input))
	if err != nil {
		return "", err
	}

	return fmt.Sprintln(resp.Candidates[0].Content.Parts[0]), nil
}

func (c client) Close() {
	c.Client.Close()
}
