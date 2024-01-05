package handler

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
	handlerv1 "github.com/kyong0612/fitness-supporter/generated/proto/handler/v1"
)

func (h handler) AnalyzeImage(ctx context.Context, req *connect.Request[handlerv1.AnalyzeImageRequest]) (*connect.Response[handlerv1.AnalyzeImageResponse], error) {
	slog.Info("Request headers: ",
		slog.Any("headers", req.Header()),
	)
	slog.Info("Request body: ",
		slog.Any("image_url", req.Msg.ImageUrl),
		slog.Any("user_id", req.Msg.UserId),
	)
	res := connect.NewResponse(&handlerv1.AnalyzeImageResponse{
		Ok: true,
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
