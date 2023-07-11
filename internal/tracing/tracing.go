package tracing

import (
	"context"
	"sync"

	"github.com/uptrace/uptrace-go/uptrace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var resource *sdkresource.Resource
var initResourcesOnce sync.Once

func getSampler(environment string) sdktrace.Sampler {
	switch environment {
	case "development":
		return sdktrace.AlwaysSample()
	case "production":
		return sdktrace.ParentBased(sdktrace.TraceIDRatioBased(0.5))
	default:
		return sdktrace.AlwaysSample()
	}
}

func initResource(service string) *sdkresource.Resource {
	initResourcesOnce.Do(func() {
		extraResources, _ := sdkresource.New(
			context.Background(),
			sdkresource.WithOS(),
			sdkresource.WithProcess(),
			sdkresource.WithContainer(),
			sdkresource.WithHost(),
			sdkresource.WithAttributes(attribute.Key("service.name").String(service)),
		)
		resource, _ = sdkresource.Merge(
			sdkresource.Default(),
			extraResources,
		)
	})
	return resource
}

func InitTracer(service string, environment string) (*sdktrace.TracerProvider, error) {
	exporter, err := otlptracegrpc.New(context.Background())
	if err != nil {
		return nil, err
	}

	// For the demonstration, use sdktrace.AlwaysSample sampler to sample all traces.
	// In a production application, use sdktrace.ProbabilitySampler with a desired probability.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(getSampler(environment)),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(initResource(service)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	uptrace.ConfigureOpentelemetry(
		uptrace.WithServiceName(service),
	)

	return tp, err
}

func InitMeter() (*sdkmetric.MeterProvider, error) {
	exp, err := otlpmetricgrpc.New(
		context.Background(),
	)
	if err != nil {
		return nil, err
	}

	mp := sdkmetric.NewMeterProvider(sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exp)))
	otel.SetMeterProvider(mp)
	return mp, nil
}
