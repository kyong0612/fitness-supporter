package handler

import (
	"context"

	"connectrpc.com/connect"
	handlerv1 "github.com/kyong0612/fitness-supporter/proto/generated/proto/handler/v1"
	"go.opentelemetry.io/otel"
)

func (h handler) RMUAppleHealthcare(ctx context.Context, req *connect.Request[handlerv1.RMUAppleHealthcareRequest]) (*connect.Response[handlerv1.RMUAppleHealthcareResponse], error) {
	_, span := otel.Tracer("").Start(ctx, "UpdateReadModelHealthcareApple")
	defer span.End()

	// persist to bq

	res := connect.NewResponse(&handlerv1.RMUAppleHealthcareResponse{
		Ok: true,
	})
	res.Header().Set("Handler-Version", "v1")

	return res, nil
}
