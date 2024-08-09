package gemini

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
	"github.com/google/generative-ai-go/genai"
	"github.com/kyong0612/fitness-supporter/infra/config"
	"go.opentelemetry.io/otel"
	"google.golang.org/api/option"
)

type client struct {
	*genai.Client
}

type Client interface {
	GenerateContentByText(ctx context.Context, input string) (string, error)
	GenerateContentByImage(ctx context.Context, minetype string, input []byte) (string, error)
	AnalyzeImage(ctx context.Context, minetype string, input []byte) (string, error)
	Close()
}

func NewClient(ctx context.Context) (Client, error) {
	c, err := genai.NewClient(ctx,
		option.WithAPIKey(
			config.Get().GeminiAPIKey,
		))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create gemini client")
	}

	return &client{c}, nil
}

func (c client) GenerateContentByText(ctx context.Context, input string) (string, error) {
	ctx, span := otel.Tracer("").Start(ctx, "gemini.GenerateContentByText")
	defer span.End()

	modal := c.GenerativeModel("gemini-1.5-flash")

	resp, err := modal.GenerateContent(ctx, genai.Text(PromptTextReplyInput(input)))
	if err != nil {
		return "", errors.Wrap(err, "failed to generate content by text")
	}

	return fmt.Sprintln(resp.Candidates[0].Content.Parts[0]), nil
}

func (c client) GenerateContentByImage(ctx context.Context, minetype string, input []byte) (string, error) {
	if len(input) == 0 {
		return "", errors.New("input image is empty")
	}

	modal := c.GenerativeModel("gemini-1.5-flash")

	resp, err := modal.GenerateContent(ctx,
		genai.Text(PromptImageReplyInput()),
		genai.Blob{
			MIMEType: minetype,
			Data:     input,
		},
	)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate content by image")
	}

	return fmt.Sprintln(resp.Candidates[0].Content.Parts[0]), nil
}

func (c client) AnalyzeImage(ctx context.Context, minetype string, input []byte) (string, error) {
	if len(input) == 0 {
		return "", errors.New("input image is empty")
	}

	modal := c.GenerativeModel("gemini-1.5-flash")

	resp, err := modal.GenerateContent(ctx,
		genai.Text(PromptAnalyzeImageInput()),
		genai.Blob{
			MIMEType: minetype,
			Data:     input,
		},
	)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate content by image")
	}

	return fmt.Sprintln(resp.Candidates[0].Content.Parts[0]), nil
}

func (c client) Close() {
	c.Client.Close()
}
