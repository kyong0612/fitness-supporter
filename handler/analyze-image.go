package handler

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"regexp"

	"connectrpc.com/connect"
	"github.com/cockroachdb/errors"
	handlerv1 "github.com/kyong0612/fitness-supporter/generated/proto/handler/v1"
	"github.com/kyong0612/fitness-supporter/infra/gemini"
)

func (h handler) AnalyzeImage(ctx context.Context, req *connect.Request[handlerv1.AnalyzeImageRequest]) (*connect.Response[handlerv1.AnalyzeImageResponse], error) {
	slog.Info("Request body: ",
		slog.Any("image_url", req.Msg.GetImageUrl()),
		slog.Any("user_id", req.Msg.GetUserId()),
	)

	// get image from url
	resp, err := http.NewRequestWithContext(ctx, http.MethodGet, req.Msg.GetImageUrl(), nil)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get image from url",
			slog.Any("err", err),
		)

		return nil, errors.Wrap(err, "failed to get image from url")
	}
	defer resp.Body.Close()

	minetype := resp.Header.Get("Content-Type")

	file, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	gemini, err := gemini.NewClient(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create gemini client")
	}
	defer gemini.Close()

	analyzed, err := gemini.AnalyzeImage(ctx, minetype, file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate reply by image")
	}

	slog.InfoContext(ctx, "analyzed image",
		slog.String("analyzed", analyzed),
	)

	// parse json
	regx := regexp.MustCompile("{.*}")
	josnData := regx.FindString(analyzed)
	_ = josnData

	// TODO: persist analyzed data to bigquery
	// storage write API: https://cloud.google.com/bigquery/docs/write-api-batch?hl=ja
	// samples: https://github.com/GoogleCloudPlatform/golang-samples/tree/main/bigquery

	res := connect.NewResponse(&handlerv1.AnalyzeImageResponse{
		Ok: true,
	})
	res.Header().Set("Greet-Version", "v1")

	return res, nil
}
