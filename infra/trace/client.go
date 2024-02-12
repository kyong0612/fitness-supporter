package trace

import (
	"context"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.11.0"
)

const (
	spanInterval = 50 * time.Millisecond
	genInterval  = 10 * time.Second
)

var (
	OTLPEndpoint = "localhost:4317"
	ServiceName  = "fitness-supporter"
)

func InitTracer(env string) (*sdktrace.TracerProvider, error) {
	e := os.Getenv("OTLP_ENDPOINT")
	if e == "" {
		e = OTLPEndpoint
	}

	// OTLP exporter config for Collector (using default config)
	exporter, err := otlptracegrpc.New(
		context.Background(),
		otlptracegrpc.WithEndpoint(e),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	sname := os.Getenv("SERVICE_NAME")
	if sname == "" {
		sname = ServiceName
	}
	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(sname),
			semconv.ServiceVersionKey.String("1.0.0"),
			semconv.DeploymentEnvironmentKey.String("production"),
			semconv.TelemetrySDKNameKey.String("opentelemetry"),
			semconv.TelemetrySDKLanguageKey.String("go"),
			semconv.TelemetrySDKVersionKey.String("0.13.0"),
		),
	)
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return tp, nil
}
